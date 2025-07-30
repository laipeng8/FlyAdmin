package repositorys

import (
	"gorm.io/gorm"
	"log"
	"server/app/models"
	"server/app/requests"
	"server/global"
)

type DepartmentRepository struct {
	Department models.Department
	Where      map[string]interface{}
}

func (r *DepartmentRepository) List(page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.Department
		offSet int
		err    error
	)

	// 初始化查询，预加载AdminUser但不加载Roles
	db := global.Db.Model(&r.Department).Preload("AdminUser", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "nickname", "real_name", "email", "name", "avatar", "phone", "gender")
	})

	// 处理查询条件
	if len(r.Where) > 0 {
		for key, value := range r.Where {
			if value == nil || value == "" {
				continue // 跳过空值
			}

			switch key {
			case "name":
				db = db.Where("name LIKE ?", "%"+value.(string)+"%")
			case "status":
				db = db.Where("status = ?", value)
			default:
				db = db.Where(key+" = ?", value)
			}
		}
	}

	// 获取总数
	if err = db.Count(&total).Error; err != nil {
		return map[string]interface{}{
			"error": "获取总数失败: " + err.Error(),
		}
	}

	// 计算偏移量
	offSet = (page - 1) * pageSize

	// 执行查询
	if err = db.Limit(pageSize).
		Order(sortField + " desc, id desc").
		Offset(offSet).
		Find(&data).Error; err != nil {
		return map[string]interface{}{
			"error": "查询数据失败: " + err.Error(),
		}
	}

	// 返回分页数据
	return global.Pages(page, pageSize, int(total), data)
}

func (r *DepartmentRepository) Add(data requests.DepartmentAdd) (*gorm.DB, models.Department) {

	r.Department.Name = data.Name
	r.Department.Status = data.Status
	r.Department.Sort = data.Sort
	return global.Db.Create(&r.Department), r.Department
}

func (r *DepartmentRepository) Update(data requests.DepartmentUpdate) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		var model models.Department
		model.ID = data.Id

		// 更新仓库中的部门信息
		r.Department.ID = data.Id
		r.Department.Name = data.Name
		r.Department.Status = data.Status
		r.Department.Sort = data.Sort

		// 执行更新操作
		db := tx.Where("id = ?", data.Id).Updates(&r.Department)
		if db.Error != nil {
			// 记录错误日志
			log.Printf("Failed to upload department with ID %d: %v", data.Id, db.Error)
			return db.Error
		}

		// 检查是否有记录受到影响
		if db.RowsAffected == 0 {
			log.Printf("No department record found with ID %d", data.Id)
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}

// 更新部门下的管理员（先删除原有关系，再添加新关系）
func (r *DepartmentRepository) UpdateDepartmentUsers(departmentID uint, userIDs []uint) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		// 删除部门原有的所有用户关系
		if err := tx.Where("department_id = ?", departmentID).
			Delete(&models.UserDepartment{}).Error; err != nil {
			return err
		}

		// 添加新的用户关系
		var userDepartments []models.UserDepartment
		for _, userID := range userIDs {
			userDepartments = append(userDepartments, models.UserDepartment{
				AdminUserID:  userID,
				DepartmentID: departmentID,
			})
		}

		if len(userDepartments) > 0 {
			if err := tx.Create(&userDepartments).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// 为部门添加管理员（不删除原有关系）
func (r *DepartmentRepository) AddDepartmentUsers(departmentID uint, userIDs []uint) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		// 查询已存在的用户关系
		var existingRelations []models.UserDepartment
		if err := tx.Where("department_id = ? AND admin_user_id IN ?",
			departmentID, userIDs).Find(&existingRelations).Error; err != nil {
			return err
		}

		// 构建已存在用户ID的map用于快速查找
		existingUserIDs := make(map[uint]bool)
		for _, relation := range existingRelations {
			existingUserIDs[relation.AdminUserID] = true
		}

		// 添加新的用户关系（跳过已存在的）
		var newRelations []models.UserDepartment
		for _, userID := range userIDs {
			if !existingUserIDs[userID] {
				newRelations = append(newRelations, models.UserDepartment{
					AdminUserID:  userID,
					DepartmentID: departmentID,
				})
			}
		}

		if len(newRelations) > 0 {
			if err := tx.Create(&newRelations).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

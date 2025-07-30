package repositorys

import (
	"gorm.io/gorm"
	"server/app/models"
	"server/app/requests"
	"server/global"
)

type RoleRepository struct {
	Model models.Role
	Where map[string]interface{}
}

func (r *RoleRepository) Group() []models.Role {
	var (
		group []models.Role
		err   error
	)
	db := global.Db.Model(&models.Role{})
	if err = db.Find(&group).Error; err != nil {
		return []models.Role{}
	}
	return group
}

func (r *RoleRepository) List(page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.Role
		offSet int
		err    error
	)

	// 初始化查询
	db := global.Db.Model(&r.Model)

	// 处理查询条件
	if len(r.Where) > 0 {
		for key, value := range r.Where {
			if value == nil || value == "" {
				continue // 跳过空值
			}

			switch key {
			case "label":
				// 角色名称模糊查询
				db = db.Where("label LIKE ?", "%"+value.(string)+"%")
			case "status":
				// 状态精确查询
				db = db.Where("status = ?", value)
			default:
				// 其他字段精确查询
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
	if err = db.Preload("Menus").
		Limit(pageSize).
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

/*
添加角色
*/
func (r *RoleRepository) Add(post requests.Role) error {
	db := global.Db.Create(&models.Role{
		Alias:  post.Alias,
		Label:  post.Label,
		Sort:   post.Sort,
		Remark: post.Remark,
		Status: &post.Status,
	})
	return db.Error
}

/*
更新角色
*/
func (r *RoleRepository) Update(post requests.Role) error {

	return global.Db.Transaction(func(sessionDb *gorm.DB) error {
		return sessionDb.Debug().Where("id = ?", post.Id).Updates(&models.Role{
			Alias:  post.Alias,
			Label:  post.Label,
			Sort:   post.Sort,
			Remark: post.Remark,
			Status: &post.Status,
		}).Error
	})

}

func (r *RoleRepository) UpMenus(post requests.RoleUpMenus) error {
	var role models.Role
	role.ID = post.Id

	if len(post.Menus) > 0 {

		var replace []models.AdminMenu

		for _, v := range post.Menus {
			var li models.AdminMenu
			li.ID = v
			replace = append(replace, li)

		}
		return global.Db.Model(&role).Omit("Menus.*").Association("Menus").Replace(replace)
	} else {
		return global.Db.Model(&role).Association("Menus").Clear()
	}

}

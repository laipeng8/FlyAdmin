package repositorys

import (
	"server/app/models"
	"server/global"
	"strings"
)

type MenuApiRepository struct {
	BaseRepository
	MenuApiList models.MenuApiList
	Where       map[string]interface{}
}

func (u *MenuApiRepository) List(page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total int64
		data  []models.MenuApiList
	)

	// 初始化数据库查询
	db := global.Db.Model(&u.MenuApiList)

	// 处理查询条件
	for key, value := range u.Where {
		if value == nil || value == "" {
			continue
		}

		// 统一转换为字符串处理
		strValue, ok := value.(string)
		if !ok {
			strValue = ""
		}

		switch key {
		case "nickname", "name", "real_name":
			db = db.Where(key+" LIKE ?", "%"+strings.TrimSpace(strValue)+"%")
		default:
			db = db.Where(key+" = ?", value)
		}
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return global.Pages(page, pageSize, 0, nil)
	}

	// 处理排序字段
	sortOrder := "desc"
	if sortField == "" {
		sortField = "created_at"
	}

	// 执行查询
	offset := (page - 1) * pageSize
	err := db.Limit(pageSize).
		Offset(offset).
		Order(sortField + " " + sortOrder).
		Order("id " + sortOrder).
		Find(&data).Error

	if err != nil {
		return global.Pages(page, pageSize, 0, nil)
	}

	return global.Pages(page, pageSize, int(total), data)
}

// Create 创建API
func (r *MenuApiRepository) Create(api *models.MenuApiList) error {
	return global.Db.Create(api).Error
}

// Update 更新API
func (r *MenuApiRepository) Update(api *models.MenuApiList) error {
	return global.Db.Save(api).Error
}

// Delete 删除API
func (r *MenuApiRepository) Delete(id uint) error {
	return global.Db.Delete(&models.MenuApiList{}, id).Error
}

// BatchDelete 批量删除API
func (r *MenuApiRepository) BatchDelete(ids []uint) error {
	return global.Db.Delete(&models.MenuApiList{}, ids).Error
}

// FindByID 根据ID查找API
func (r *MenuApiRepository) FindByID(id uint) (*models.MenuApiList, error) {
	var api models.MenuApiList
	err := global.Db.Preload("Menu").First(&api, id).Error
	return &api, err
}

// FindAll 查找所有API
func (r *MenuApiRepository) FindAll() ([]models.MenuApiList, error) {
	var apis []models.MenuApiList
	err := global.Db.Preload("Menu").Find(&apis).Error
	return apis, err
}

// FindByMenuID 根据菜单ID查找API
func (r *MenuApiRepository) FindByMenuID(menuID uint) ([]models.MenuApiList, error) {
	var apis []models.MenuApiList
	err := global.Db.Where("menu_id = ?", menuID).Preload("Menu").Find(&apis).Error
	return apis, err
}

// FindByCode 根据代码查找API
func (r *MenuApiRepository) FindByCode(code string) (*models.MenuApiList, error) {
	var api models.MenuApiList
	err := global.Db.Where("code = ?", code).Preload("Menu").First(&api).Error
	return &api, err
}

// FindByURL 根据URL查找API
func (r *MenuApiRepository) FindByURL(url string) (*models.MenuApiList, error) {
	var api models.MenuApiList
	err := global.Db.Where("url = ?", url).Preload("Menu").First(&api).Error
	return &api, err
}

// CheckCodeExists 检查代码是否存在
func (r *MenuApiRepository) CheckCodeExists(code string, excludeID ...uint) (bool, error) {
	var count int64
	query := global.Db.Model(&models.MenuApiList{}).Where("code = ?", code)
	if len(excludeID) > 0 {
		query = query.Where("id != ?", excludeID[0])
	}
	err := query.Count(&count).Error
	return count > 0, err
}

// CheckURLExists 检查URL是否存在
func (r *MenuApiRepository) CheckURLExists(url string, excludeID ...uint) (bool, error) {
	var count int64
	query := global.Db.Model(&models.MenuApiList{}).Where("url = ?", url)
	if len(excludeID) > 0 {
		query = query.Where("id != ?", excludeID[0])
	}
	err := query.Count(&count).Error
	return count > 0, err
}

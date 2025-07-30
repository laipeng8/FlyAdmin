package repositorys

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/app/models"
	"server/app/requests"
	"server/global"
	"strconv"
)

type SystemMenuRepository struct {
	MenuModel models.AdminMenu
	BaseRepository
}

// Add
//
//	@Description:	添加菜单
//	@receiver		menu *SystemMenuRepository
//	@param			post	requests.MenuPost
//	@return			*gorm.DB
//	@return			models.AdminMenu
func (menu *SystemMenuRepository) Add(post requests.MenuPost) (*gorm.DB, models.AdminMenu) {
	menu.MenuModel.Name = post.Name
	menu.MenuModel.Component = post.Component
	menu.MenuModel.Meta = post.Meta
	menu.MenuModel.ParentId = post.ParentId
	menu.MenuModel.Path = post.Path
	menu.MenuModel.Redirect = post.Redirect
	fmt.Println(menu.MenuModel.ParentId)
	return menu.getDb().Create(&menu.MenuModel), menu.MenuModel
}

// Update 更新菜单
func (menu *SystemMenuRepository) Update(post requests.MenuPost) (error, models.AdminMenu) {
	var updateData models.AdminMenu
	updateData.Name = post.Name
	updateData.Component = post.Component
	updateData.Meta = post.Meta
	updateData.ParentId = post.ParentId
	updateData.Path = post.Path
	updateData.Redirect = post.Redirect
	updateData.Sort = post.Sort
	menu.MenuModel.ID = uint(post.Id)

	return global.Db.Transaction(func(sessionDb *gorm.DB) error {
		var notDelIds []uint
		for _, v := range post.ApiList {
			var apiList models.MenuApiList
			if len(v["id"]) > 0 {
				//更新业务逻辑
				id, _ := strconv.Atoi(v["id"])
				apiList.ID = uint(id)
				delete(v, "id")
				upDb := sessionDb.Model(&apiList).Updates(models.MenuApiList{
					Code: v["code"],
					Url:  v["url"],
				})

				if upDb.Error != nil {
					return upDb.Error
				}
				notDelIds = append(notDelIds, apiList.ID)
			} else {
				//新增业务逻辑
				var addModel = models.MenuApiList{
					MenuId: menu.MenuModel.ID,
					Code:   v["code"],
					Url:    v["url"],
				}
				if adDb := sessionDb.Create(&addModel); adDb.Error != nil {
					return adDb.Error
				}
				notDelIds = append(notDelIds, addModel.ID)

			}
		}

		//同步，自动删除不存在的id
		var syncDb *gorm.DB
		if len(notDelIds) > 0 {
			syncDb = sessionDb.Not(notDelIds).Where("menu_id = ?", post.Id).Delete(&models.MenuApiList{})
		} else {
			syncDb = sessionDb.Debug().Where("menu_id = ?", post.Id).Delete(&models.MenuApiList{})
		}

		if syncDb.Error != nil {
			return syncDb.Error
		}
		fmt.Printf("---->%+v", updateData)
		return sessionDb.Model(&menu.MenuModel).Select("*").Omit("id", "created_at", "deleted_at", "updated_at").Updates(updateData).Error
	}), menu.MenuModel
}

func (menu *SystemMenuRepository) ArrayToTree(arr []models.AdminMenu, pid uint) []*models.TreeMenu {

	// 创建一个 map 用来存储所有的节点
	nodeMap := make(map[uint]*models.TreeMenu)
	rootNodes := make([]*models.TreeMenu, 0)

	// 遍历节点列表，将每个节点放入 map 中
	for _, menu := range arr {
		menuCopy := menu // 创建副本，避免引用原始数据
		nodeMap[menuCopy.ID] = &models.TreeMenu{
			AdminMenu: menuCopy,
		}
	}

	// 遍历节点列表，建立父子关系
	for _, menu := range arr {
		if parentId, exists := nodeMap[menu.ParentId]; exists {
			parentId.Children = append(parentId.Children, nodeMap[menu.ID])
		} else {
			// 如果找不到父节点，则认为这是一个根节点
			rootNodes = append(rootNodes, nodeMap[menu.ID])
		}
	}
	return rootNodes
}

func (menu *SystemMenuRepository) MenuTree() interface{} {
	var all []models.AdminMenu
	global.Db.Order("sort desc").Find(&all)

	return menu.ArrayToTree(all, 0)
}

func (menu *SystemMenuRepository) GetCustomClaims(c *gin.Context) (*models.CustomClaims, bool) {
	v, ok := c.Get("claims")
	claims, err := v.(*models.CustomClaims)
	if ok && err {
		return claims, true
	} else {
		return &models.CustomClaims{}, false
	}
}

// GetApiList 根据当前登陆得用户获得api 权限
func (menu *SystemMenuRepository) GetApiList(c *gin.Context, apiList *[]models.MenuApiList) error {
	claims, ok := menu.GetCustomClaims(c)
	if ok {
		var adminUser models.AdminUser
		adminUser.ID = claims.Id
		global.Db.Model(&adminUser).Preload("Roles.Menus.ApiList").First(&adminUser)
		return nil
	} else {
		return errors.New("无法处理")
	}
}

// BuildMenuTree 构建菜单树
func (menu *SystemMenuRepository) BuildMenuTree(menus []models.AdminMenu) []*models.TreeMenu {
	menuMap := make(map[uint]*models.TreeMenu)
	var rootMenus []*models.TreeMenu

	// 先将所有菜单转换为 TreeMenu 并存储到 map 中
	for _, m := range menus {
		treeMenu := &models.TreeMenu{
			AdminMenu: m,
			Children:  []*models.TreeMenu{},
		}
		menuMap[m.ID] = treeMenu
	}

	// 构建树结构
	for _, m := range menus {
		if m.ParentId == 0 {
			rootMenus = append(rootMenus, menuMap[m.ID])
		} else {
			if parent, ok := menuMap[m.ParentId]; ok {
				parent.Children = append(parent.Children, menuMap[m.ID])
			}
		}
	}

	return rootMenus
}

// GetPermissionByUser 根据传递用户对象
func (menu *SystemMenuRepository) GetPermissionByUser(adminUser models.AdminUser, permission *[]string) error {
	for _, role := range adminUser.Roles {
		for _, menu := range role.Menus {
			*permission = append(*permission, menu.Name)
		}
	}
	return nil
}

// GetApiListToMap 获取map apiList
func (menu *SystemMenuRepository) GetApiListToMap(c *gin.Context, apiListMap *map[string]string) error {
	var apiList []models.MenuApiList
	err := menu.GetApiList(c, &apiList)
	if err != nil {
		return err
	} else {
		for _, v := range apiList {
			(*apiListMap)[v.Url] = v.Code
		}
		return nil
	}
}

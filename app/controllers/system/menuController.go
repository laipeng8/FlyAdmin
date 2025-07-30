package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"server/app/models"
	"server/app/repositorys"
	"server/app/requests"
	"server/global"
	"server/global/response"
)

type MenuController struct{}

func (m *MenuController) Add(c *gin.Context) {

	var (
		postData       requests.MenuPost
		menuRepository repositorys.SystemMenuRepository
	)
	_ = c.ShouldBind(&postData)
	result, model := menuRepository.Add(postData)

	if result.Error == nil {
		response.Success(c, "ok", model)
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (m *MenuController) Update(c *gin.Context) {
	var (
		postData       requests.MenuPost
		menuRepository repositorys.SystemMenuRepository
	)
	if bindErr := c.ShouldBindBodyWith(&postData, binding.JSON); bindErr != nil {
		response.Failed(c, bindErr.Error())
		return
	}
	global.Logger.Infof("%+v", postData)
	err, model := menuRepository.Update(postData)

	if err == nil {
		global.Db.Preload("ApiList").Find(&model)
		response.Success(c, "ok", model)
	} else {
		response.Failed(c, err.Error())
	}
}

func (m *MenuController) All(c *gin.Context) {
	var (
		menuRepository repositorys.SystemMenuRepository
	)
	response.Success(c, "ok", menuRepository.MenuTree())
}

func (m *MenuController) Del(c *gin.Context) {

	var (
		delIds requests.MenuDel
	)
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.AdminMenu{}, delIds.Ids)

	if result.Error == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (m *MenuController) MenuPermissions(c *gin.Context) {
	var (
		myMenus        []*models.TreeMenu
		menus          []models.AdminMenu
		adminUser      models.AdminUser
		menuRepository repositorys.SystemMenuRepository
	)

	v, ok := c.Get("claims")
	if !ok {
		response.Success(c, "ok", []interface{}{})
		return
	}

	claims, ok := v.(*models.CustomClaims)
	if !ok {
		response.Success(c, "ok", []interface{}{})
		return
	}

	if global.IsSuperAdmin(claims.Roles, global.SuperAdmin) {
		global.Db.Order("sort desc").Find(&menus)
	} else {
		adminUser.ID = claims.Id
		global.Db.Model(&adminUser).Preload("Roles").Preload("Roles.Menus", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort desc")
		}).First(&adminUser)
		for _, v := range adminUser.Roles {
			menus = append(menus, v.Menus...)
		}
	}

	myMenus = menuRepository.ArrayToTree(menus, 0)
	cleanMenus := removeMenuFields(myMenus)
	response.Success(c, "ok", cleanMenus)
}

func removeMenuFields(menus []*models.TreeMenu) []map[string]interface{} {
	var result []map[string]interface{}

	for _, menu := range menus {
		menuMap := make(map[string]interface{})

		// 按照指定顺序添加字段，并检查是否为空
		if menu.ID != 0 {
			menuMap["id"] = menu.ID
		}

		if menu.Name != "" {
			menuMap["name"] = menu.Name
		}

		if menu.Path != "" {
			menuMap["path"] = menu.Path
		}

		if menu.Component != "" {
			menuMap["component"] = menu.Component
		}

		if menu.Meta != nil {
			menuMap["meta"] = menu.Meta
		}

		// 处理子菜单
		if menu.Children != nil && len(menu.Children) > 0 {
			children := removeMenuFields(menu.Children)
			if len(children) > 0 {
				menuMap["children"] = children
			}
		}

		// 只有当有字段时才添加到结果中
		if len(menuMap) > 0 {
			result = append(result, menuMap)
		}
	}

	return result
}

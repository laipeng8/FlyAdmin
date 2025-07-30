package system

import (
	"strconv"

	"server/app/models"
	"server/app/repositorys"
	"server/app/requests"
	"server/global"
	"server/global/response"

	"github.com/gin-gonic/gin"
)

type MenuApiController struct{}

// GetMenuApiList 获取API列表
func (c *MenuApiController) GetMenuApiList(ctx *gin.Context) {
	var (
		params            global.List
		menuApiRepository repositorys.MenuApiRepository
	)
	_ = ctx.ShouldBind(&params)

	menuApiRepository.Where = params.Where
	response.Success(ctx, "ok", menuApiRepository.List(params.Page, params.PageSize, "created_at"))
}

// CreateMenuApi 创建API
func (c *MenuApiController) CreateMenuApi(ctx *gin.Context) {
	var req requests.CreateMenuApiRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}

	repo := &repositorys.MenuApiRepository{}

	// 检查代码是否已存在
	exists, err := repo.CheckCodeExists(req.Code)
	if err != nil {
		response.FailWithMessage("检查API代码失败: "+err.Error(), ctx)
		return
	}
	if exists {
		response.FailWithMessage("API代码已存在", ctx)
		return
	}

	// 检查URL是否已存在
	exists, err = repo.CheckURLExists(req.Url)
	if err != nil {
		response.FailWithMessage("检查API地址失败: "+err.Error(), ctx)
		return
	}
	if exists {
		response.FailWithMessage("API地址已存在", ctx)
		return
	}

	// 验证菜单是否存在
	var menu models.AdminMenu
	err = global.Db.First(&menu, req.MenuId).Error
	if err != nil {
		response.FailWithMessage("菜单不存在", ctx)
		return
	}

	api := &models.MenuApiList{
		Code:     req.Code,
		Url:      req.Url,
		MenuId:   req.MenuId,
		Describe: req.Describe,
	}

	if err := repo.Create(api); err != nil {
		response.FailWithMessage("创建API失败: "+err.Error(), ctx)
		return
	}

	response.OkWithMessage("API创建成功", ctx)
}

// UpdateMenuApi 更新API
func (c *MenuApiController) UpdateMenuApi(ctx *gin.Context) {
	var req requests.UpdateMenuApiRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}

	repo := &repositorys.MenuApiRepository{}

	// 检查API是否存在
	api, err := repo.FindByID(req.ID)
	if err != nil {
		response.FailWithMessage("API不存在", ctx)
		return
	}

	// 检查代码是否已存在（排除当前记录）
	exists, err := repo.CheckCodeExists(req.Code, req.ID)
	if err != nil {
		response.FailWithMessage("检查API代码失败: "+err.Error(), ctx)
		return
	}
	if exists {
		response.FailWithMessage("API代码已存在", ctx)
		return
	}

	// 检查URL是否已存在（排除当前记录）
	exists, err = repo.CheckURLExists(req.Url, req.ID)
	if err != nil {
		response.FailWithMessage("检查API地址失败: "+err.Error(), ctx)
		return
	}
	if exists {
		response.FailWithMessage("API地址已存在", ctx)
		return
	}

	// 验证菜单是否存在
	var menu models.AdminMenu
	err = global.Db.First(&menu, req.MenuId).Error
	if err != nil {
		response.FailWithMessage("菜单不存在", ctx)
		return
	}

	// 更新API信息
	api.Code = req.Code
	api.Url = req.Url
	api.MenuId = req.MenuId
	api.Describe = req.Describe

	if err := repo.Update(api); err != nil {
		response.FailWithMessage("更新API失败: "+err.Error(), ctx)
		return
	}

	response.OkWithMessage("API更新成功", ctx)
}

// DeleteMenuApi 删除API
func (c *MenuApiController) DeleteMenuApi(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("API ID格式错误", ctx)
		return
	}

	repo := &repositorys.MenuApiRepository{}

	// 检查API是否存在
	_, err = repo.FindByID(uint(id))
	if err != nil {
		response.FailWithMessage("API不存在", ctx)
		return
	}

	if err := repo.Delete(uint(id)); err != nil {
		response.FailWithMessage("删除API失败: "+err.Error(), ctx)
		return
	}

	response.OkWithMessage("API删除成功", ctx)
}

// BatchDeleteMenuApi 批量删除API
func (c *MenuApiController) BatchDeleteMenuApi(ctx *gin.Context) {
	var req requests.BatchDeleteMenuApiRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}

	repo := &repositorys.MenuApiRepository{}

	if err := repo.BatchDelete(req.IDs); err != nil {
		response.FailWithMessage("批量删除API失败: "+err.Error(), ctx)
		return
	}

	response.OkWithMessage("批量删除API成功", ctx)
}

// GetMenuApi 获取API详情
func (c *MenuApiController) GetMenuApi(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("API ID格式错误", ctx)
		return
	}

	repo := &repositorys.MenuApiRepository{}
	api, err := repo.FindByID(uint(id))
	if err != nil {
		response.FailWithMessage("API不存在", ctx)
		return
	}

	response.OkWithData(api, ctx)
}

// GetMenuApiByMenuID 根据菜单ID获取API列表
func (c *MenuApiController) GetMenuApiByMenuID(ctx *gin.Context) {
	menuIDStr := ctx.Param("menu_id")
	menuID, err := strconv.ParseUint(menuIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("菜单ID格式错误", ctx)
		return
	}

	repo := &repositorys.MenuApiRepository{}
	apis, err := repo.FindByMenuID(uint(menuID))
	if err != nil {
		response.FailWithMessage("查询API失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(apis, ctx)
}

// GetAllMenuApis 获取所有API
func (c *MenuApiController) GetAllMenuApis(ctx *gin.Context) {
	repo := &repositorys.MenuApiRepository{}
	apis, err := repo.FindAll()
	if err != nil {
		response.FailWithMessage("查询API失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(apis, ctx)
}

// CheckMenuApiCode 检查API代码是否存在
func (c *MenuApiController) CheckMenuApiCode(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		response.FailWithMessage("API代码不能为空", ctx)
		return
	}

	excludeIDStr := ctx.Query("exclude_id")
	var excludeID uint
	if excludeIDStr != "" {
		if id, err := strconv.ParseUint(excludeIDStr, 10, 64); err == nil {
			excludeID = uint(id)
		}
	}

	repo := &repositorys.MenuApiRepository{}
	exists, err := repo.CheckCodeExists(code, excludeID)
	if err != nil {
		response.FailWithMessage("检查API代码失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(gin.H{
		"exists": exists,
		"code":   code,
	}, ctx)
}

// CheckMenuApiURL 检查API地址是否存在
func (c *MenuApiController) CheckMenuApiURL(ctx *gin.Context) {
	url := ctx.Query("url")
	if url == "" {
		response.FailWithMessage("API地址不能为空", ctx)
		return
	}

	excludeIDStr := ctx.Query("exclude_id")
	var excludeID uint
	if excludeIDStr != "" {
		if id, err := strconv.ParseUint(excludeIDStr, 10, 64); err == nil {
			excludeID = uint(id)
		}
	}

	repo := &repositorys.MenuApiRepository{}
	exists, err := repo.CheckURLExists(url, excludeID)
	if err != nil {
		response.FailWithMessage("检查API地址失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(gin.H{
		"exists": exists,
		"url":    url,
	}, ctx)
}

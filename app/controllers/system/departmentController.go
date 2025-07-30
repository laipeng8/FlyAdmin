package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/app/models"
	"server/app/repositorys"
	"server/app/requests"
	"server/global"
	"server/global/response"
	"strconv"
)

type DepartmentController struct {
}

func (controller *DepartmentController) Add(c *gin.Context) {
	var (
		departmentRepository repositorys.DepartmentRepository
		data                 requests.DepartmentAdd
		err                  error
	)
	err = c.ShouldBind(&data)
	if err != nil {
		response.FrontDataError(err, c)
		return
	}
	result, model := departmentRepository.Add(data)
	if result.Error == nil {
		response.Success(c, "ok", model)
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (controller *DepartmentController) Edit(c *gin.Context) {
	var (
		departmentRepository repositorys.DepartmentRepository
		data                 requests.DepartmentUpdate
		err                  error
	)
	err = c.ShouldBind(&data)
	if err != nil {
		response.FrontDataError(err, c)
		return
	}
	err = departmentRepository.Update(data)
	if err != nil {
		response.UpdateErrorMsg(err, c)
	} else {
		response.UpdateSuccessMsg(c)
	}
}

func (controller *DepartmentController) Delete(c *gin.Context) {
	var delIds global.Del
	_ = c.ShouldBind(&delIds)
	result := global.Db.Delete(&models.Department{}, delIds.Ids)
	if result.Error == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, result.Error.Error())
	}
}
func (controller *DepartmentController) Del(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.FrontDataError(err, c)
		return
	}
	result := global.Db.Delete(&models.Department{}, id)
	if result.Error == nil {
		response.Success(c, "ok", "")
	} else {
		response.Failed(c, result.Error.Error())
	}
}

func (controller *DepartmentController) List(c *gin.Context) {
	var (
		params               requests.RoleList
		departmentRepository repositorys.DepartmentRepository
		err                  error
	)
	err = c.ShouldBind(&params)
	if err = c.ShouldBind(&params); err != nil {
		response.FrontDataError(err, c)
		return
	}
	// 设置默认值
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}

	departmentRepository.Where = params.Where
	result := departmentRepository.List(params.Page, params.PageSize, "sort")

	if _, ok := result["error"]; ok {
		response.Fail(c, result["error"].(string), nil)
		return
	}
	response.Success(c, "获取部门列表成功", result)

}

func (controller *DepartmentController) UserUpDepart(c *gin.Context) {
	var (
		params         requests.DepartmentUserUpdate
		departmentRepo repositorys.DepartmentRepository
	)

	// 加强绑定验证
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, "请求参数错误: "+err.Error(), nil)
		return
	}

	// 验证部门是否存在
	var department models.Department
	if err := global.Db.First(&department, params.DepartmentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, "部门不存在", nil)
		} else {
			response.Fail(c, "查询部门失败", nil)
		}
		return
	}

	// 验证所有用户是否存在
	var userCount int64
	if err := global.Db.Model(&models.AdminUser{}).
		Where("id IN ?", params.UserIDs).
		Count(&userCount).Error; err != nil {
		response.Fail(c, "查询用户失败", nil)
		return
	}

	if userCount != int64(len(params.UserIDs)) {
		response.Fail(c, "部分用户不存在", nil)
		return
	}

	// 执行更新
	if err := departmentRepo.UpdateDepartmentUsers(params.DepartmentID, params.UserIDs); err != nil {
		response.Fail(c, "更新部门用户失败: "+err.Error(), nil)
		return
	}

	response.Success(c, "更新部门用户成功", nil)
}

func (controller *DepartmentController) UserAddDepart(c *gin.Context) {
	var (
		params         requests.DepartmentUserAdd
		departmentRepo repositorys.DepartmentRepository
	)

	// 加强绑定验证
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, "请求参数错误: "+err.Error(), nil)
		return
	}

	// 验证部门是否存在
	var department models.Department
	if err := global.Db.First(&department, params.DepartmentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, "部门不存在", nil)
		} else {
			response.Fail(c, "查询部门失败", nil)
		}
		return
	}

	// 验证所有用户是否存在
	var userCount int64
	if err := global.Db.Model(&models.AdminUser{}).
		Where("id IN ?", params.UserIDs).
		Count(&userCount).Error; err != nil {
		response.Fail(c, "查询用户失败", nil)
		return
	}

	if userCount != int64(len(params.UserIDs)) {
		response.Fail(c, "部分用户不存在", nil)
		return
	}

	// 执行添加
	if err := departmentRepo.AddDepartmentUsers(params.DepartmentID, params.UserIDs); err != nil {
		response.Fail(c, "添加部门用户失败: "+err.Error(), nil)
		return
	}

	response.Success(c, "添加部门用户成功", nil)
}

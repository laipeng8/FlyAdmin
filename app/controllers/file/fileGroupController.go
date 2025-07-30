package file

import (
	"github.com/gin-gonic/gin"
	"server/app/repositorys"
	"server/app/requests"
	"server/global"
	"server/global/response"
	"strconv"
)

type FileGroupController struct{}

// Index 根据ID查询文件组及其子级树形结构
func (controller *FileGroupController) Index(ctx *gin.Context) {
	// 1. 获取ID参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 2. 查询数据
	repo := repositorys.FileGroupRepository{}
	tree, err := repo.GetTreeByRootID(uint(id))
	if err != nil {
		global.Logger.Error("查询文件组失败:", err)
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if tree == nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 3. 返回结果
	response.OkWithData(tree, ctx)
}
func (controller *FileGroupController) Save(ctx *gin.Context) {
	var fileGroup requests.FileGroupAdd
	if err := ctx.ShouldBindJSON(&fileGroup); err != nil {
		response.FailWithMessage("参数绑定失败: "+err.Error(), ctx)
		return
	}

	repo := repositorys.FileGroupRepository{}
	if err := repo.Create(&fileGroup); err != nil {
		global.Logger.Error("创建文件组失败:", err)
		response.FailWithMessage("创建文件组失败: "+err.Error(), ctx)
		return
	}

	response.OkWithData(fileGroup, ctx)
}

func (controller *FileGroupController) Edit(ctx *gin.Context) {
	var (
		data requests.FileGroupUpdate
		repo repositorys.FileGroupRepository
	)
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	err = repo.Update(data)
	if err == nil {
		response.UpdateSuccessMsg(ctx)
	} else {
		response.UpdateErrorMsg(err, ctx)
	}
}

func (controller *FileGroupController) Delete(ctx *gin.Context) {
	// 获取ID参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}

	// 执行强制删除
	repo := repositorys.FileGroupRepository{}
	if err := repo.ForceDelete(uint(id)); err != nil {
		global.Logger.Error("强制删除失败:", err)
		response.FailWithMessage("强制删除失败: "+err.Error(), ctx)
		return
	}

	response.DeleteSuccessMsg(ctx)
}
func (controller *FileGroupController) List(ctx *gin.Context) {
	repo := repositorys.FileGroupRepository{}

	// 查询所有文件组
	fileGroups, err := repo.FindAll()
	if err != nil {
		global.Logger.Error("查询文件组失败:", err)
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 构建树形结构
	tree := repo.BuildTree(fileGroups, 0)

	response.OkWithData(tree, ctx)
}

func (controller *FileGroupController) Check(ctx *gin.Context) {
	// 1. 获取ID参数
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}

	// 2. 创建Repository实例
	repo := repositorys.FileGroupRepository{}

	// 3. 检查文件组是否存在
	exists, err := repo.Exists(uint(id))
	if err != nil {
		global.Logger.Error("检查文件组失败:", err)
		response.FailWithMessage("检查文件组失败: "+err.Error(), ctx)
		return
	}
	if !exists {
		response.FailWithMessage("文件组不存在", ctx)
		return
	}

	// 4. 获取子文件组信息
	children, err := repo.GetChildren(uint(id))
	if err != nil {
		global.Logger.Error("获取子文件组失败:", err)
		response.FailWithMessage("获取子文件组失败: "+err.Error(), ctx)
		return
	}

	// 5. 获取关联文件信息
	files, err := repo.GetFiles(uint(id))
	if err != nil {
		global.Logger.Error("获取关联文件失败:", err)
		response.FailWithMessage("获取关联文件失败: "+err.Error(), ctx)
		return
	}

	// 6. 构造返回数据
	responseData := gin.H{
		"group_id":     id,
		"has_children": len(children) > 0,
		"children":     children,
		"has_files":    len(files) > 0,
		"files":        files,
	}

	// 7. 返回成功响应
	response.OkWithData(responseData, ctx)
}

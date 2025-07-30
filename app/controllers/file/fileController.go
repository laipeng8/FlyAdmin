package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
	"os"
	"path/filepath"
	"server/app/models"
	"server/app/repositorys"
	"server/global"
	"server/global/response"
	"strconv"
	"strings"
	"time"
)

type FileController struct{}

func (controller *FileController) Test(c *gin.Context) {
	allow := map[string]string{
		"image/jpeg": "jpg",
		"image/png":  "png",
	}

	newUploadPath := "." + global.Config.App.UploadFile + "/" + time.Now().Format("20060102")
	file, _ := c.FormFile("file")
	fileType, ok := allow[file.Header.Get("Content-Type")]
	if !ok {
		response.Failed(c, "当前类型不允许上传！")
		return
	}

	uuid := uuid2.NewV4()

	dirErr := os.MkdirAll(newUploadPath, os.ModePerm)

	if dirErr != nil {
		response.Failed(c, "文件目录创建错误:"+dirErr.Error())
		return
	}

	fileName := uuid.String() + "." + fileType

	allDir := newUploadPath + "/" + fileName
	err := c.SaveUploadedFile(file, allDir)

	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	response.Success(c, "ok", gin.H{
		"id":       uuid,
		"fileName": fileName,
		"src":      global.Config.WX.Url + "/static/system/common/file" + allDir[8:],
	})
}
func (controller *FileController) Upload(c *gin.Context) {
	// 1. 解析表单数据
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("获取文件失败: "+err.Error(), c)
		return
	}

	fileTypeStr := c.PostForm("type")
	uploaderStr := c.PostForm("uploader")
	groupIDStr := c.PostForm("group_id")

	// 2. 参数验证
	fileType, err := strconv.Atoi(fileTypeStr)
	if err != nil {
		response.FailWithMessage("文件类型参数错误", c)
		return
	}

	uploader, err := strconv.ParseUint(uploaderStr, 10, 64)
	if err != nil {
		response.FailWithMessage("上传者参数错误", c)
		return
	}

	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("文件组参数错误", c)
		return
	}

	// 3. 获取文件组信息
	groupRepo := repositorys.FileGroupRepository{}
	fileGroup, err := groupRepo.FindByID(uint(groupID))
	if err != nil {
		response.FailWithMessage("文件组不存在", c)
		return
	}

	// 4. 创建存储目录
	basePath := "." + global.Config.App.UploadFile
	datePath := time.Now().Format("20060102")
	groupPath := fileGroup.Name // 使用文件组名作为目录名

	fullPath := filepath.Join(basePath, datePath, groupPath)
	if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
		response.FailWithMessage("创建目录失败: "+err.Error(), c)
		return
	}

	// 5. 处理文件名冲突
	originalFilename := file.Filename
	ext := filepath.Ext(originalFilename)
	nameWithoutExt := strings.TrimSuffix(originalFilename, ext)

	newFilename := originalFilename
	counter := 1

	for {
		if _, err := os.Stat(filepath.Join(fullPath, newFilename)); os.IsNotExist(err) {
			break
		}
		newFilename = fmt.Sprintf("%s_%d%s", nameWithoutExt, counter, ext)
		counter++
	}

	// 6. 保存文件
	filePath := filepath.Join(fullPath, newFilename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		response.FailWithMessage("文件保存失败: "+err.Error(), c)
		return
	}

	// 7. 构建文件URL
	relativePath := strings.ReplaceAll(
		strings.TrimPrefix(filePath, "."),
		"\\", "/",
	)
	fileUrl := global.Config.App.ImgUrl + "/static/system/common/file/" + relativePath

	// 8. 保存到数据库
	fileRepo := repositorys.FileRepository{}
	newFile := models.File{
		FileName: newFilename,
		FilePath: relativePath,
		FileUrl:  fileUrl,
		Type:     fileType,
		Uploader: uint(uploader),
		GroupID:  uint(groupID),
	}

	if err := fileRepo.Create(&newFile); err != nil {
		// 如果数据库保存失败，删除已上传的文件
		_ = os.Remove(filePath)
		response.FailWithMessage("文件信息保存失败: "+err.Error(), c)
		return
	}

	// 9. 返回成功响应
	response.OkWithData(gin.H{
		"id":       newFile.ID,
		"fileName": newFilename,
		"filePath": relativePath,
		"fileUrl":  fileUrl,
		"type":     fileType,
		"uploader": uploader,
		"group_id": groupID,
	}, c)
}

func (controller *FileController) Index(ctx *gin.Context) {
	// 获取ID参数
	idStr := ctx.Param("id")
	groupID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}
	// 通过Repository查询数据
	re := repositorys.FileRepository{}
	files, err := re.GetFilesByGroupID(uint(groupID))
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), ctx)
		return
	}

	// 返回查询结果
	response.OkWithData(files, ctx)
}

func (controller *FileController) Edit(ctx *gin.Context) {
	// 绑定请求数据
	var req models.File
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}

	// 检查必要字段
	if req.ID == 0 {
		response.FailWithMessage("文件ID不能为空", ctx)
		return
	}

	// 调用Repository
	re := repositorys.FileRepository{}
	updatedFile, err := re.UpdateFile(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 返回更新后的数据
	response.OkWithData(updatedFile, ctx)
}

// Delete 删除单个文件
func (controller *FileController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	fileID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}
	re := repositorys.FileRepository{}
	if err := re.DeleteFile(uint(fileID)); err != nil {
		response.FailWithMessage("删除失败: "+err.Error(), ctx)
		return
	}

	response.OkWithMessage("删除成功", ctx)
}

// BatchDelete 批量删除文件
func (controller *FileController) BatchDelete(ctx *gin.Context) {
	var request struct {
		IDs []uint `json:"ids" binding:"required"` // 前端传入：[1, 2, 3]
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), ctx)
		return
	}

	if len(request.IDs) == 0 {
		response.FailWithMessage("未选择文件", ctx)
		return
	}
	re := repositorys.FileRepository{}
	if err := re.BatchDeleteFiles(request.IDs); err != nil {
		response.FailWithMessage("批量删除失败: "+err.Error(), ctx)
		return
	}

	response.OkWithMessage(fmt.Sprintf("成功删除 %d 个文件", len(request.IDs)), ctx)
}

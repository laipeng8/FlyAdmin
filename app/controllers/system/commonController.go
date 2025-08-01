package system

import (
	"encoding/json"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"server/global"
	"server/global/response"
	"strconv"
	"time"
)

type CommonController struct {
	allowType      map[string]string
	uploadPath     string
	uploadPathBase string
}

func (p *CommonController) GetFileBasePath() string {
	return "." + global.Config.App.UploadFile
}

func (p *CommonController) UpLoad(c *gin.Context) {

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
		"src":      global.Config.WX.Url + "/api/system/common/file" + allDir[8:],
	})

}

func (p *CommonController) CaptchaInfo(c *gin.Context) {
	id := captcha.NewLen(4)
	response.Success(c, "ok", gin.H{
		"id":  id,
		"url": global.Config.App.Host + "/api/common/captcha/img/" + id,
	})
}

func (p *CommonController) CaptchaImage(c *gin.Context) {
	w, _ := strconv.Atoi(c.Param("w"))
	h, _ := strconv.Atoi(c.Param("h"))
	fmt.Println(c.Param("id"))
	_ = global.CaptchaServe(c.Writer, c.Request, c.Param("id"), ".png", "zh", false, w, h)
}

func (p *CommonController) GetVersion(c *gin.Context) {

	res, err := http.DefaultClient.Get("https://gitee.com/kevn/gsadmin/tags/names.json?search=&page=1")
	if err != nil {
		response.Failed(c, err.Error())
		return
	}

	if res.StatusCode != 200 {
		response.Failed(c, fmt.Sprintf("%s %s", "Http Status Code:", res.Status))
		return
	}

	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()

	versionInfo := make(map[string]interface{}, 1)

	err = json.Unmarshal(all, &versionInfo)
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	if tags, ok := versionInfo["tags"]; ok {
		response.Success(c, "ok", tags)
	} else {
		response.Success(c, "ok", nil)
	}

}

package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	r := Response{
		Code:    200,
		Msg:     msg,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusOK, r)
}

// Failed Deprecated
func Failed(c *gin.Context, err string) {
	r := Response{
		Code:    422,
		Msg:     err,
		Message: err,
		Data:    []string{},
	}
	c.JSON(http.StatusOK, r)
}

type response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListPageResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Succeed = 200
	Error   = 0
)

func Result(code int, data any, MSG string, c *gin.Context) {
	c.JSON(http.StatusOK, response{code, data, MSG})
}

func Ok(data any, Msg string, c *gin.Context) {
	Result(Succeed, data, Msg, c)

}
func OkWithData(data any, c *gin.Context) {
	Result(Succeed, data, "成功", c)

}
func OkWithListPage(List any, count int64, c *gin.Context) {
	OkWithData(ListPageResponse[any]{
		Count: count,
		List:  List,
	}, c)
}

func OkWithMessage(MSG string, c *gin.Context) {
	Result(Succeed, map[string]any{}, MSG, c)

}
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]interface{}{}, msg, c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := GetValidMsg(err, obj)
	FailWithMessage(msg, c)

}

// 查询数据成功
func GetSuccessMsg(c *gin.Context) {
	Result(Succeed, map[string]interface{}{}, "查询数据成功", c)
}

// 插入数据成功
func InsertSuccessMsg(c *gin.Context) {
	Result(Succeed, map[string]interface{}{}, "插入数据成功", c)
}

// 更新数据成功
func UpdateSuccessMsg(c *gin.Context) {
	Result(Succeed, map[string]interface{}{}, "更新数据成功", c)
}

// 删除数据成功
func DeleteSuccessMsg(c *gin.Context) {
	Result(Succeed, map[string]interface{}{}, "删除数据成功", c)
}

// 前端提交数据有误
func FrontDataError(err error, c *gin.Context) {
	Result(401, map[string]interface{}{}, "前端提交数据有误："+err.Error(), c)
}

// 查询数据失败
func GetErrorMsg(err error, c *gin.Context) {
	if err == nil {
		Result(501, map[string]interface{}{}, "查询数据失败", c)
	} else {
		Result(501, map[string]interface{}{}, "查询数据失败："+err.Error(), c)
	}
}

// 查询数据失败
func InsetErrorMsg(err error, c *gin.Context) {
	if err == nil {
		Result(502, map[string]interface{}{}, "插入数据失败", c)
	} else {
		Result(502, map[string]interface{}{}, "插入数据失败："+err.Error(), c)
	}
}

// 更新数据失败
func UpdateErrorMsg(err error, c *gin.Context) {
	if err == nil {
		Result(503, map[string]interface{}{}, "更新数据失败", c)
	} else {
		Result(503, map[string]interface{}{}, "更新数据失败："+err.Error(), c)
	}
}

// 删除数据失败
func DeleteErrorMsg(err error, c *gin.Context) {
	if err == nil {
		Result(504, map[string]any{}, "删除数据失败", c)
	} else {
		Result(504, map[string]any{}, "删除数据失败："+err.Error(), c)
	}
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(int(code), map[string]any{}, "未知错误", c)
}
func OkWith(c *gin.Context) {
	Result(Succeed, map[string]any{}, "成功", c)
}

func GetValidMsg(err error, obj any) string {
	//使用的时候，需要传入obj的指针
	getObj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			//根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg

			}
		}
	}
	return err.Error()
}

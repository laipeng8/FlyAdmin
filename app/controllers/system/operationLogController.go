package system

import (
	"github.com/gin-gonic/gin"
	"server/app/repositorys"
	"server/global"
	"server/global/response"
)

type OperationLogController struct{}

func (o *OperationLogController) List(c *gin.Context) {

	var (
		params       global.List
		operationLog repositorys.OperationLogRepository
	)
	_ = c.ShouldBind(&params)
	operationLog.Where = params.Where

	response.Success(c, "ok", operationLog.List(params.Page, params.PageSize, "created_at"))
}

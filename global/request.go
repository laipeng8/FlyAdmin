package global

import "time"

type List struct {
	Page     int                    `json:"page" form:"page"`         // 当前页码
	PageSize int                    `json:"pageSize" form:"pageSize"` // 每页数量
	Where    map[string]interface{} `json:"where" form:"where"`       // 筛选条件
}

type ListTime struct {
	Page      int                    `json:"page" form:"page"`             // 当前页码
	PageSize  int                    `json:"pageSize" form:"pageSize"`     // 每页数量
	Where     map[string]interface{} `json:"where" form:"where"`           // 筛选条件
	StartTime *time.Time             `json:"start_time" form:"start_time"` // 创建时间范围的开始时间
	EndTime   *time.Time             `json:"end_time" form:"end_time"`     // 创建时间范围的结束时间
}

type Del struct {
	Ids []uint `json:"id"`
}

package requests

// CreateTimerTaskRequest 创建定时任务请求
type CreateTimerTaskRequest struct {
	Name           string `json:"name" binding:"required" validate:"required" label:"任务名称"`
	Description    string `json:"description" validate:"max=500" label:"任务描述"`
	CronExpression string `json:"cron_expression" binding:"required" validate:"required" label:"Cron表达式"`
	TargetURL      string `json:"target_url" binding:"required" validate:"required,url" label:"目标URL"`
	Method         string `json:"method" binding:"required" validate:"required,oneof=GET POST PUT DELETE PATCH" label:"请求方法"`
	Headers        string `json:"headers" validate:"max=2000" label:"请求头"`
	Body           string `json:"body" validate:"max=5000" label:"请求体"`
	Timeout        int    `json:"timeout" validate:"min=1,max=300" label:"超时时间"`
	RetryCount     int    `json:"retry_count" validate:"min=0,max=10" label:"重试次数"`
	RetryInterval  int    `json:"retry_interval" validate:"min=0,max=3600" label:"重试间隔"`
}

// UpdateTimerTaskRequest 更新定时任务请求
type UpdateTimerTaskRequest struct {
	ID             uint   `json:"id" binding:"required" validate:"required" label:"任务ID"`
	Name           string `json:"name" binding:"required" validate:"required" label:"任务名称"`
	Description    string `json:"description" validate:"max=500" label:"任务描述"`
	CronExpression string `json:"cron_expression" binding:"required" validate:"required" label:"Cron表达式"`
	TargetURL      string `json:"target_url" binding:"required" validate:"required,url" label:"目标URL"`
	Method         string `json:"method" binding:"required" validate:"required,oneof=GET POST PUT DELETE PATCH" label:"请求方法"`
	Headers        string `json:"headers" validate:"max=2000" label:"请求头"`
	Body           string `json:"body" validate:"max=5000" label:"请求体"`
	Status         int    `json:"status" validate:"oneof=0 1" label:"状态"`
	Timeout        int    `json:"timeout" validate:"min=1,max=300" label:"超时时间"`
	RetryCount     int    `json:"retry_count" validate:"min=0,max=10" label:"重试次数"`
	RetryInterval  int    `json:"retry_interval" validate:"min=0,max=3600" label:"重试间隔"`
}

// TimerTaskListRequest 定时任务列表请求
type TimerTaskListRequest struct {
	Page     int    `json:"page" validate:"min=1" label:"页码"`
	PageSize int    `json:"page_size" validate:"min=1,max=100" label:"每页数量"`
	Name     string `json:"name" validate:"max=255" label:"任务名称"`
	Status   *int   `json:"status" validate:"omitempty,oneof=0 1" label:"状态"`
}

// TimerTaskLogRequest 定时任务日志请求
type TimerTaskLogRequest struct {
	TaskID   uint `json:"task_id" binding:"required" validate:"required" label:"任务ID"`
	Page     int  `json:"page" validate:"min=1" label:"页码"`
	PageSize int  `json:"page_size" validate:"min=1,max=100" label:"每页数量"`
}

// ExecuteTaskRequest 手动执行任务请求
type ExecuteTaskRequest struct {
	TaskID uint `json:"task_id" binding:"required" validate:"required" label:"任务ID"`
}

// TestTaskRequest 测试任务请求
type TestTaskRequest struct {
	TargetURL string `json:"target_url" binding:"required" validate:"required,url" label:"目标URL"`
	Method    string `json:"method" binding:"required" validate:"required,oneof=GET POST PUT DELETE PATCH" label:"请求方法"`
	Headers   string `json:"headers" validate:"max=2000" label:"请求头"`
	Body      string `json:"body" validate:"max=5000" label:"请求体"`
	Timeout   int    `json:"timeout" validate:"min=1,max=300" label:"超时时间"`
}

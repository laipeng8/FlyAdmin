package system

import (
	"strconv"

	"server/app/models"
	"server/app/repositorys"
	"server/app/requests"
	"server/global"
	"server/global/response"
	"server/pkg/timer"

	"github.com/gin-gonic/gin"
)

type TimerController struct{}

// StartTimer 启动定时任务管理器
func (TimerController) StartTimer(c *gin.Context) {
	taskManager := timer.GetTaskManager()
	if err := taskManager.Start(); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("定时任务管理器启动成功", c)
}

// StopTimer 停止定时任务管理器
func (TimerController) StopTimer(c *gin.Context) {
	taskManager := timer.GetTaskManager()
	if err := taskManager.Stop(); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("定时任务管理器停止成功", c)
}

// GetTimerStatus 获取定时任务管理器状态
func (TimerController) GetTimerStatus(c *gin.Context) {
	taskManager := timer.GetTaskManager()
	status := taskManager.GetTaskStatus()
	response.OkWithData(status, c)
}

// CreateTask 创建定时任务
func (TimerController) CreateTask(c *gin.Context) {
	var req requests.CreateTimerTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 获取当前用户ID
	userID := global.GetUserID(c)

	task := &models.TimerTask{
		Name:           req.Name,
		Description:    req.Description,
		CronExpression: req.CronExpression,
		TargetURL:      req.TargetURL,
		Method:         req.Method,
		Headers:        req.Headers,
		Body:           req.Body,
		Status:         1, // 默认启用
		Timeout:        req.Timeout,
		RetryCount:     req.RetryCount,
		RetryInterval:  req.RetryInterval,
		Creator:        userID,
	}

	repo := &repositorys.TimerRepository{}
	if err := repo.Create(task); err != nil {
		response.FailWithMessage("创建任务失败: "+err.Error(), c)
		return
	}

	// 如果任务管理器已启动，添加任务到调度器
	taskManager := timer.GetTaskManager()
	if err := taskManager.AddTask(task); err != nil {
		global.Logger.Errorf("添加任务到调度器失败: %v", err)
	}

	response.OkWithMessage("定时任务创建成功", c)
}

// UpdateTask 更新定时任务
func (TimerController) UpdateTask(c *gin.Context) {
	var req requests.UpdateTimerTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	repo := &repositorys.TimerRepository{}
	task, err := repo.FindByID(req.ID)
	if err != nil {
		response.FailWithMessage("任务不存在", c)
		return
	}

	// 更新任务信息
	task.Name = req.Name
	task.Description = req.Description
	task.CronExpression = req.CronExpression
	task.TargetURL = req.TargetURL
	task.Method = req.Method
	task.Headers = req.Headers
	task.Body = req.Body
	task.Status = req.Status
	task.Timeout = req.Timeout
	task.RetryCount = req.RetryCount
	task.RetryInterval = req.RetryInterval

	if err := repo.Update(task); err != nil {
		response.FailWithMessage("更新任务失败: "+err.Error(), c)
		return
	}

	// 更新任务管理器中的任务
	taskManager := timer.GetTaskManager()
	if err := taskManager.UpdateTask(task); err != nil {
		global.Logger.Errorf("更新任务管理器失败: %v", err)
	}

	response.OkWithMessage("定时任务更新成功", c)
}

// DeleteTask 删除定时任务
func (TimerController) DeleteTask(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("任务ID格式错误", c)
		return
	}

	repo := &repositorys.TimerRepository{}
	if err := repo.Delete(uint(taskID)); err != nil {
		response.FailWithMessage("删除任务失败: "+err.Error(), c)
		return
	}

	// 从任务管理器中移除任务
	taskManager := timer.GetTaskManager()
	if err := taskManager.RemoveTask(uint(taskID)); err != nil {
		global.Logger.Errorf("从任务管理器移除任务失败: %v", err)
	}

	response.OkWithMessage("定时任务删除成功", c)
}

// GetTask 获取定时任务详情
func (TimerController) GetTask(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("任务ID格式错误", c)
		return
	}

	repo := &repositorys.TimerRepository{}
	task, err := repo.FindByID(uint(taskID))
	if err != nil {
		response.FailWithMessage("任务不存在", c)
		return
	}

	response.OkWithData(task, c)
}

// GetTaskList 获取定时任务列表
func (TimerController) GetTaskList(c *gin.Context) {
	var req requests.TimerTaskListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 构建查询条件
	query := global.Db.Model(&models.TimerTask{})
	if req.Name != "" {
		query = query.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	// 统计总数
	var total int64
	query.Count(&total)

	// 分页查询
	var tasks []models.TimerTask
	offset := (req.Page - 1) * req.PageSize
	err := query.Offset(offset).Limit(req.PageSize).Order("created_at desc").Find(&tasks).Error
	if err != nil {
		response.FailWithMessage("查询任务列表失败: "+err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"list":  tasks,
		"total": total,
		"page":  req.Page,
		"size":  req.PageSize,
	}, c)
}

// ExecuteTask 手动执行任务
func (TimerController) ExecuteTask(c *gin.Context) {
	var req requests.ExecuteTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	taskManager := timer.GetTaskManager()
	if err := taskManager.ExecuteTask(req.TaskID); err != nil {
		response.FailWithMessage("执行任务失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("任务执行成功", c)
}

// TestTask 测试任务
func (TimerController) TestTask(c *gin.Context) {
	var req requests.TestTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	taskManager := timer.GetTaskManager()
	result, err := taskManager.TestTask(req.TargetURL, req.Method, req.Headers, req.Body, req.Timeout)
	if err != nil {
		response.FailWithMessage("测试任务失败: "+err.Error(), c)
		return
	}

	response.OkWithData(result, c)
}

// GetTaskLogs 获取任务执行日志
func (TimerController) GetTaskLogs(c *gin.Context) {
	var req requests.TimerTaskLogRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	// 统计总数
	var total int64
	global.Db.Model(&models.TimerTaskLog{}).Where("task_id = ?", req.TaskID).Count(&total)

	// 分页查询
	var logs []models.TimerTaskLog
	offset := (req.Page - 1) * req.PageSize
	err := global.Db.Where("task_id = ?", req.TaskID).
		Order("created_at desc").
		Offset(offset).
		Limit(req.PageSize).
		Find(&logs).Error

	if err != nil {
		response.FailWithMessage("查询任务日志失败: "+err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"list":  logs,
		"total": total,
		"page":  req.Page,
		"size":  req.PageSize,
	}, c)
}

// ToggleTaskStatus 切换任务状态
func (TimerController) ToggleTaskStatus(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("任务ID格式错误", c)
		return
	}

	repo := &repositorys.TimerRepository{}
	task, err := repo.FindByID(uint(taskID))
	if err != nil {
		response.FailWithMessage("任务不存在", c)
		return
	}

	// 切换状态
	if task.Status == 1 {
		task.Status = 0
	} else {
		task.Status = 1
	}

	if err := repo.Update(task); err != nil {
		response.FailWithMessage("更新任务状态失败: "+err.Error(), c)
		return
	}

	// 更新任务管理器
	taskManager := timer.GetTaskManager()
	if err := taskManager.UpdateTask(task); err != nil {
		global.Logger.Errorf("更新任务管理器失败: %v", err)
	}

	statusText := "启用"
	if task.Status == 0 {
		statusText = "禁用"
	}

	response.OkWithMessage("任务已"+statusText, c)
}

// GetCronExamples 获取Cron表达式示例
func (TimerController) GetCronExamples(c *gin.Context) {
	examples := []map[string]string{
		{"name": "每分钟执行", "expression": "0 * * * * *", "description": "每分钟的第0秒执行"},
		{"name": "每小时执行", "expression": "0 0 * * * *", "description": "每小时的第0分0秒执行"},
		{"name": "每天凌晨执行", "expression": "0 0 0 * * *", "description": "每天凌晨0点0分0秒执行"},
		{"name": "每周一执行", "expression": "0 0 0 * * 1", "description": "每周一凌晨0点0分0秒执行"},
		{"name": "每月1号执行", "expression": "0 0 0 1 * *", "description": "每月1号凌晨0点0分0秒执行"},
		{"name": "每5分钟执行", "expression": "0 */5 * * * *", "description": "每5分钟执行一次"},
		{"name": "每天上午9点和下午6点执行", "expression": "0 0 9,18 * * *", "description": "每天上午9点和下午6点执行"},
		{"name": "工作日执行", "expression": "0 0 9 * * 1-5", "description": "周一到周五上午9点执行"},
	}

	response.OkWithData(examples, c)
}

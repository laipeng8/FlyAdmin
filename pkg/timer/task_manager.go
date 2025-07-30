package timer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"server/app/models"
	"server/app/repositorys"
	"server/global"

	"github.com/robfig/cron/v3"
)

// TaskManager 定时任务管理器
type TaskManager struct {
	cron     *cron.Cron
	tasks    map[uint]cron.EntryID // 任务ID -> Cron Entry ID
	taskMu   sync.RWMutex
	repo     *repositorys.TimerRepository
	running  bool
	stopChan chan struct{}
}

var (
	manager *TaskManager
	once    sync.Once
)

// GetTaskManager 获取任务管理器单例
func GetTaskManager() *TaskManager {
	once.Do(func() {
		manager = &TaskManager{
			cron:     cron.New(cron.WithSeconds()),
			tasks:    make(map[uint]cron.EntryID),
			repo:     &repositorys.TimerRepository{},
			stopChan: make(chan struct{}),
		}
	})
	return manager
}

// Start 启动任务管理器
func (tm *TaskManager) Start() error {
	tm.taskMu.Lock()
	defer tm.taskMu.Unlock()

	if tm.running {
		return fmt.Errorf("任务管理器已启动")
	}

	// 加载所有启用的任务
	tasks, err := tm.repo.FindEnabled()
	if err != nil {
		return fmt.Errorf("加载任务失败: %w", err)
	}

	// 添加所有任务到cron
	for _, task := range tasks {
		if err := tm.addTask(&task); err != nil {
			global.Logger.Errorf("添加任务失败 [%d]: %v", task.ID, err)
			continue
		}
	}

	tm.cron.Start()
	tm.running = true
	global.Logger.Info("定时任务管理器启动成功")
	return nil
}

// Stop 停止任务管理器
func (tm *TaskManager) Stop() error {
	tm.taskMu.Lock()
	defer tm.taskMu.Unlock()

	if !tm.running {
		return fmt.Errorf("任务管理器未启动")
	}

	tm.cron.Stop()
	tm.running = false
	close(tm.stopChan)
	global.Logger.Info("定时任务管理器已停止")
	return nil
}

// AddTask 添加新任务
func (tm *TaskManager) AddTask(task *models.TimerTask) error {
	tm.taskMu.Lock()
	defer tm.taskMu.Unlock()

	if !tm.running {
		return fmt.Errorf("任务管理器未启动")
	}

	return tm.addTask(task)
}

// UpdateTask 更新任务
func (tm *TaskManager) UpdateTask(task *models.TimerTask) error {
	tm.taskMu.Lock()
	defer tm.taskMu.Unlock()

	if !tm.running {
		return fmt.Errorf("任务管理器未启动")
	}

	// 移除旧任务
	if entryID, exists := tm.tasks[task.ID]; exists {
		tm.cron.Remove(entryID)
		delete(tm.tasks, task.ID)
	}

	// 如果任务启用，添加新任务
	if task.Status == 1 {
		return tm.addTask(task)
	}

	return nil
}

// RemoveTask 移除任务
func (tm *TaskManager) RemoveTask(taskID uint) error {
	tm.taskMu.Lock()
	defer tm.taskMu.Unlock()

	if !tm.running {
		return fmt.Errorf("任务管理器未启动")
	}

	if entryID, exists := tm.tasks[taskID]; exists {
		tm.cron.Remove(entryID)
		delete(tm.tasks, taskID)
	}

	return nil
}

// ExecuteTask 手动执行任务
func (tm *TaskManager) ExecuteTask(taskID uint) error {
	task, err := tm.repo.FindByID(taskID)
	if err != nil {
		return fmt.Errorf("任务不存在: %w", err)
	}

	return tm.executeTask(task)
}

// TestTask 测试任务
func (tm *TaskManager) TestTask(targetURL, method, headers, body string, timeout int) (map[string]interface{}, error) {
	startTime := time.Now()

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	// 创建请求
	var req *http.Request
	var err error

	if method == "GET" {
		req, err = http.NewRequest(method, targetURL, nil)
	} else {
		req, err = http.NewRequest(method, targetURL, strings.NewReader(body))
	}

	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	if headers != "" {
		var headerMap map[string]string
		if err := json.Unmarshal([]byte(headers), &headerMap); err == nil {
			for key, value := range headerMap {
				req.Header.Set(key, value)
			}
		}
	}

	// 设置默认Content-Type
	if method != "GET" && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求执行失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	duration := time.Since(startTime).Milliseconds()

	return map[string]interface{}{
		"status_code": resp.StatusCode,
		"duration":    duration,
		"response":    string(respBody),
		"headers":     resp.Header,
	}, nil
}

// GetTaskStatus 获取任务状态
func (tm *TaskManager) GetTaskStatus() map[string]interface{} {
	tm.taskMu.RLock()
	defer tm.taskMu.RUnlock()

	entries := tm.cron.Entries()
	taskStatus := make(map[string]interface{})

	for _, entry := range entries {
		taskStatus[fmt.Sprintf("entry_%d", entry.ID)] = map[string]interface{}{
			"next_run": entry.Next,
			"prev_run": entry.Prev,
		}
	}

	return map[string]interface{}{
		"running":    tm.running,
		"task_count": len(tm.tasks),
		"entries":    taskStatus,
	}
}

// addTask 内部方法：添加任务到cron
func (tm *TaskManager) addTask(task *models.TimerTask) error {
	entryID, err := tm.cron.AddFunc(task.CronExpression, func() {
		tm.executeTask(task)
	})
	if err != nil {
		return fmt.Errorf("添加cron任务失败: %w", err)
	}

	tm.tasks[task.ID] = entryID
	return nil
}

// executeTask 执行任务
func (tm *TaskManager) executeTask(task *models.TimerTask) error {
	startTime := time.Now()

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: time.Duration(task.Timeout) * time.Second,
	}

	// 创建请求
	var req *http.Request
	var err error

	if task.Method == "GET" {
		req, err = http.NewRequest(task.Method, task.TargetURL, nil)
	} else {
		req, err = http.NewRequest(task.Method, task.TargetURL, strings.NewReader(task.Body))
	}

	if err != nil {
		return tm.handleTaskError(task, err, startTime)
	}

	// 设置请求头
	if task.Headers != "" {
		var headerMap map[string]string
		if err := json.Unmarshal([]byte(task.Headers), &headerMap); err == nil {
			for key, value := range headerMap {
				req.Header.Set(key, value)
			}
		}
	}

	// 设置默认Content-Type
	if task.Method != "GET" && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return tm.handleTaskError(task, err, startTime)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return tm.handleTaskError(task, err, startTime)
	}

	duration := time.Since(startTime).Milliseconds()

	// 判断是否成功
	success := resp.StatusCode >= 200 && resp.StatusCode < 300

	// 记录执行日志
	log := &models.TimerTaskLog{
		TaskID:     task.ID,
		Status:     map[bool]int{true: 1, false: 0}[success],
		Message:    fmt.Sprintf("HTTP %d", resp.StatusCode),
		Duration:   duration,
		Response:   string(respBody),
		StatusCode: resp.StatusCode,
		RunTime:    time.Now(),
	}

	if err := tm.repo.CreateLog(log); err != nil {
		global.Logger.Errorf("记录任务日志失败 [%d]: %v", task.ID, err)
	}

	// 更新任务统计
	task.RunCount++
	if success {
		task.SuccessCount++
	} else {
		task.FailCount++
	}
	task.LastRunTime = time.Now()

	// 计算下次执行时间
	if entryID, exists := tm.tasks[task.ID]; exists {
		entries := tm.cron.Entries()
		for _, entry := range entries {
			if entry.ID == entryID {
				task.NextRunTime = entry.Next
				break
			}
		}
	}

	if err := tm.repo.UpdateRunStats(task.ID, task.RunCount, task.SuccessCount, task.FailCount); err != nil {
		global.Logger.Errorf("更新任务统计失败 [%d]: %v", task.ID, err)
	}

	if err := tm.repo.UpdateLastRunTime(task.ID, task.LastRunTime, task.NextRunTime); err != nil {
		global.Logger.Errorf("更新任务时间失败 [%d]: %v", task.ID, err)
	}

	// 如果失败且需要重试
	if !success && task.RetryCount > 0 {
		go tm.retryTask(task, 1)
	}

	global.Logger.Infof("任务执行完成 [%d]: %s, 耗时: %dms, 状态: %d",
		task.ID, task.Name, duration, resp.StatusCode)

	return nil
}

// handleTaskError 处理任务执行错误
func (tm *TaskManager) handleTaskError(task *models.TimerTask, err error, startTime time.Time) error {
	duration := time.Since(startTime).Milliseconds()

	// 记录错误日志
	log := &models.TimerTaskLog{
		TaskID:     task.ID,
		Status:     0,
		Message:    err.Error(),
		Duration:   duration,
		Response:   "",
		StatusCode: 0,
		RunTime:    time.Now(),
	}

	if logErr := tm.repo.CreateLog(log); logErr != nil {
		global.Logger.Errorf("记录任务错误日志失败 [%d]: %v", task.ID, logErr)
	}

	// 更新任务统计
	task.RunCount++
	task.FailCount++
	task.LastRunTime = time.Now()

	if updateErr := tm.repo.UpdateRunStats(task.ID, task.RunCount, task.SuccessCount, task.FailCount); updateErr != nil {
		global.Logger.Errorf("更新任务统计失败 [%d]: %v", task.ID, updateErr)
	}

	// 如果失败且需要重试
	if task.RetryCount > 0 {
		go tm.retryTask(task, 1)
	}

	global.Logger.Errorf("任务执行失败 [%d]: %s, 错误: %v", task.ID, task.Name, err)
	return err
}

// retryTask 重试任务
func (tm *TaskManager) retryTask(task *models.TimerTask, retryCount int) {
	if retryCount > task.RetryCount {
		return
	}

	// 等待重试间隔
	time.Sleep(time.Duration(task.RetryInterval) * time.Second)

	// 重新执行任务
	if err := tm.executeTask(task); err != nil {
		global.Logger.Errorf("任务重试失败 [%d]: 第%d次重试, 错误: %v", task.ID, retryCount, err)
		// 继续重试
		go tm.retryTask(task, retryCount+1)
	} else {
		global.Logger.Infof("任务重试成功 [%d]: 第%d次重试", task.ID, retryCount)
	}
}

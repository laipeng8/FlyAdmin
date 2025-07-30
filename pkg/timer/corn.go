package timer

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"server/config"
	"strings"
	"sync"
)

var (
	cronInstance          *cron.Cron
	cronMu                sync.Mutex
	cronRunning           bool
	cronEntryID           cron.EntryID // 用于存储当前任务的 ID
	currentCronExpression string       // 存储当前的定时任务时间规则
)

// StartCronJob 启动定时任务
func StartCronJob(cfg *config.Config) error {
	cronMu.Lock()
	defer cronMu.Unlock()

	if cronRunning {
		return fmt.Errorf("定时任务已启动")
	}

	cronInstance = cron.New()

	// 添加订单状态更新任务
	entryID, err := cronInstance.AddFunc(cfg.Cron.OrderStatusUpdate, func() {
		fmt.Println("Running order status upload task...")
		UpdateOrderStatus()
	})
	if err != nil {
		return fmt.Errorf("添加定时任务失败: %w", err)
	}

	cronEntryID = entryID                              // 保存任务 ID
	currentCronExpression = cfg.Cron.OrderStatusUpdate // 保存当前的时间规则
	cronInstance.Start()
	cronRunning = true
	return nil
}

// StopCronJob 停止定时任务
func StopCronJob() error {
	cronMu.Lock()
	defer cronMu.Unlock()

	if !cronRunning {
		return fmt.Errorf("定时任务未启动")
	}

	cronInstance.Stop()
	cronRunning = false
	currentCronExpression = "" // 清空时间规则
	return nil
}

// UpdateCronJob 更新定时任务的时间规则
func UpdateCronJob(cronExpression string) error {
	cronMu.Lock()
	defer cronMu.Unlock()

	if !cronRunning {
		return fmt.Errorf("定时任务未启动")
	}

	// 移除旧的任务
	cronInstance.Remove(cronEntryID)

	// 添加新的任务
	entryID, err := cronInstance.AddFunc(cronExpression, func() {
		fmt.Println("Running order status upload task with new schedule...")
		UpdateOrderStatus()
	})
	if err != nil {
		return fmt.Errorf("更新定时任务失败: %w", err)
	}

	cronEntryID = entryID                  // 更新任务 ID
	currentCronExpression = cronExpression // 更新当前的时间规则
	return nil
}

// CronJobStatus 获取定时任务状态和时间规则描述
func CronJobStatus() (string, string, error) {
	cronMu.Lock()
	defer cronMu.Unlock()

	if !cronRunning {
		return "stopped", "", nil
	}

	// 解析 Cron 表达式为自然语言描述
	description, err := ParseCronExpression(currentCronExpression)
	if err != nil {
		return "running", "", err
	}

	return "running", description, nil
}

// ParseCronExpression 将 Cron 表达式转换为自然语言描述
func ParseCronExpression(cronExpr string) (string, error) {
	// 解析 Cron 表达式
	_, err := cron.ParseStandard(cronExpr)
	if err != nil {
		return "", fmt.Errorf("无效的 Cron 表达式: %v", err)
	}

	// 获取 Cron 表达式的各个字段
	fields := strings.Fields(cronExpr)
	if len(fields) != 5 {
		return "", fmt.Errorf("Cron 表达式必须包含 5 个字段")
	}

	// 解析分钟和小时7
	minute := fields[0]
	hour := fields[1]

	// 转换为自然语言描述
	var description string
	switch {
	case minute == "0" && hour != "*":
		description = fmt.Sprintf("每天 %s 点", hour)
	case minute != "0" && hour != "*":
		description = fmt.Sprintf("每天 %s 点 %s 分", hour, minute)
	case strings.HasPrefix(minute, "*/"):
		interval := strings.TrimPrefix(minute, "*/")
		description = fmt.Sprintf("每隔 %s 分钟", interval)
	default:
		description = "自定义时间规则"
	}

	return description, nil
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus() {
	// 实现订单状态更新逻辑
	fmt.Println("Updating order status...")
}

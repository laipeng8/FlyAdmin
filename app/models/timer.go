package models

import (
	"fmt"
	"server/global"
	"time"
)

// TimerTask 定时任务模型
type TimerTask struct {
	global.GAD_MODEL
	Name           string    `gorm:"type:varchar(255);column:name;comment:任务名称" json:"name"`
	Description    string    `gorm:"type:varchar(500);column:description;comment:任务描述" json:"description"`
	CronExpression string    `gorm:"type:varchar(100);column:cron_expression;comment:cron表达式" json:"cron_expression"`
	TargetURL      string    `gorm:"type:varchar(500);column:target_url;comment:目标接口URL" json:"target_url"`
	Method         string    `gorm:"type:varchar(10);column:method;comment:请求方法" json:"method"`
	Headers        string    `gorm:"type:text;column:headers;comment:请求头" json:"headers"`
	Body           string    `gorm:"type:text;column:body;comment:请求体" json:"body"`
	Status         int       `gorm:"type:tinyint;column:status;default:1;comment:状态 1:启用 0:禁用" json:"status"`
	LastRunTime    time.Time `gorm:"column:last_run_time;comment:最后执行时间" json:"last_run_time"`
	NextRunTime    time.Time `gorm:"column:next_run_time;comment:下次执行时间" json:"next_run_time"`
	RunCount       int64     `gorm:"column:run_count;default:0;comment:执行次数" json:"run_count"`
	SuccessCount   int64     `gorm:"column:success_count;default:0;comment:成功次数" json:"success_count"`
	FailCount      int64     `gorm:"column:fail_count;default:0;comment:失败次数" json:"fail_count"`
	Timeout        int       `gorm:"column:timeout;default:30;comment:超时时间(秒)" json:"timeout"`
	RetryCount     int       `gorm:"column:retry_count;default:0;comment:重试次数" json:"retry_count"`
	RetryInterval  int       `gorm:"column:retry_interval;default:60;comment:重试间隔(秒)" json:"retry_interval"`
	Creator        uint      `gorm:"column:creator;comment:创建者" json:"creator"`
}

// TimerTaskLog 定时任务执行日志
type TimerTaskLog struct {
	global.GAD_MODEL
	TaskID     uint      `gorm:"column:task_id;comment:任务ID" json:"task_id"`
	Status     int       `gorm:"type:tinyint;column:status;comment:执行状态 1:成功 0:失败" json:"status"`
	Message    string    `gorm:"type:text;column:message;comment:执行结果" json:"message"`
	Duration   int64     `gorm:"column:duration;comment:执行时长(毫秒)" json:"duration"`
	Response   string    `gorm:"type:text;column:response;comment:响应内容" json:"response"`
	StatusCode int       `gorm:"column:status_code;comment:HTTP状态码" json:"status_code"`
	RunTime    time.Time `gorm:"column:run_time;comment:执行时间" json:"run_time"`
}

func (m *TimerTask) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "timer_tasks")
}

func (m *TimerTaskLog) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "timer_task_logs")
}

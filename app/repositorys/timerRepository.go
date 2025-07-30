package repositorys

import (
	"server/app/models"
	"server/global"
)

type TimerRepository struct {
	BaseRepository
}

// Create 创建定时任务
func (r *TimerRepository) Create(task *models.TimerTask) error {
	return global.Db.Create(task).Error
}

// Update 更新定时任务
func (r *TimerRepository) Update(task *models.TimerTask) error {
	return global.Db.Save(task).Error
}

// Delete 删除定时任务
func (r *TimerRepository) Delete(id uint) error {
	return global.Db.Delete(&models.TimerTask{}, id).Error
}

// FindByID 根据ID查找定时任务
func (r *TimerRepository) FindByID(id uint) (*models.TimerTask, error) {
	var task models.TimerTask
	err := global.Db.First(&task, id).Error
	return &task, err
}

// FindAll 查找所有定时任务
func (r *TimerRepository) FindAll() ([]models.TimerTask, error) {
	var tasks []models.TimerTask
	err := global.Db.Find(&tasks).Error
	return tasks, err
}

// FindEnabled 查找所有启用的定时任务
func (r *TimerRepository) FindEnabled() ([]models.TimerTask, error) {
	var tasks []models.TimerTask
	err := global.Db.Where("status = ?", 1).Find(&tasks).Error
	return tasks, err
}

// UpdateRunStats 更新任务执行统计
func (r *TimerRepository) UpdateRunStats(id uint, runCount, successCount, failCount int64) error {
	return global.Db.Model(&models.TimerTask{}).Where("id = ?", id).Updates(map[string]interface{}{
		"run_count":     runCount,
		"success_count": successCount,
		"fail_count":    failCount,
	}).Error
}

// UpdateLastRunTime 更新最后执行时间
func (r *TimerRepository) UpdateLastRunTime(id uint, lastRunTime, nextRunTime interface{}) error {
	return global.Db.Model(&models.TimerTask{}).Where("id = ?", id).Updates(map[string]interface{}{
		"last_run_time": lastRunTime,
		"next_run_time": nextRunTime,
	}).Error
}

// CreateLog 创建执行日志
func (r *TimerRepository) CreateLog(log *models.TimerTaskLog) error {
	return global.Db.Create(log).Error
}

// FindLogsByTaskID 根据任务ID查找执行日志
func (r *TimerRepository) FindLogsByTaskID(taskID uint, limit int) ([]models.TimerTaskLog, error) {
	var logs []models.TimerTaskLog
	err := global.Db.Where("task_id = ?", taskID).Order("created_at desc").Limit(limit).Find(&logs).Error
	return logs, err
}

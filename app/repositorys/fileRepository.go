package repositorys

import (
	"fmt"
	"os"
	"server/app/models"
	"server/global"
)

type FileRepository struct {
	FileGroupModel models.File
	BaseRepository
}

// FileRepository 中添加
func (repo *FileRepository) Create(file *models.File) error {
	return global.Db.Create(file).Error
}

// GetFilesByGroupID 根据分组ID获取文件列表
func (repo *FileRepository) GetFilesByGroupID(groupID uint) ([]models.File, error) {
	var files []models.File
	err := global.Db.Where("group_id = ?", groupID).Find(&files).Error
	return files, err
}

// DeleteFile 删除单个文件（根据 file.id）
func (repo *FileRepository) DeleteFile(id uint) error {
	var file models.File
	if err := global.Db.First(&file, id).Error; err != nil {
		return err // 文件不存在
	}

	// 删除数据库记录 + 实际文件（事务保证一致性）
	tx := global.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Delete(&file).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 可选：删除实际文件
	if file.FilePath != "" {
		if err := os.Remove(file.FilePath); err != nil && !os.IsNotExist(err) {
			tx.Rollback()
			return fmt.Errorf("文件删除失败: %v", err)
		}
	}

	return tx.Commit().Error
}

// BatchDeleteFiles 批量删除文件（根据 file.id 列表）
func (repo *FileRepository) BatchDeleteFiles(ids []uint) error {
	// 先查询所有文件记录（获取文件路径，用于物理删除）
	var files []models.File
	if err := global.Db.Where("id IN ?", ids).Find(&files).Error; err != nil {
		return err
	}

	// 开启事务
	tx := global.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 删除数据库记录
	if err := tx.Where("id IN ?", ids).Delete(&models.File{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 2. 删除实际文件（可选）
	for _, file := range files {
		if file.FilePath != "" {
			if err := os.Remove(file.FilePath); err != nil && !os.IsNotExist(err) {
				tx.Rollback()
				return fmt.Errorf("文件删除失败: %v", err)
			}
		}
	}

	return tx.Commit().Error
}

// UpdateFile 更新文件信息（全字段更新）
func (repo *FileRepository) UpdateFile(file *models.File) (*models.File, error) {
	// 直接更新所有字段（GORM会自动忽略零值字段如created_at）
	result := global.Db.Model(&models.File{}).Where("id = ?", file.ID).Updates(file)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("文件不存在或未修改")
	}

	// 返回更新后的完整数据
	var updatedFile models.File
	global.Db.First(&updatedFile, file.ID)
	return &updatedFile, nil
}

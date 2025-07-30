package repositorys

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"server/app/models"
	"server/app/requests"
	"server/global"
)

type FileGroupRepository struct {
	FileGroupModel models.FileGroup
	BaseRepository
}

// FindByID 根据ID查询单个文件组
func (repo *FileGroupRepository) FindByID(id uint) (*models.FileGroup, error) {
	var fileGroup models.FileGroup
	if err := global.Db.First(&fileGroup, id).Error; err != nil {
		return nil, err
	}
	return &fileGroup, nil
}

func (repo *FileGroupRepository) FindAll() ([]models.FileGroup, error) {
	var fileGroups []models.FileGroup
	err := global.Db.Find(&fileGroups).Error
	if err != nil {
		return nil, err
	}
	return fileGroups, nil
}

func (repo *FileGroupRepository) BuildTree(groups []models.FileGroup, parentID uint) []*models.TreeFileGroup {
	var tree []*models.TreeFileGroup

	for _, group := range groups {
		if group.ParentID == parentID {
			node := &models.TreeFileGroup{
				FileGroup: group,
				Children:  repo.BuildTree(groups, group.ID),
			}
			tree = append(tree, node)
		}
	}

	return tree
}

// GetTreeByRootID 根据根节点ID获取子树
func (repo *FileGroupRepository) GetTreeByRootID(rootID uint) (*models.TreeFileGroup, error) {
	// 1. 查询所有文件组
	allGroups, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	// 2. 查找根节点
	var rootGroup *models.FileGroup
	for _, group := range allGroups {
		if group.ID == rootID {
			rootGroup = &group
			break
		}
	}
	if rootGroup == nil {
		return nil, nil
	}

	// 3. 构建子树
	rootTree := &models.TreeFileGroup{
		FileGroup: *rootGroup,
		Children:  repo.BuildTree(allGroups, rootID),
	}

	return rootTree, nil
}

// Create 创建文件组
func (repo *FileGroupRepository) Create(fileGroup *requests.FileGroupAdd) error {
	repo.FileGroupModel.Name = fileGroup.Name
	repo.FileGroupModel.ParentID = fileGroup.ParentID
	if err := global.Db.Create(&repo.FileGroupModel).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新文件组
func (repo *FileGroupRepository) Update(data requests.FileGroupUpdate) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		// 1. 检查文件组是否存在
		var existingGroup models.FileGroup
		if err := tx.First(&existingGroup, data.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("文件组不存在")
			}
			return fmt.Errorf("查询文件组失败: %w", err)
		}

		// 2. 检查是否尝试将父级设为自己的子级(防止循环引用)
		if data.ParentID != 0 {
			// 2.1 不能设置自己为父级
			if data.ParentID == data.ID {
				return fmt.Errorf("不能将父级设置为自己")
			}

			// 2.2 检查新父级是否是自己的子级
			if err := checkIsChild(tx, data.ID, data.ParentID); err != nil {
				return err
			}
		}

		// 3. 执行更新
		updates := map[string]interface{}{
			"name":      data.Name,
			"parent_id": data.ParentID,
		}

		if err := tx.Model(&models.FileGroup{}).
			Where("id = ?", data.ID).
			Updates(updates).Error; err != nil {
			return fmt.Errorf("更新文件组失败: %w", err)
		}

		return nil
	})
}

// checkIsChild 检查targetID是否是parentID的子级
func checkIsChild(tx *gorm.DB, parentID, targetID uint) error {
	var children []models.FileGroup
	if err := tx.Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return fmt.Errorf("查询子文件组失败: %w", err)
	}

	for _, child := range children {
		if child.ID == targetID {
			return fmt.Errorf("不能将父级设置为自己的子级")
		}
		if err := checkIsChild(tx, child.ID, targetID); err != nil {
			return err
		}
	}
	return nil
}

// ForceDelete 强制删除文件组及其子文件组和关联文件
func (repo *FileGroupRepository) ForceDelete(id uint) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		// 1. 递归获取所有子文件组ID（包括自身）
		allGroupIDs, err := repo.getAllChildGroupIDs(tx, id)
		if err != nil {
			return fmt.Errorf("获取子文件组失败: %w", err)
		}
		allGroupIDs = append(allGroupIDs, id) // 包含自身

		// 2. 删除所有关联文件
		if err := tx.Where("`group_id` IN (?)", allGroupIDs).
			Delete(&models.File{}).Error; err != nil {
			return fmt.Errorf("删除关联文件失败: %w", err)
		}

		// 3. 删除所有文件组（从叶子节点开始删除）
		if err := repo.deleteGroupsRecursively(tx, id); err != nil {
			return fmt.Errorf("删除文件组失败: %w", err)
		}

		return nil
	})
}

// getAllChildGroupIDs 递归获取所有子文件组ID
func (repo *FileGroupRepository) getAllChildGroupIDs(tx *gorm.DB, parentID uint) ([]uint, error) {
	var childIDs []uint
	var children []models.FileGroup

	// 获取直接子文件组
	if err := tx.Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return nil, err
	}

	// 递归获取所有子级
	for _, child := range children {
		ids, err := repo.getAllChildGroupIDs(tx, child.ID)
		if err != nil {
			return nil, err
		}
		childIDs = append(childIDs, child.ID)
		childIDs = append(childIDs, ids...)
	}

	return childIDs, nil
}

// deleteGroupsRecursively 递归删除文件组（从叶子节点开始）
func (repo *FileGroupRepository) deleteGroupsRecursively(tx *gorm.DB, parentID uint) error {
	var children []models.FileGroup

	// 1. 先处理所有子文件组
	if err := tx.Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return err
	}

	for _, child := range children {
		if err := repo.deleteGroupsRecursively(tx, child.ID); err != nil {
			return err
		}
	}

	// 2. 最后删除当前文件组
	if err := tx.Delete(&models.FileGroup{}, parentID).Error; err != nil {
		return err
	}

	return nil
}

// Exists 检查文件组是否存在
func (repo *FileGroupRepository) Exists(id uint) (bool, error) {
	var count int64
	if err := global.Db.Model(&models.FileGroup{}).
		Where("id = ?", id).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetChildren 获取子文件组列表
func (repo *FileGroupRepository) GetChildren(parentID uint) ([]models.FileGroup, error) {
	var children []models.FileGroup
	if err := global.Db.Where("parent_id = ?", parentID).
		Find(&children).Error; err != nil {
		return nil, err
	}
	return children, nil
}

// GetFiles 获取关联文件列表
func (repo *FileGroupRepository) GetFiles(groupID uint) ([]models.File, error) {
	var files []models.File
	if err := global.Db.Where("`group_id` = ?", groupID).
		Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

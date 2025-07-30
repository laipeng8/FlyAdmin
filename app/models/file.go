package models

import (
	"fmt"
	"server/global"
)

type File struct {
	global.GAD_MODEL
	FileName string `gorm:"type:varchar(255);column:file_name;comment:文件名" json:"file_name"`
	FilePath string `gorm:"type:varchar(255);column:file_path;comment:文件路径" json:"file_path"`
	FileUrl  string `gorm:"type:varchar(255);column:file_url;comment:文件网络链接" json:"file_url"`
	Type     int    `gorm:"type:int;column:type;comment:1图片，2视频，3html" json:"type"`
	Uploader uint   `gorm:"type:bigint;column:uploader;comment:上传者" json:"uploader"`
	GroupID  uint   `gorm:"type:bigint;column:group_id;comment:分组_id" json:"group_id"`
}

type FileGroup struct {
	global.GAD_MODEL
	Name     string `gorm:"type:varchar(255);column:name;comment:文件组名" json:"name"`
	ParentID uint   `gorm:"type:bigint;column:parent_id;comment:文件组上级id" json:"parent_id"`
}

type TreeFileGroup struct {
	FileGroup
	Children []*TreeFileGroup `json:"children"`
}

func (m *File) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "files")
}

func (m *FileGroup) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "file_group")
}

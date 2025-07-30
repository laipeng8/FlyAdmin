package models

import (
	"server/global"
)

type MenuApiList struct {
	global.GAD_MODEL
	Code   string    `gorm:"column:code;type:varchar(100);comment:关键字" json:"code"`
	Url    string    `gorm:"column:url;type:varchar(100);comment:地址" json:"url"`
	MenuId uint      `gorm:"column:menu_id;type:int;" json:"menu_id"`
	Describe string `gorm:"type:varchar(255);column:describe;comment:描述" json:"describe"`
	Menu   AdminMenu `gorm:"foreignKey:id;references:menu_id"`
}

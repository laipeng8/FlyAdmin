package models

import (
	"fmt"
	"server/global"
)

type Department struct {
	global.GAD_MODEL
	Name      string      `gorm:"column:name;type:varchar(255)" json:"name"`
	Status    *int        `gorm:"column:status;type:int" json:"status"`
	Sort      int         `gorm:"column:sort;type:int" json:"sort"`
	AdminUser []AdminUser `gorm:"many2many:user_department;" json:"admin_users"`
}

func (m *Department) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "department")
}

type UserDepartment struct {
	AdminUserID  uint `gorm:"primaryKey;column:admin_user_id"`
	DepartmentID uint `gorm:"primaryKey;column:department_id"`
}

func (m *UserDepartment) TableName() string {
	return fmt.Sprint(global.Config.Db.TablePrefix, "user_department")
}

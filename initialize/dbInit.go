package initialize

import (
	"fmt"
	"server/app/models"
	"server/config"
	"server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DbInit(c *config.Config) *gorm.DB {
	dsn := fmt.Sprint(c.Db.User, ":", c.Db.PassWord, "@tcp(", c.Db.Host, ":", c.Db.Port, ")/", c.Db.Database, "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Db.TablePrefix,
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err.Error())
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(c.Db.MaxOpenConns)
	sqlDb.SetMaxIdleConns(c.Db.MaxIdleConns)
	sqlDb.SetConnMaxIdleTime(time.Hour)
	return db
}

// AutoMigrate 自动迁移数据库表
func AutoMigrate(db *gorm.DB) {
	// 自动迁移表结构
	err := db.AutoMigrate(
		&models.AdminUser{},
		&models.Role{},
		&models.AdminMenu{},
		&models.MenuApiList{},
		&models.Department{},
		&models.OperationLog{},
		&models.File{},
		&models.FileGroup{},
		&models.TimerTask{},    // 添加定时任务表
		&models.TimerTaskLog{}, // 添加定时任务日志表
	)
	if err != nil {
		global.Logger.Errorf("数据库迁移失败: %v", err)
		panic(err)
	}
	global.Logger.Info("数据库迁移成功")
}

func DbClose(db *gorm.DB) func() {
	return func() {
		db, _ := db.DB()
		_ = db.Close()
	}
}

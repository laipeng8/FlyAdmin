package repositorys

import (
	"gorm.io/gorm"
	"server/global"
)

type BaseRepository struct {
	db *gorm.DB
}

// SetDb
//
//	@Description:	设置Db
//	@receiver		r *BaseRepository
//	@param			db	*gorm.DB
//	@return			*BaseRepository
func (r *BaseRepository) SetDb(db *gorm.DB) {
	r.db = db
}

// getDb
//
//	@Description:	私有方法 主要应对事务问题
//	@receiver		r *BaseRepository
//	@return			*gorm.DB
func (r *BaseRepository) getDb() *gorm.DB {
	var db *gorm.DB
	if r.db == nil {
		db = global.Db
	} else {
		db = r.db
	}
	return db.Debug() // 启用调试模式
}

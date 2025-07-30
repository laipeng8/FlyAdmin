package global

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type GAD_MODEL struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"` // 主键ID
	CreatedAt *LocalTime     `json:"created_at" gorm:"type:datetime(3)"` // 创建时间
	UpdatedAt *LocalTime     `json:"updated_at" gorm:"type:datetime(3)"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                     // 删除时间
}

type LocalTime time.Time

// MarshalJSON 实现 JSON 序列化
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// UnmarshalJSON 实现 JSON 反序列化
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	// 去除 JSON 字符串的引号
	str := string(data)
	if str == "null" {
		return nil
	}
	str = str[1 : len(str)-1]

	// 解析时间
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return err
	}
	*t = LocalTime(parsedTime)
	return nil
}

// Scan 实现数据库扫描接口
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Value 实现数据库值接口
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

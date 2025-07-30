package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONRawMessageMap 定义自定义类型
type JSONRawMessageMap map[string]json.RawMessage

// Value 实现 driver.Valuer 接口，用于序列化
func (j JSONRawMessageMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 实现 sql.Scanner 接口，用于反序列化
func (j *JSONRawMessageMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for JSONRawMessageMap")
	}
	return json.Unmarshal(bytes, j)
}

package reflect

import "reflect"

// updateFields 通用的反射工具方法，用于将一个结构体（src）中的非空字段值更新到另一个结构体（dst）中。它的主要作用是避免手动逐个字段赋值，同时支持跳过空值字段和不可更新的字段。
func UpdateFields(dst, src interface{}) error {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()
	dstType := dstValue.Type()

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		if srcField.Kind() == reflect.Ptr && !srcField.IsNil() {
			fieldName := dstType.Field(i).Name
			dstField := dstValue.FieldByName(fieldName)

			// 检查字段是否可更新
			if dstField.IsValid() && dstField.CanSet() {
				// 如果字段标记为不可更新（gorm:"-"），则跳过
				if gormTag := dstType.Field(i).Tag.Get("gorm"); gormTag == "-" {
					continue
				}

				// 检查字段值是否为空
				switch srcField.Elem().Kind() {
				case reflect.String:
					if srcField.Elem().String() == "" { // 空字符串不更新
						continue
					}
				}

				// 更新字段
				dstField.Set(srcField)
			}
		}
	}

	return nil
}

package myutils

import (
	"reflect"
	"strings"
)

type IStruct interface {
	GetStructData() interface{}
}

// StructToMap 结构体转map
// fields为 不需要 转换的结构体的字段名，多个字段直接连着写
func StructToMap[T any](st T, fields string) map[string]interface{} {
	m := make(map[string]interface{})

	val := reflect.ValueOf(st)

	// 等效 valType := val.Type()
	valType := reflect.TypeOf(st)

	// 判断如果val是指针，指针指向的值的反射对象
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// 判断如果val为空，直接返回空map
	if val.Kind() != reflect.Struct {
		return m
	}

	// NumField 返回该类型的结构体字段数量
	for i := 0; i < valType.NumField(); i++ {
		name := valType.Field(i).Name
		tag := valType.Field(i).Tag.Get("json")
		if !strings.Contains(fields, name) {
			//// 查找多次括号
			//tagParts := strings.Split(tag, ",")
			//if len(tagParts) > 0 {
			//	name = tagParts[0]
			//}

			// 只需要查找一次逗号
			if tag != "" {
				index := strings.Index(tag, ",")
				if index == -1 {
					name = tag
				} else {
					name = tag[:index]
				}
			}

			m[name] = val.Field(i).Interface()
		}
	}
	return m
}

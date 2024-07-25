package utils

import (
	"fmt"
	"reflect"
	"time"
)

// HasValue 检查 reflect.Value 是否有值
func HasValue(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface, reflect.Chan:
		return !val.IsNil()
	case reflect.Bool:
		return val.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return val.Float() != 0
	case reflect.String:
		return val.String() != ""
	case reflect.Struct:
		// 特别处理 time.Time 类型
		if _, ok := val.Interface().(time.Time); ok {
			return !val.Interface().(time.Time).IsZero()
		}
		// 处理其他 struct 类型，你可以根据需要扩展这部分
		return true // 或者其他逻辑
	default:
		return false
	}
}

// StructToMap 遍历结构体字段，如果字段有值，则将其保存到 map 中
func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	// 确保我们处理的是结构体
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		fmt.Println("Expecting a struct")
		return nil
	}
	// 遍历结构体的所有字段
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		fieldName := typeField.Name

		// 检查字段是否有值
		if HasValue(valueField) {
			result[fieldName] = valueField.Interface()
		}
	}
	return result
}

func StructsJsonToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	// 确保我们处理的是结构体
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		fmt.Println("Expecting a struct")
		return nil
	}
	// 遍历结构体的所有字段
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		fieldName := typeField.Tag.Get("json")

		// 检查字段是否有值
		if HasValue(valueField) {
			result[fieldName] = valueField.Interface()
		}
	}
	return result
}

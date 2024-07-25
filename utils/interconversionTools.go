package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

func StrToInt64(req string) int64 {
	num, _ := strconv.ParseInt(req, 10, 64)
	return num
}

func StrToFloat64(req string) float64 {
	floatValue, _ := strconv.ParseFloat(req, 64)
	return floatValue
}

func GroupByField(slice interface{}, fieldName string) (map[interface{}][]interface{}, error) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("provided value is not a slice")
	}
	grouped := make(map[interface{}][]interface{})
	for i := 0; i < sliceVal.Len(); i++ {
		element := sliceVal.Index(i).Interface()
		fieldVal := reflect.ValueOf(element).FieldByName(fieldName)

		if !fieldVal.IsValid() {
			return nil, fmt.Errorf("field %s does not exist in the slice element type", fieldName)
		}
		key := fieldVal.Interface()
		grouped[key] = append(grouped[key], element)
	}

	return grouped, nil
}

func TransferData(src, dst interface{}) {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		name := srcVal.Type().Field(i).Name
		dstField := dstVal.FieldByName(name)

		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(srcVal.Field(i))
		}
	}
}

func ConvertInterfaceSliceToStringSlice(interfaces []interface{}) ([]string, error) {
	var strings []string
	for _, v := range interfaces {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("类型断言失败: 预期 string，实际为 %v", reflect.TypeOf(v))
		}
		strings = append(strings, str)
	}
	return strings, nil
}

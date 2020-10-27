package transfer

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func MapToStruct(s map[string]interface{}, obj interface{}) interface{} {
	err := mapstructure.Decode(s, &obj)
	if err != nil {
		fmt.Println(err)
	}
	return obj
}

func StringToInt(slice []string) (rs []int) {
	if len(slice) <= 0 {
		return
	}
	for _, num := range slice {
		inum, err := strconv.Atoi(num)
		if err != nil {
			rs = append(rs, 0)
			continue
		} else {
			rs = append(rs, inum)
		}
	}
	return
}

func IntToString(slice []int) (rs []string) {
	if len(slice) <= 0 {
		return
	}
	for _, num := range slice {
		rs = append(rs, strconv.Itoa(num))
	}
	return
}

func StringToInt64(slice []string) (rs []int64) {
	if len(slice) <= 0 {
		return
	}
	for _, num := range slice {
		inum, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			rs = append(rs, 0)
			continue
		} else {
			rs = append(rs, inum)
		}
	}
	return
}

func Int64ToString(slice []int64) (rs []string) {
	if len(slice) <= 0 {
		return
	}
	for _, num := range slice {
		rs = append(rs, strconv.FormatInt(num, 10))
	}
	return
}

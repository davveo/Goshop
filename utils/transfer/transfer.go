package transfer

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
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

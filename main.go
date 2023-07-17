package main

import (
	"fmt"
	"reflect"
)

type SimpleStruct struct {
	Name string `keyname:"name"`
	ID   int    `keyname:"id"`
}

func main() {
	simple := SimpleStruct{
		Name: "Superman",
		ID:   1522,
	}
	StructToMap(simple)
}

func StructToMap(item interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	structReflectValue := reflect.ValueOf(item)
	structReflectType := reflect.TypeOf(item)
	numberF := structReflectValue.NumField()
	for i := 0; i < numberF; i++ {
		v := structReflectValue.Field(i)
		t := structReflectType.Field(i).Name
		if v.Kind() == reflect.String {
			newMap[t] = v.Interface().(string)
		} else {
			newMap[t] = v.Interface().(int)
		}

		fmt.Println(t, v)
	}
	fmt.Println(newMap)
	return nil
}

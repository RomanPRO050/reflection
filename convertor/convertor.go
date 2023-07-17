package convertor

import "reflect"

func MapToStruct(mp map[string]interface{}, item interface{}) error {
	mapR := reflect.TypeOf(mp)

	return nil
}

func StructToMap(item interface{}) map[string]interface{} {
	structReflectType := reflect.TypeOf(item)
	field := structReflectType.Field(0)
	structReflectValue := reflect.ValueOf(item)
	return nil
}

package util

import (
	"errors"
	"reflect"
)

func StructToSlice(x interface{}) []interface{} {
	v := reflect.ValueOf(x)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return values
}

func ValidateBlankFields(x interface{}) error {
	y := StructToSlice(x)
	for _, v := range y {
		if v == nil || v == "" {
			return errors.New("all fields must be filled")
		}
	}
	return nil
}

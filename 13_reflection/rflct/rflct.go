package rflct

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	value := getValue(x)
	numValues := 0

	var getField func(int) reflect.Value

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walk(value.Field(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(value.MapIndex(key).Interface(), fn)
		}
	}
	for i := 0; i < numValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}

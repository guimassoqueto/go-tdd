package main

import "reflect"


func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numOfValues = val.Len()
		getField = val.Index
	}
	
	for i := 0; i < numOfValues; i++ {
		Walk(getField(i).Interface(), fn)
	}
}


func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}

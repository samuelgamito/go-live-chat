package misc

import "reflect"

func SliceContainsString(slice interface{}, target string, searchString string) int {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic("SliceContainsString expects a slice")
	}

	for i := 0; i < val.Len(); i++ {
		item := val.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		if item.Kind() == reflect.Struct {
			field := item.FieldByName(target)
			if field.IsValid() && field.Interface() == searchString {
				return i
			}
		} else {
			panic("SliceContainsString needs a struct")
		}

	}
	return -1
}

func RemoveByIndex(slice interface{}, index int) interface{} {
	val := reflect.ValueOf(slice)

	if val.Kind() != reflect.Slice {
		panic("Input is not a slice")
	}

	if index < 0 || index >= val.Len() {
		panic("Index out of range")
	}

	result := reflect.AppendSlice(val.Slice(0, index), val.Slice(index+1, val.Len()))
	return result.Interface()
}

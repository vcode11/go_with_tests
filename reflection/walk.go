package reflection

import "reflect"

// Take a struct x of interface{} and call fn for all strings inside
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	var getField func(int) reflect.Value
	numValues := 0
	switch val.Kind() {
	case reflect.Struct:
		getField = val.Field
		numValues = val.NumField()
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		getField = val.Index
		numValues = val.Len()
	case reflect.String:
		fn(val.String())
	}
	for i := 0; i < numValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
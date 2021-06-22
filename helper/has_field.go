package helper

import "reflect"

// HasField checks to see if the field is present in the struct
func HasField(data interface{}, name string) bool {
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

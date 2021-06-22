package helper

import "reflect"

func Last(x int, a interface{}) bool {
	return x == reflect.ValueOf(a).Len()-1
}

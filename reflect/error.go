package reflect

import "reflect"

func ErrorIsNil(err error) bool {
	if err == nil {
		return true
	}
	val := reflect.ValueOf(err)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return true
	}
	return false
}

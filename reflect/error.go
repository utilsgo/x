package reflect

import "reflect"

// ErrorIsNil returns true if the error is nil or a pointer to nil.
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

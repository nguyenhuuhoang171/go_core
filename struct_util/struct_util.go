package struct_util

import "reflect"

// IsEmptyStruct reports whether v is a nil pointer, or a struct (or pointer
// to struct) whose fields are all zero-valued.
func IsEmptyStruct[T any](v T) bool {
	rv := reflect.ValueOf(v)

	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return true
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if !field.CanInterface() {
			continue
		}

		if field.Kind() == reflect.Ptr {
			if !field.IsNil() {
				return false
			}
		} else if !field.IsZero() {
			return false
		}
	}

	return true
}

package gorm_tag

import (
	"reflect"
	"strings"
)

// HasZeroPrimaryKey reports whether entity is a nil pointer, or a struct (or
// pointer to struct) whose gorm primary-key field(s) are zero-valued.
func HasZeroPrimaryKey(entity any) bool {
	v := reflect.ValueOf(entity)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return false
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		if !isPrimaryKeyTag(t.Field(i).Tag.Get("gorm")) {
			continue
		}
		if v.Field(i).IsZero() {
			return true
		}
	}
	return false
}

// isPrimaryKeyTag reports whether a gorm struct tag marks the field as a primary
// key (gorm accepts both "primaryKey" and the legacy "primary_key").
func isPrimaryKeyTag(tag string) bool {
	for _, part := range strings.Split(tag, ";") {
		part = strings.TrimSpace(part)
		if strings.EqualFold(part, "primaryKey") || strings.EqualFold(part, "primary_key") {
			return true
		}
	}
	return false
}

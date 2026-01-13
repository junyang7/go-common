package _is

import (
	"reflect"
)

func Empty(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.String() == ""
	}
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Array, reflect.Struct:
		return v.IsZero()
	case reflect.Ptr:
		return v.IsNil()
	}
	if v.IsNil() {
		return true
	}
	return false
}
func Numeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
func Alpha(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
			return false
		}
	}
	return true
}
func AlphaLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}
func AlphaUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < 'A' || c > 'Z' {
			return false
		}
	}
	return true
}

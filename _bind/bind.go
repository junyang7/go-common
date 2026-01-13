package _bind

import (
	"github.com/junyang7/go-common/_as"
	"reflect"
)

func Do(v interface{}, data map[string]interface{}) error {
	if data == nil {
		return nil
	}
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldName := fieldType.Name
		tagName := fieldType.Tag.Get("json")
		if tagName == "" {
			tagName = fieldName
		}
		dataValue, exists := data[tagName]
		if !exists {
			continue
		}
		if dataValue == nil {
			continue
		}
		if field.Kind() == reflect.Struct {
			if dataMap, ok := dataValue.(map[string]interface{}); ok {
				_ = Do(field.Addr().Interface(), dataMap)
			}
		} else if field.Kind() == reflect.Map {
			if field.Type().Key().Kind() == reflect.String {
				if dataMap, ok := dataValue.(map[string]interface{}); ok {
					if field.IsNil() {
						field.Set(reflect.MakeMap(field.Type()))
					}
					for k, v := range dataMap {
						elemKind := field.Type().Elem().Kind()
						if elemKind == reflect.Interface {
							field.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
						} else if elemKind == reflect.String {
							field.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(_as.String(v)))
						} else if elemKind == reflect.Int {
							field.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(_as.Int(v)))
						} else if elemKind == reflect.Float64 {
							field.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(_as.Float64(v)))
						} else {
							continue
						}
					}
				} else {
					continue
				}
			}
		} else {
			if field.CanSet() {
				switch field.Kind() {
				case reflect.String:
					field.SetString(_as.String(dataValue))
				case reflect.Bool:
					field.SetBool(_as.Bool(dataValue))
				case reflect.Int:
					field.SetInt(int64(_as.Int(dataValue)))
				case reflect.Int64:
					field.SetInt(_as.Int64(dataValue))
				case reflect.Uint:
					field.SetUint(uint64(_as.Uint(dataValue)))
				case reflect.Uint64:
					field.SetUint(_as.Uint64(dataValue))
				case reflect.Float64:
					field.SetFloat(_as.Float64(dataValue))
				case reflect.Slice:
					if field.Type().Elem().Kind() == reflect.Uint8 {
						field.SetBytes(_as.ByteList(dataValue))
					}
				default:
					continue
				}
			}
		}
	}
	return nil
}

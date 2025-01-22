package _as

import (
	"strconv"
	"strings"
)

func String(value interface{}) string {
	switch value.(type) {
	case []byte:
		return string(value.(interface{}).([]byte))
	case string:
		return value.(interface{}).(string)
	case int8:
		return strconv.FormatInt(int64(value.(interface{}).(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(interface{}).(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(interface{}).(int32)), 10)
	case int64:
		return strconv.FormatInt(value.(interface{}).(int64), 10)
	case int:
		return strconv.FormatInt(int64(value.(interface{}).(int)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(interface{}).(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(interface{}).(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(interface{}).(uint32)), 10)
	case uint64:
		return strconv.FormatUint(value.(interface{}).(uint64), 10)
	case uint:
		return strconv.FormatUint(uint64(value.(interface{}).(uint)), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(interface{}).(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value.(interface{}).(float64), 'f', -1, 64)
	case bool:
		if value.(interface{}).(bool) {
			return "1"
		} else {
			return ""
		}
	default:
		return ""
	}
}
func Bool(value interface{}) bool {
	switch value.(type) {
	case []byte:
		return len(value.(interface{}).([]byte)) > 0
	case string:
		return "" != strings.TrimSpace(value.(interface{}).(string))
	case int8:
		return value.(interface{}).(int8) > 0
	case int16:
		return value.(interface{}).(int16) > 0
	case int32:
		return value.(interface{}).(int32) > 0
	case int64:
		return value.(interface{}).(int64) > 0
	case int:
		return value.(interface{}).(int) > 0
	case uint8:
		return value.(interface{}).(uint8) > 0
	case uint16:
		return value.(interface{}).(uint16) > 0
	case uint32:
		return value.(interface{}).(uint32) > 0
	case uint64:
		return value.(interface{}).(uint64) > 0
	case uint:
		return value.(interface{}).(uint) > 0
	case float32:
		return value.(interface{}).(float32) > 0
	case float64:
		return value.(interface{}).(float64) > 0
	case bool:
		return value.(interface{}).(bool)
	default:
		return false
	}
}
func ByteList(value interface{}) []byte {
	return []byte(String(value))
}
func Float64(value interface{}) float64 {
	switch value.(type) {
	case []byte:
		return 0
	case string:
		f64, err := strconv.ParseFloat(value.(interface{}).(string), 64)
		if nil != err {
			return 0
		}
		return f64
	case int8:
		return float64(value.(interface{}).(int8))
	case int16:
		return float64(value.(interface{}).(int16))
	case int32:
		return float64(value.(interface{}).(int32))
	case int64:
		return float64(value.(interface{}).(int64))
	case int:
		return float64(value.(interface{}).(int))
	case uint8:
		return float64(value.(interface{}).(uint8))
	case uint16:
		return float64(value.(interface{}).(uint16))
	case uint32:
		return float64(value.(interface{}).(uint32))
	case uint64:
		return float64(value.(interface{}).(uint64))
	case uint:
		return float64(value.(interface{}).(uint))
	case float32:
		return float64(value.(interface{}).(float32))
	case float64:
		return value.(interface{}).(float64)
	case bool:
		if value.(interface{}).(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}
func Int64(value interface{}) int64 {
	switch value.(type) {
	case []byte:
		i64, err := strconv.ParseInt(strings.TrimSpace(string(value.(interface{}).([]byte))), 10, 64)
		if nil != err {
			return 0
		}
		return i64
	case string:
		i64, err := strconv.ParseInt(strings.TrimSpace(value.(interface{}).(string)), 10, 64)
		if nil != err {
			return 0
		}
		return i64
	case int8:
		return int64(value.(interface{}).(int8))
	case int16:
		return int64(value.(interface{}).(int16))
	case int32:
		return int64(value.(interface{}).(int32))
	case int64:
		return value.(interface{}).(int64)
	case int:
		return int64(value.(interface{}).(int))
	case uint8:
		return int64(value.(interface{}).(uint8))
	case uint16:
		return int64(value.(interface{}).(uint16))
	case uint32:
		return int64(value.(interface{}).(uint32))
	case uint64:
		return int64(value.(interface{}).(uint64))
	case uint:
		return int64(value.(interface{}).(uint))
	case float32:
		return int64(value.(interface{}).(float32))
	case float64:
		return int64(value.(interface{}).(float64))
	case bool:
		if value.(interface{}).(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}
func Int(value interface{}) int {
	return int(Int64(value))
}
func Uint64(value interface{}) uint64 {
	switch value.(type) {
	case []byte:
		ui64, err := strconv.ParseUint(strings.TrimSpace(string(value.(interface{}).([]byte))), 10, 64)
		if nil != err {
			return 0
		}
		return ui64
	case string:
		ui64, err := strconv.ParseUint(strings.TrimSpace(value.(interface{}).(string)), 10, 64)
		if nil != err {
			return 0
		}
		return ui64
	case int8:
		return uint64(value.(interface{}).(int8))
	case int16:
		return uint64(value.(interface{}).(int16))
	case int32:
		return uint64(value.(interface{}).(int32))
	case int64:
		return uint64(value.(interface{}).(int64))
	case int:
		return uint64(value.(interface{}).(int))
	case uint8:
		return uint64(value.(interface{}).(uint8))
	case uint16:
		return uint64(value.(interface{}).(uint16))
	case uint32:
		return uint64(value.(interface{}).(uint32))
	case uint64:
		return value.(interface{}).(uint64)
	case uint:
		return uint64(value.(interface{}).(uint))
	case float32:
		return uint64(value.(interface{}).(float32))
	case float64:
		return uint64(value.(interface{}).(float64))
	case bool:
		if value.(interface{}).(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}
func Uint(value interface{}) uint {
	return uint(Uint64(value))
}

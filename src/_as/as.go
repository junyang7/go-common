package _as

import (
	"strconv"
	"strings"
)

func String(data interface{}) string {
	switch data.(type) {
	case []byte:
		return string(data.(interface{}).([]byte))
	case string:
		return data.(interface{}).(string)
	case int8:
		return strconv.FormatInt(int64(data.(interface{}).(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(data.(interface{}).(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(data.(interface{}).(int32)), 10)
	case int64:
		return strconv.FormatInt(data.(interface{}).(int64), 10)
	case int:
		return strconv.FormatInt(int64(data.(interface{}).(int)), 10)
	case uint8:
		return strconv.FormatUint(uint64(data.(interface{}).(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(data.(interface{}).(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(data.(interface{}).(uint32)), 10)
	case uint64:
		return strconv.FormatUint(data.(interface{}).(uint64), 10)
	case uint:
		return strconv.FormatUint(uint64(data.(interface{}).(uint)), 10)
	case float32:
		return strconv.FormatFloat(float64(data.(interface{}).(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(data.(interface{}).(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(data.(interface{}).(bool))
	default:
		return ""
	}
}
func Bool(data interface{}) bool {
	switch data.(type) {
	case []byte:
		return len(data.(interface{}).([]byte)) > 0
	case string:
		return "" != strings.TrimSpace(data.(interface{}).(string))
	case int8:
		return data.(interface{}).(int8) > 0
	case int16:
		return data.(interface{}).(int16) > 0
	case int32:
		return data.(interface{}).(int32) > 0
	case int64:
		return data.(interface{}).(int64) > 0
	case int:
		return data.(interface{}).(int) > 0
	case uint8:
		return data.(interface{}).(uint8) > 0
	case uint16:
		return data.(interface{}).(uint16) > 0
	case uint32:
		return data.(interface{}).(uint32) > 0
	case uint64:
		return data.(interface{}).(uint64) > 0
	case uint:
		return data.(interface{}).(uint) > 0
	case float32:
		return data.(interface{}).(float32) > 0
	case float64:
		return data.(interface{}).(float64) > 0
	case bool:
		return data.(interface{}).(bool)
	default:
		return false
	}
}
func Int64(data interface{}) int64 {
	switch data.(type) {
	case []byte:
		i64, err := strconv.ParseInt(strings.TrimSpace(string(data.(interface{}).([]byte))), 10, 64)
		if nil != err {
			return 0
		}
		return i64
	case string:
		i64, err := strconv.ParseInt(strings.TrimSpace(data.(interface{}).(string)), 10, 64)
		if nil != err {
			return 0
		}
		return i64
	case int8:
		return int64(data.(interface{}).(int8))
	case int16:
		return int64(data.(interface{}).(int16))
	case int32:
		return int64(data.(interface{}).(int32))
	case int64:
		return data.(interface{}).(int64)
	case int:
		return int64(data.(interface{}).(int))
	case uint8:
		return int64(data.(interface{}).(uint8))
	case uint16:
		return int64(data.(interface{}).(uint16))
	case uint32:
		return int64(data.(interface{}).(uint32))
	case uint64:
		return int64(data.(interface{}).(uint64))
	case uint:
		return int64(data.(interface{}).(uint))
	case float32:
		return int64(data.(interface{}).(float32))
	case float64:
		return int64(data.(interface{}).(float64))
	case bool:
		if data.(interface{}).(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}
func Int(data interface{}) int {
	return int(Int64(data))
}
func ByteList(data interface{}) []byte {
	return []byte(String(data))
}
func Float64(data interface{}) float64 {
	switch data.(type) {
	case []byte:
		return 0
	case string:
		f64, err := strconv.ParseFloat(data.(interface{}).(string), 64)
		if nil != err {
			return 0
		}
		return f64
	case int8:
		return float64(data.(interface{}).(int8))
	case int16:
		return float64(data.(interface{}).(int16))
	case int32:
		return float64(data.(interface{}).(int32))
	case int64:
		return float64(data.(interface{}).(int64))
	case int:
		return float64(data.(interface{}).(int))
	case uint8:
		return float64(data.(interface{}).(uint8))
	case uint16:
		return float64(data.(interface{}).(uint16))
	case uint32:
		return float64(data.(interface{}).(uint32))
	case uint64:
		return float64(data.(interface{}).(uint64))
	case uint:
		return float64(data.(interface{}).(uint))
	case float32:
		return float64(data.(interface{}).(float32))
	case float64:
		return data.(interface{}).(float64)
	case bool:
		if data.(interface{}).(bool) {
			return 1
		}
		return 0
	default:
		return 0
	}
}

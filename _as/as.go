package _as

import (
	"fmt"
	"strconv"
)

///
/// 基础方法
///

func String(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case nil:
		return ""
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		if v {
			return "true"
		} else {
			return "false"
		}
	default:
		return fmt.Sprintf("%v", v)
	}
}
func Bool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		return v != ""
	case []byte:
		return len(v) > 0
	case nil:
		return false
	case int:
		return v != 0
	case int8:
		return v != 0
	case int16:
		return v != 0
	case int32:
		return v != 0
	case int64:
		return v != 0
	case uint:
		return v != 0
	case uint8:
		return v != 0
	case uint16:
		return v != 0
	case uint32:
		return v != 0
	case uint64:
		return v != 0
	case float32:
		return v != 0
	case float64:
		return v != 0
	default:
		return false
	}
}
func Float64(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case string:
		f64, err := strconv.ParseFloat(v, 64)
		if nil != err {
			return 0
		}
		return f64
	case []byte:
		f64, err := strconv.ParseFloat(string(v), 64)
		if nil != err {
			return 0
		}
		return f64
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}
func ByteList(value interface{}) []byte {
	switch v := value.(type) {
	case []byte:
		return v
	default:
		return []byte(String(value))
	}
}

///
///  衍生方法
///

func Int64(value interface{}) int64 {
	return int64(Float64(value))
}
func Int(value interface{}) int {
	return int(Float64(value))
}
func Uint64(value interface{}) uint64 {
	m := Int64(value)
	if m >= 0 {
		return uint64(m)
	} else {
		return uint64(m * -1)
	}
}
func Uint(value interface{}) uint {
	m := Int(value)
	if m >= 0 {
		return uint(m)
	} else {
		return uint(m * -1)
	}
}

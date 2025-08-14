package _as

import (
	"fmt"
	"strconv"
)

func String(value interface{}) string {
	switch v := value.(type) {
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
	case string:
		return v
	case []byte:
		return string(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
func Bool(value interface{}) bool {
	switch v := value.(type) {
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
	case string:
		return v == "true"
	case []byte:
		return string(v) == "true"
	case bool:
		return v
	default:
		return false
	}
}
func ByteList(value interface{}) []byte {
	switch v := value.(type) {
	case nil:
		return []byte{}
	case string:
		return []byte(v)
	case []byte:
		return v
	default:
		return []byte(String(value))
	}
}
func Float64(value interface{}) float64 {
	switch v := value.(type) {
	case nil:
		return 0
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
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}
func Int64(value interface{}) int64 {
	switch v := value.(type) {
	case nil:
		return 0
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i64, err := strconv.ParseInt(v, 10, 64)
		if nil != err {
			return 0
		}
		return i64
	case []byte:
		i64, err := strconv.ParseInt(string(v), 10, 64)
		if nil != err {
			return 0
		}
		return i64
	default:
		return 0
	}
}
func Int(value interface{}) int {
	return int(Int64(value))
}
func Uint64(value interface{}) uint64 {
	switch v := value.(type) {
	case nil:
		return 0
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		ui64, err := strconv.ParseUint(v, 10, 64)
		if nil != err {
			return 0
		}
		return ui64
	case []byte:
		ui64, err := strconv.ParseUint(string(v), 10, 64)
		if nil != err {
			return 0
		}
		return ui64
	default:
		return 0
	}
}
func Uint(value interface{}) uint {
	return uint(Uint64(value))
}

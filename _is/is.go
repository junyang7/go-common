package _is

import (
	"fmt"
	"strings"
)

func Empty(value interface{}) bool {
	if `<nil>` == fmt.Sprintf("%v", value) {
		return true
	}
	switch value.(type) {
	case nil:
		return true
	case string:
		return "" == strings.TrimSpace(value.(string))
	case int8:
		return 0 == value.(int8)
	case int16:
		return 0 == value.(int16)
	case int32:
		return 0 == value.(int32)
	case int64:
		return 0 == value.(int64)
	case int:
		return 0 == value.(int)
	case uint8:
		return 0 == value.(uint8)
	case uint16:
		return 0 == value.(uint16)
	case uint32:
		return 0 == value.(uint32)
	case uint64:
		return 0 == value.(uint64)
	case float32:
		return 0 == value.(float32)
	case float64:
		return 0 == value.(float64)
	case bool:
		return false == value.(bool)
	case []byte:
		return 0 == len(value.([]byte))
	case map[string]string:
		return 0 == len(value.(map[string]string))
	case []map[string]string:
		return 0 == len(value.([]map[string]string))
	case [][]map[string]string:
		return 0 == len(value.([][]map[string]string))
	case []map[string]interface{}:
		return 0 == len(value.([]map[string]interface{}))
	case [][]map[string]interface{}:
		return 0 == len(value.([][]map[string]interface{}))
	default:
		return false
	}
}

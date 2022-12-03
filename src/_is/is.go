package _is

import "strings"

func Empty(value interface{}) bool {
	switch value.(type) {
	case string:
		return "" == strings.TrimSpace(value.(string))
	case int8:
		return value.(int8) > 0
	case int16:
		return value.(int16) > 0
	case int32:
		return value.(int32) > 0
	case int64:
		return value.(int64) > 0
	case int:
		return value.(int) > 0
	case uint8:
		return value.(uint8) > 0
	case uint16:
		return value.(uint16) > 0
	case uint32:
		return value.(uint32) > 0
	case uint64:
		return value.(uint64) > 0
	case float32:
		return value.(float32) > 0
	case float64:
		return value.(float64) > 0
	case bool:
		return value.(bool)
	case []byte:
		return len(value.([]byte)) > 0
	case map[string]string:
		return len(value.(map[string]string)) > 0
	case []map[string]string:
		return len(value.([]map[string]string)) > 0
	case [][]map[string]string:
		return len(value.([][]map[string]string)) > 0
	case []map[string]interface{}:
		return len(value.([]map[string]interface{})) > 0
	case [][]map[string]interface{}:
		return len(value.([][]map[string]interface{})) > 0
	default:
		return false
	}

}
func NotEmpty(value interface{}) bool {
	return !Empty(value)
}

package _is

import (
	"fmt"
	"reflect"
)

// Empty 判断一个值是否是空
// bool,int8,int16,int32,int64,int,uint8,uint16,uint32,uint64,uint,float32,float64,string,nil,[]x,map[x]x,[x]x,chan x,*x
func Empty(value interface{}) bool {
	var t string = fmt.Sprintf("%T", value)
	switch t {
	case `bool`:
		return false == value.(bool)
	case `int8`:
		return 0 == value.(int8)
	case `int16`:
		return 0 == value.(int16)
	case `int32`:
		return 0 == value.(int32)
	case `int64`:
		return 0 == value.(int64)
	case `int`:
		return 0 == value.(int)
	case `uint8`:
		return 0 == value.(uint8)
	case `uint16`:
		return 0 == value.(uint16)
	case `uint32`:
		return 0 == value.(uint32)
	case `uint64`:
		return 0 == value.(uint64)
	case `uint`:
		return 0 == value.(uint)
	case `float32`:
		return 0 == value.(float32)
	case `float64`:
		return 0 == value.(float64)
	case `string`:
		return "" == value.(string)
	case `<nil>`:
		return true
	}
	if `<nil>` == fmt.Sprintf("%v", value) {
		return true
	}
	v := reflect.ValueOf(value)
	k := v.Kind().String()
	switch k {
	case `slice`, `map`, `chan`:
		return v.Len() == 0
	case `array`:
		return v.IsZero()
	case `ptr`:
		return v.IsNil()
	}
	return true
}

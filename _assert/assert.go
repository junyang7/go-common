package _assert

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

const Default = 1
const Float = 2
const Time = 3
const List = 4

// 内部辅助函数：获取调用者信息
func getCallerInfo(skip int) (file string, line int, funcName string) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	return
}

// 内部辅助函数：打印成功信息
func printSuccess(t *testing.T, file string, funcName string, line int, tE string, tG string, e any, g any) {
	fmt.Println("\t", fmt.Sprintf("\033[0;32m success \033[0m"), file, funcName, line, tE, tG, e, g)
}

// 内部辅助函数：打印失败信息
func printFailure(t *testing.T, file string, funcName string, line int, tE string, tG string, e any, g any) {
	fmt.Println("\t", fmt.Sprintf("\033[0;31m failure \033[0m"), file, funcName, line, tE, tG, e, g)
	t.FailNow()
}

func equal(t *testing.T, e any, g any, m int) {
	file, line, funcName := getCallerInfo(2)
	tE := fmt.Sprintf("%T", e)
	tG := fmt.Sprintf("%T", g)
	if tE != tG {
		fmt.Println("\t", fmt.Sprintf("\033[0;31m failure \033[0m"), file, funcName, line, tE, tG)
		t.FailNow()
		return
	}
	switch m {
	case Float:
		l := 0
		partList := strings.Split(fmt.Sprintf("%v", e), ".")
		if len(partList) > 1 {
			l = len(partList[1])
		}
		f := "%." + fmt.Sprintf("%d", l) + "f"
		if fmt.Sprintf(f, e) != fmt.Sprintf(f, g) {
			printFailure(t, file, funcName, line, tE, tG, e, g)
			return
		}
		break
	case Time:
		if !e.(time.Time).Equal(g.(time.Time)) {
			printFailure(t, file, funcName, line, tE, tG, e, g)
			return
		}
		break
	case List:
		if !reflect.DeepEqual(e, g) {
			printFailure(t, file, funcName, line, tE, tG, e, g)
			return
		}
		break
	default:
		if fmt.Sprintf("%v", e) != fmt.Sprintf("%v", g) {
			printFailure(t, file, funcName, line, tE, tG, e, g)
			return
		}
		break
	}
	printSuccess(t, file, funcName, line, tE, tG, e, g)
}

// Equal 验证两个值相等
func Equal(t *testing.T, expect any, get any) {
	equal(t, expect, get, Default)
}

// EqualByFloat 验证两个浮点数相等（根据精度）
func EqualByFloat(t *testing.T, expect any, get any) {
	equal(t, expect, get, Float)
}

// EqualByTime 验证两个时间相等
func EqualByTime(t *testing.T, expect any, get any) {
	equal(t, expect, get, Time)
}

// EqualByList 验证两个列表/切片深度相等
func EqualByList(t *testing.T, expect any, get any) {
	equal(t, expect, get, List)
}

// NotEqual 验证两个值不相等
func NotEqual(t *testing.T, notExpect any, get any) {
	file, line, funcName := getCallerInfo(1)
	tE := fmt.Sprintf("%T", notExpect)
	tG := fmt.Sprintf("%T", get)

	// 类型不同，直接判定为不相等（成功）
	if tE != tG {
		printSuccess(t, file, funcName, line, tE, tG, notExpect, get)
		return
	}

	// 类型相同，检查值
	if fmt.Sprintf("%v", notExpect) == fmt.Sprintf("%v", get) {
		// 相等了，但期望不相等，失败
		printFailure(t, file, funcName, line, tE, tG, notExpect, get)
		return
	}

	printSuccess(t, file, funcName, line, tE, tG, notExpect, get)
}

// True 验证值为 true
func True(t *testing.T, value bool) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)

	if !value {
		printFailure(t, file, funcName, line, tV, "bool", true, value)
		return
	}

	printSuccess(t, file, funcName, line, tV, "bool", true, value)
}

// False 验证值为 false
func False(t *testing.T, value bool) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)

	if value {
		printFailure(t, file, funcName, line, tV, "bool", false, value)
		return
	}

	printSuccess(t, file, funcName, line, tV, "bool", false, value)
}

// Nil 验证值为 nil
func Nil(t *testing.T, value any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)

	if value != nil && !reflect.ValueOf(value).IsNil() {
		printFailure(t, file, funcName, line, tV, "nil", nil, value)
		return
	}

	printSuccess(t, file, funcName, line, tV, "nil", nil, value)
}

// NotNil 验证值不为 nil
func NotNil(t *testing.T, value any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)

	if value == nil {
		printFailure(t, file, funcName, line, tV, "not nil", "not nil", value)
		return
	}

	// 检查是否是 nil 接口值
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface ||
		rv.Kind() == reflect.Slice || rv.Kind() == reflect.Map ||
		rv.Kind() == reflect.Chan || rv.Kind() == reflect.Func {
		if rv.IsNil() {
			printFailure(t, file, funcName, line, tV, "not nil", "not nil", value)
			return
		}
	}

	printSuccess(t, file, funcName, line, tV, "not nil", "not nil", value)
}

// Error 验证 error 不为 nil
func Error(t *testing.T, err error) {
	file, line, funcName := getCallerInfo(1)

	if err == nil {
		printFailure(t, file, funcName, line, "error", "error", "error", nil)
		return
	}

	printSuccess(t, file, funcName, line, "error", "error", "error", err)
}

// NoError 验证 error 为 nil
func NoError(t *testing.T, err error) {
	file, line, funcName := getCallerInfo(1)

	if err != nil {
		printFailure(t, file, funcName, line, "error", "error", nil, err)
		return
	}

	printSuccess(t, file, funcName, line, "error", "error", nil, nil)
}

// Contains 验证字符串包含子串
func Contains(t *testing.T, str string, substr string) {
	file, line, funcName := getCallerInfo(1)

	if !strings.Contains(str, substr) {
		printFailure(t, file, funcName, line, "string", "string", "contains: "+substr, str)
		return
	}

	printSuccess(t, file, funcName, line, "string", "string", "contains: "+substr, str)
}

// NotContains 验证字符串不包含子串
func NotContains(t *testing.T, str string, substr string) {
	file, line, funcName := getCallerInfo(1)

	if strings.Contains(str, substr) {
		printFailure(t, file, funcName, line, "string", "string", "not contains: "+substr, str)
		return
	}

	printSuccess(t, file, funcName, line, "string", "string", "not contains: "+substr, str)
}

// Empty 验证字符串、切片、数组、map 为空
func Empty(t *testing.T, value any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)

	isEmpty := false

	if value == nil {
		isEmpty = true
	} else {
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
			isEmpty = rv.Len() == 0
		default:
			// 其他类型用字符串判断
			isEmpty = fmt.Sprintf("%v", value) == ""
		}
	}

	if !isEmpty {
		printFailure(t, file, funcName, line, tV, "empty", "empty", value)
		return
	}

	printSuccess(t, file, funcName, line, tV, "empty", "empty", value)
}

// NotEmpty 验证字符串、切片、数组、map 不为空
func NotEmpty(t *testing.T, value any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)

	isEmpty := false

	if value == nil {
		isEmpty = true
	} else {
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
			isEmpty = rv.Len() == 0
		default:
			isEmpty = fmt.Sprintf("%v", value) == ""
		}
	}

	if isEmpty {
		printFailure(t, file, funcName, line, tV, "not empty", "not empty", value)
		return
	}

	printSuccess(t, file, funcName, line, tV, "not empty", "not empty", value)
}

// Len 验证长度
func Len(t *testing.T, value any, length int) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)

	actualLen := 0
	if value != nil {
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
			actualLen = rv.Len()
		default:
			actualLen = len(fmt.Sprintf("%v", value))
		}
	}

	if actualLen != length {
		printFailure(t, file, funcName, line, tV, "length", length, actualLen)
		return
	}

	printSuccess(t, file, funcName, line, tV, "length", length, actualLen)
}

// Panics 验证函数会 panic
func Panics(t *testing.T, fn func()) {
	file, line, funcName := getCallerInfo(1)

	defer func() {
		if r := recover(); r == nil {
			// 没有 panic，失败
			printFailure(t, file, funcName, line, "panic", "panic", "should panic", "no panic")
		} else {
			// 有 panic，成功
			printSuccess(t, file, funcName, line, "panic", "panic", "panics", r)
		}
	}()

	fn()
}

// NotPanics 验证函数不会 panic
func NotPanics(t *testing.T, fn func()) {
	file, line, funcName := getCallerInfo(1)

	defer func() {
		if r := recover(); r != nil {
			// 发生了 panic，失败
			printFailure(t, file, funcName, line, "no panic", "panic", "should not panic", r)
		}
	}()

	fn()

	// 没有 panic，成功
	printSuccess(t, file, funcName, line, "no panic", "no panic", "no panic", "no panic")
}

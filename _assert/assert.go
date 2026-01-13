package _assert

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"
)

const Default = 1
const Float = 2
const Time = 3
const List = 4

func getCallerInfo(skip int) (file string, line int, funcName string) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
	return
}
func printSuccess(t *testing.T, file string, funcName string, line int, tE string, tG string, e any, g any) {
	fmt.Println("\t", fmt.Sprintf("\033[0;32m success \033[0m"), file, funcName, line, tE, tG, e, g)
}
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
func Regexp(t *testing.T, str string, pattern string) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", str)
	tP := fmt.Sprintf("%T", pattern)
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		printFailure(t, file, funcName, line, tV, tP, "regexp error", err)
		return
	}
	if !matched {
		printFailure(t, file, funcName, line, tV, tP, "does not match", str)
		return
	}
	printSuccess(t, file, funcName, line, tV, tP, "matches", str)
}
func Equal(t *testing.T, expect any, get any) {
	equal(t, expect, get, Default)
}
func EqualByFloat(t *testing.T, expect any, get any) {
	equal(t, expect, get, Float)
}
func EqualByTime(t *testing.T, expect any, get any) {
	equal(t, expect, get, Time)
}
func EqualByList(t *testing.T, expect any, get any) {
	equal(t, expect, get, List)
}
func NotEqual(t *testing.T, notExpect any, get any) {
	file, line, funcName := getCallerInfo(1)
	tE := fmt.Sprintf("%T", notExpect)
	tG := fmt.Sprintf("%T", get)
	if tE != tG {
		printSuccess(t, file, funcName, line, tE, tG, notExpect, get)
		return
	}
	if fmt.Sprintf("%v", notExpect) == fmt.Sprintf("%v", get) {
		printFailure(t, file, funcName, line, tE, tG, notExpect, get)
		return
	}
	printSuccess(t, file, funcName, line, tE, tG, notExpect, get)
}
func True(t *testing.T, value bool) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	if !value {
		printFailure(t, file, funcName, line, tV, "bool", true, value)
		return
	}
	printSuccess(t, file, funcName, line, tV, "bool", true, value)
}
func False(t *testing.T, value bool) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	if value {
		printFailure(t, file, funcName, line, tV, "bool", false, value)
		return
	}
	printSuccess(t, file, funcName, line, tV, "bool", false, value)
}
func Nil(t *testing.T, value any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	if value != nil && !reflect.ValueOf(value).IsNil() {
		printFailure(t, file, funcName, line, tV, "nil", nil, value)
		return
	}
	printSuccess(t, file, funcName, line, tV, "nil", nil, value)
}
func NotNil(t *testing.T, value any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	if value == nil {
		printFailure(t, file, funcName, line, tV, "not nil", "not nil", value)
		return
	}
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
func Error(t *testing.T, err error) {
	file, line, funcName := getCallerInfo(1)
	if err == nil {
		printFailure(t, file, funcName, line, "error", "error", "error", nil)
		return
	}
	printSuccess(t, file, funcName, line, "error", "error", "error", err)
}
func NoError(t *testing.T, err error) {
	file, line, funcName := getCallerInfo(1)
	if err != nil {
		printFailure(t, file, funcName, line, "error", "error", nil, err)
		return
	}
	printSuccess(t, file, funcName, line, "error", "error", nil, nil)
}
func Contains(t *testing.T, str string, substr string) {
	file, line, funcName := getCallerInfo(1)
	if !strings.Contains(str, substr) {
		printFailure(t, file, funcName, line, "string", "string", "contains: "+substr, str)
		return
	}
	printSuccess(t, file, funcName, line, "string", "string", "contains: "+substr, str)
}
func NotContains(t *testing.T, str string, substr string) {
	file, line, funcName := getCallerInfo(1)
	if strings.Contains(str, substr) {
		printFailure(t, file, funcName, line, "string", "string", "not contains: "+substr, str)
		return
	}
	printSuccess(t, file, funcName, line, "string", "string", "not contains: "+substr, str)
}
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
			isEmpty = fmt.Sprintf("%v", value) == ""
		}
	}
	if !isEmpty {
		printFailure(t, file, funcName, line, tV, "empty", "empty", value)
		return
	}
	printSuccess(t, file, funcName, line, tV, "empty", "empty", value)
}
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
func Panics(t *testing.T, fn func()) {
	file, line, funcName := getCallerInfo(1)
	defer func() {
		if r := recover(); r == nil {
			printFailure(t, file, funcName, line, "panic", "panic", "should panic", "no panic")
		} else {
			printSuccess(t, file, funcName, line, "panic", "panic", "panics", r)
		}
	}()
	fn()
}
func NotPanics(t *testing.T, fn func()) {
	file, line, funcName := getCallerInfo(1)
	defer func() {
		if r := recover(); r != nil {
			printFailure(t, file, funcName, line, "no panic", "panic", "should not panic", r)
		}
	}()
	fn()
	printSuccess(t, file, funcName, line, "no panic", "no panic", "no panic", "no panic")
}
func Greater(t *testing.T, value any, expected any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	tE := fmt.Sprintf("%T", expected)
	if fmt.Sprintf("%v", value) <= fmt.Sprintf("%v", expected) {
		printFailure(t, file, funcName, line, tV, tE, value, expected)
		return
	}
	printSuccess(t, file, funcName, line, tV, tE, value, expected)
}
func GreaterOrEqual(t *testing.T, value any, expected any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	tE := fmt.Sprintf("%T", expected)
	if fmt.Sprintf("%v", value) < fmt.Sprintf("%v", expected) {
		printFailure(t, file, funcName, line, tV, tE, value, expected)
		return
	}
	printSuccess(t, file, funcName, line, tV, tE, value, expected)
}
func Less(t *testing.T, value any, expected any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	tE := fmt.Sprintf("%T", expected)
	if fmt.Sprintf("%v", value) >= fmt.Sprintf("%v", expected) {
		printFailure(t, file, funcName, line, tV, tE, value, expected)
		return
	}
	printSuccess(t, file, funcName, line, tV, tE, value, expected)
}
func LessOrEqual(t *testing.T, value any, expected any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	tE := fmt.Sprintf("%T", expected)
	if fmt.Sprintf("%v", value) > fmt.Sprintf("%v", expected) {
		printFailure(t, file, funcName, line, tV, tE, value, expected)
		return
	}
	printSuccess(t, file, funcName, line, tV, tE, value, expected)
}
func In(t *testing.T, value any, collection any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	tC := fmt.Sprintf("%T", collection)
	val := reflect.ValueOf(collection)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array && val.Kind() != reflect.Map {
		printFailure(t, file, funcName, line, tV, tC, value, collection)
		return
	}
	found := false
	for i := 0; i < val.Len(); i++ {
		if reflect.DeepEqual(value, val.Index(i).Interface()) {
			found = true
			break
		}
	}
	if !found {
		printFailure(t, file, funcName, line, tV, tC, value, collection)
		return
	}
	printSuccess(t, file, funcName, line, tV, tC, value, collection)
}
func NotIn(t *testing.T, value any, collection any) {
	file, line, funcName := getCallerInfo(1)
	tV := fmt.Sprintf("%T", value)
	tC := fmt.Sprintf("%T", collection)
	val := reflect.ValueOf(collection)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array && val.Kind() != reflect.Map {
		printFailure(t, file, funcName, line, tV, tC, value, collection)
		return
	}
	found := false
	for i := 0; i < val.Len(); i++ {
		if reflect.DeepEqual(value, val.Index(i).Interface()) {
			found = true
			break
		}
	}
	if found {
		printFailure(t, file, funcName, line, tV, tC, value, collection)
		return
	}
	printSuccess(t, file, funcName, line, tV, tC, value, collection)
}

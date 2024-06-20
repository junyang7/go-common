package _assert

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"
)

const Default = 1
const Float = 2
const Time = 3
const List = 4

func equal(t *testing.T, e any, g any, m int) {
	file := ""
	line := 0
	funcName := ""
	pc, file, line, ok := runtime.Caller(2) // 调用栈的深度
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}
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
			fmt.Println("\t", fmt.Sprintf("\033[0;31m failure \033[0m"), file, funcName, line, tE, tG, e, g)
			t.FailNow()
			return
		}
		break
	case Time:
		if !e.(time.Time).Equal(g.(time.Time)) {
			fmt.Println("\t", fmt.Sprintf("\033[0;31m failure \033[0m"), file, funcName, line, tE, tG, e, g)
			t.FailNow()
			return
		}
		break
	default:
		if fmt.Sprintf("%v", e) != fmt.Sprintf("%v", g) {
			fmt.Println("\t", fmt.Sprintf("\033[0;31m failure \033[0m"), file, funcName, line, tE, tG, e, g)
			t.FailNow()
			return
		}
		break
	}
	fmt.Println("\t", fmt.Sprintf("\033[0;32m success \033[0m"), file, funcName, line, tE, tG, e, g)
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

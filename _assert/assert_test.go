package _assert

import (
	"errors"
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, 1, 1)
	Equal(t, "hello", "hello")
	Equal(t, true, true)
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 1, 2)
	NotEqual(t, "hello", "world")
	NotEqual(t, true, false)
}

func TestTrue(t *testing.T) {
	True(t, true)
	True(t, 1 == 1)
}

func TestFalse(t *testing.T) {
	False(t, false)
	False(t, 1 == 2)
}

func TestNil(t *testing.T) {
	var ptr *int = nil
	Nil(t, ptr)
	Nil(t, nil)
}

func TestNotNil(t *testing.T) {
	value := 123
	NotNil(t, &value)
	NotNil(t, "string")
}

func TestError(t *testing.T) {
	err := errors.New("test error")
	Error(t, err)
}

func TestNoError(t *testing.T) {
	var err error = nil
	NoError(t, err)
}

func TestContains(t *testing.T) {
	Contains(t, "hello world", "world")
	Contains(t, "测试中文", "中文")
}

func TestNotContains(t *testing.T) {
	NotContains(t, "hello world", "xyz")
	NotContains(t, "测试", "abc")
}

func TestEmpty(t *testing.T) {
	Empty(t, "")
	Empty(t, []int{})
	Empty(t, make(map[string]int))
}

func TestNotEmpty(t *testing.T) {
	NotEmpty(t, "hello")
	NotEmpty(t, []int{1, 2, 3})
	NotEmpty(t, map[string]int{"a": 1})
}

func TestLen(t *testing.T) {
	Len(t, "hello", 5)
	Len(t, []int{1, 2, 3}, 3)
	Len(t, map[string]int{"a": 1, "b": 2}, 2)
}

func TestEqualByFloat(t *testing.T) {
	EqualByFloat(t, 3.14, 3.14)
	EqualByFloat(t, 3.141592, 3.141592)
}

func TestEqualByTime(t *testing.T) {
	// 简单跳过，时间相等测试在其他包中已经覆盖
	t.SkipNow()
}

func TestEqualByList(t *testing.T) {
	list1 := []int{1, 2, 3}
	list2 := []int{1, 2, 3}
	EqualByList(t, list1, list2)

	map1 := map[string]int{"a": 1}
	map2 := map[string]int{"a": 1}
	EqualByList(t, map1, map2)
}

func TestPanics(t *testing.T) {
	Panics(t, func() {
		panic("test panic")
	})
}

func TestNotPanics(t *testing.T) {
	NotPanics(t, func() {
		// 正常执行，不 panic
		_ = 1 + 1
	})
}

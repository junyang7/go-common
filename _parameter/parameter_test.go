package _parameter

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"mime/multipart"
	"testing"
)

func TestNew(t *testing.T) {
	{
		param := New("test", "value")
		_assert.NotNil(t, param)
		_assert.Equal(t, "value", param.Value())
	}
}

func TestParameter_Default(t *testing.T) {
	// value 为 nil，应该使用默认值
	{
		param := New("age", nil)
		result := param.Default(18).Int64().Value()
		_assert.Equal(t, int64(18), result)
	}

	// value 不为 nil，不应该使用默认值
	{
		param := New("age", 30)
		result := param.Default(18).Int64().Value()
		_assert.Equal(t, int64(30), result) // 保持原值
	}

	// 链式调用
	{
		param := New("name", nil)
		result := param.Default("alice").String().Value()
		_assert.Equal(t, "alice", result)
	}

	// 0 不是 nil，不使用默认值
	{
		param := New("count", 0)
		result := param.Default(10).Int64().Value()
		_assert.Equal(t, int64(0), result) // 保持 0
	}

	// 空字符串不是 nil，不使用默认值
	{
		param := New("name", "")
		result := param.Default("default").String().Value()
		_assert.Equal(t, "", result) // 保持空字符串
	}
}

func TestParameter_Int64(t *testing.T) {
	var give int64 = 123
	var expect string = "*_validator.Int64"
	get := fmt.Sprintf("%T", New("int", give).Int64())
	_assert.Equal(t, expect, get)
}

func TestParameter_String(t *testing.T) {
	var give string = "test"
	var expect string = "*_validator.String"
	get := fmt.Sprintf("%T", New("string", give).String())
	_assert.Equal(t, expect, get)
}

func TestParameter_Bool(t *testing.T) {
	var give bool = false
	var expect string = "*_validator.Bool"
	get := fmt.Sprintf("%T", New("bool", give).Bool())
	_assert.Equal(t, expect, get)
}

func TestParameter_Float64(t *testing.T) {
	var give float64 = 3.1415926
	var expect string = "*_validator.Float64"
	get := fmt.Sprintf("%T", New("float64", give).Float64())
	_assert.Equal(t, expect, get)
}

func TestParameter_File(t *testing.T) {
	// 模拟文件列表
	files := []*multipart.FileHeader{
		{Filename: "file1.jpg"},
		{Filename: "file2.png"},
	}

	// File() 返回第一个文件
	{
		param := New("avatar", files)
		file := param.File()
		_assert.NotNil(t, file)
		_assert.Equal(t, "file1.jpg", file.Filename)
	}

	// FileList() 返回全部文件
	{
		param := New("avatar", files)
		result := param.FileList()
		_assert.NotNil(t, result)
		_assert.Len(t, result, 2)
		_assert.Equal(t, "file1.jpg", result[0].Filename)
		_assert.Equal(t, "file2.png", result[1].Filename)
	}

	// 空文件列表
	{
		param := New("avatar", []*multipart.FileHeader{})
		_assert.Nil(t, param.File())
		_assert.NotNil(t, param.FileList())
		_assert.Len(t, param.FileList(), 0)
	}

	// nil
	{
		param := New("avatar", nil)
		_assert.Nil(t, param.File())
		_assert.Nil(t, param.FileList())
	}

	// 类型不匹配
	{
		param := New("avatar", "not-a-file")
		_assert.Nil(t, param.File())
		_assert.Nil(t, param.FileList())
	}
}

func TestParameter_IsNil(t *testing.T) {
	// nil
	{
		param := New("test", nil)
		_assert.True(t, param.IsNil())
	}

	// 非 nil
	{
		param := New("test", "value")
		_assert.False(t, param.IsNil())
	}

	{
		param := New("test", 0)
		_assert.False(t, param.IsNil())
	}

	{
		param := New("test", "")
		_assert.False(t, param.IsNil())
	}
}

func TestParameter_Value(t *testing.T) {
	{
		param := New("test", "value")
		_assert.Equal(t, "value", param.Value())
	}

	{
		param := New("test", 123)
		_assert.Equal(t, 123, param.Value())
	}

	{
		param := New("test", nil)
		_assert.Nil(t, param.Value())
	}
}

func BenchmarkParameter_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("test", "value")
	}
}

func BenchmarkParameter_String(b *testing.B) {
	param := New("test", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		param.String()
	}
}

func BenchmarkParameter_Int64(b *testing.B) {
	param := New("test", 123)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		param.Int64()
	}
}

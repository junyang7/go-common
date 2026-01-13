package _directory

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestExists(t *testing.T) {
	{
		var expect bool = true
		get := Exists(".")
		_assert.Equal(t, expect, get)
	}
	{
		var expect bool = false
		get := Exists("./_directory_not_exist")
		_assert.Equal(t, expect, get)
	}
}
func TestCreate(t *testing.T) {
	{
		path := "test_create"
		Delete(path)
		Create(path)
		var expect bool = true
		get := Exists(path)
		_assert.Equal(t, expect, get)
		Delete(path)
	}
}
func TestDelete(t *testing.T) {
	{
		path := "test_delete"
		Create(path)
		Delete(path)
		var expect bool = false
		get := Exists(path)
		_assert.Equal(t, expect, get)
	}
}
func TestRename(t *testing.T) {
	{
		pathOld := "test_rename_old"
		pathNew := "test_rename_new"
		Delete(pathOld)
		Delete(pathNew)
		Create(pathOld)
		Rename(pathOld, pathNew)
		var expect bool = false
		get := Exists(pathOld)
		_assert.Equal(t, expect, get)
		expect = true
		get = Exists(pathNew)
		_assert.Equal(t, expect, get)
		Delete(pathNew)
	}
}
func TestCurrent(t *testing.T) {
	{
		dir := Current()
		_assert.NotEqual(t, "", dir)
		_assert.Equal(t, true, Exists(dir))
	}
}
func TestName(t *testing.T) {
	{
		path := "/home/user/go/src/myproject/module/file.txt"
		expected := "/home/user/go/src/myproject/module"
		result := Name(path, 1)
		_assert.Equal(t, expected, result)
	}
	{
		path := "/home/user/go/src/myproject/module/file.txt"
		expected := "/home/user/go/src/myproject"
		result := Name(path, 2)
		_assert.Equal(t, expected, result)
	}
	{
		path := "/home/user/go/src/myproject/module/file.txt"
		expected := "/home/user/go/src"
		result := Name(path, 3)
		_assert.Equal(t, expected, result)
	}
	{
		path := "/home/user/go/src/myproject/module/file.txt"
		expected := "/"
		result := Name(path, 10)
		_assert.Equal(t, expected, result)
	}
	{
		path := ""
		expected := ""
		result := Name(path, 1)
		_assert.Equal(t, expected, result)
	}
	{
		path := "/home/user/go/src/myproject/module/file.txt"
		expected := path
		result := Name(path, 0)
		_assert.Equal(t, expected, result)
	}
	{
		path := "/home/user"
		expected := "/home"
		result := Name(path, 1)
		_assert.Equal(t, expected, result)
	}
	{
		path := "/home/user/go/src/myproject/module/file.txt"
		expected := path
		result := Name(path, -1)
		_assert.Equal(t, expected, result)
	}
}

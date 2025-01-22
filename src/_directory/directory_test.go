package _directory

import (
	"github.com/junyang7/go-common/src/_assert"
	"testing"
)

func TestExists(t *testing.T) {
	{
		var expect bool = true
		get := Exists("../_directory")
		_assert.Equal(t, expect, get)
	}
	{
		var expect bool = false
		get := Exists("/_directory")
		_assert.Equal(t, expect, get)
	}
}
func TestCreate(t *testing.T) {
	path := "test"
	Delete(path)
	Create(path)
	var expect bool = true
	get := Exists(path)
	_assert.Equal(t, expect, get)
}
func TestDelete(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestRename(t *testing.T) {
	pathOld := "test"
	pathNew := "test_renamed"
	if !Exists(pathOld) {
		Create(pathOld)
	}
	Delete(pathNew)
	Rename(pathOld, pathNew)
	var expect bool = true
	get := Exists(pathNew)
	_assert.Equal(t, expect, get)
}
func TestCurrent(t *testing.T) {
	// no need to test
	t.SkipNow()
}

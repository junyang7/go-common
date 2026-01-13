package _dao

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_hash"
	"strings"
	"testing"
)

func Test_Build(t *testing.T) {
	{
		path := _directory.Current() + "/dao"
		if !_directory.Exists(path) {
			_directory.Create(path)
		}
		dbName := "Db"
		tbName := "Tb"
		Build(path, dbName, tbName)
		var expect string = "4547a1bed05f9d3c70156ace9f00f0fc"
		get := _hash.Md5(_file.ReadAllAsString(path + "/" + strings.ToLower(dbName) + "_" + strings.ToLower(tbName) + ".go"))
		_assert.Equal(t, expect, get)
		_directory.Delete(path)
	}
}
func Test_BuildByAuto(t *testing.T) {
	{
		path := _directory.Current() + "/dao"
		if !_directory.Exists(path) {
			_directory.Create(path)
		}
		BuildByAuto()
		_directory.Delete(path)
	}
}

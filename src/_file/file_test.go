package _file

import (
	"github.com/junyang7/go-common/src/_assert"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestExists(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestDelete(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestReadAll(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestReadAllAsString(t *testing.T) {
	{
		path := "TestReadAllAsString"
		Delete(path)
		Write(path, path, os.ModePerm)
		var expect string = path
		get := ReadAllAsString(path)
		_assert.Equal(t, expect, get)
	}
}
func TestReadAllAsInt64(t *testing.T) {
	{
		path := "TestReadAllAsInt64"
		content := 99
		Delete(path)
		Write(path, content, os.ModePerm)
		var expect int64 = 99
		get := ReadAllAsInt64(path)
		_assert.Equal(t, expect, get)
	}
}
func TestCopy(t *testing.T) {
	{
		pathA := "TestCopy.a.txt"
		pathB := "TestCopy.b.txt"
		content := "TestCopy"
		Delete(pathA)
		Delete(pathB)
		Write(pathA, content, os.ModePerm)
		Copy(pathA, pathB)
		var expect bool = true
		get := Exists(pathB)
		_assert.Equal(t, expect, get)
	}
}
func TestRename(t *testing.T) {
	{
		pathA := "TestRename.a.txt"
		pathB := "TestRename.b.txt"
		content := "TestRename"
		Delete(pathA)
		Delete(pathB)
		Write(pathA, content, os.ModePerm)
		Rename(pathA, pathB)
		var expect bool = true
		get := Exists(pathB)
		_assert.Equal(t, expect, get)
	}
}
func TestOpen(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestFile_Truncate(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestFile_Write(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestFile_WriteAt(t *testing.T) {
	{
		path := "TestFile_WriteAt.txt"
		content := path
		var offset int64 = 10
		var size int64 = 30
		Delete(path)
		f := Open(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
		f.Truncate(size)
		f.WriteAt(content, offset)
	}
}
func TestFile_ReadAll(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestFile_ReadAllAsString(t *testing.T) {
	{
		path := "TestFile_ReadAllAsString"
		content := path
		Delete(path)
		{
			f := Open(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
			f.Write(content)
			f.Close()
		}
		{
			f := Open(path, os.O_RDONLY, os.ModePerm)
			var expect string = path
			get := f.ReadAllAsString()
			_assert.Equal(t, expect, get)
		}
	}
}
func TestFile_ReadAllAsIne64(t *testing.T) {
	{
		path := "TestFile_ReadAllAsIne64"
		content := 99
		Delete(path)
		{
			f := Open(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
			f.Write(content)
			f.Close()
		}
		{
			f := Open(path, os.O_RDONLY, os.ModePerm)
			var expect int64 = 99
			get := f.ReadAllAsInt64()
			_assert.Equal(t, expect, get)
		}
	}
}
func TestFile_Close(t *testing.T) {
	// no need to test
	t.SkipNow()
}

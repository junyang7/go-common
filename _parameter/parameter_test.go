package _parameter

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_file"
	"mime/multipart"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	{
		param := New("username", "testUser")
		_assert.Equal(t, "username", param.name)
		_assert.Equal(t, "testUser", param.value)
	}
}
func TestDefault(t *testing.T) {
	{
		param := New("username", nil).Default("defaultUser")
		_assert.Equal(t, "defaultUser", param.value)
	}
	{
		param := New("username", "existingUser").Default("defaultUser")
		_assert.Equal(t, "existingUser", param.value)
	}
}
func TestRequired(t *testing.T) {
	{
		param := New("username", nil)
		_assert.Panics(t, func() { param.Required() })
	}
	{
		param := New("username", "testUser")
		param.Required()
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
func TestInt64(t *testing.T) {
	{
		param := New("age", 25)
		validator := param.Int64()
		_assert.NotNil(t, validator)
	}
	{
		param := New("age", "invalid")
		validator := param.Int64()
		_assert.NotNil(t, validator)
	}
}
func TestString(t *testing.T) {
	{
		param := New("username", "testUser")
		validator := param.String()
		_assert.NotNil(t, validator)
	}
	{
		param := New("username", 123)
		validator := param.String()
		_assert.NotNil(t, validator)
	}
}
func TestBool(t *testing.T) {
	{
		param := New("isActive", true)
		validator := param.Bool()
		_assert.NotNil(t, validator)
	}
	{
		param := New("isActive", "invalid")
		validator := param.Bool()
		_assert.NotNil(t, validator)
	}
}
func TestFloat64(t *testing.T) {
	{
		param := New("score", 98.5)
		validator := param.Float64()
		_assert.NotNil(t, validator)
	}
	{
		param := New("score", "invalid")
		validator := param.Float64()
		_assert.NotNil(t, validator)
	}
}
func TestFile(t *testing.T) {
	{
		path := "test_write.txt"
		content := "Hello, World!"
		perm := os.FileMode(0644)
		_file.Write(path, content, perm)
		_assert.True(t, _file.Exists(path))
		fileHeader := &multipart.FileHeader{
			Filename: path,
			Size:     int64(len(content)),
		}
		param := New("file", []*multipart.FileHeader{fileHeader})
		file := param.File()
		_assert.NotNil(t, file)
		_assert.Equal(t, path, file.Filename)
		_file.Delete(path)
	}
}
func TestFileList(t *testing.T) {
	{
		path1 := "test_write_list1.txt"
		content := "File content"
		perm := os.FileMode(0644)
		_file.Write(path1, content, perm)
		_assert.True(t, _file.Exists(path1))
		path2 := "test_write_list2.txt"
		_file.Write(path2, content, perm)
		_assert.True(t, _file.Exists(path2))
		fileHeader1 := &multipart.FileHeader{
			Filename: path1,
			Size:     int64(len(content)),
		}
		fileHeader2 := &multipart.FileHeader{
			Filename: path2,
			Size:     int64(len(content)),
		}
		param := New("files", []*multipart.FileHeader{fileHeader1, fileHeader2})
		files := param.FileList()
		_assert.NotNil(t, files)
		_assert.Equal(t, 2, len(files))
		_assert.Equal(t, path1, files[0].Filename)
		_assert.Equal(t, path2, files[1].Filename)
		_file.Delete(path1)
		_file.Delete(path2)
	}
}

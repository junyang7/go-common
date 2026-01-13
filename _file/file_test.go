package _file

import (
	"github.com/junyang7/go-common/_assert"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	{
		path := "test_write.txt"
		content := "Hello, World!"
		Write(path, content)
		_assert.True(t, Exists(path))
		Delete(path)
	}
	{
		path := "test_write_empty.txt"
		content := ""
		Write(path, content)
		_assert.True(t, Exists(path))
		Delete(path)
	}
}
func TestExists(t *testing.T) {
	{
		path := "test_exists.txt"
		content := "File exists"
		Write(path, content)
		_assert.True(t, Exists(path))
		Delete(path)
	}
	{
		path := "non_existent_file.txt"
		_assert.False(t, Exists(path))
	}
}
func TestDelete(t *testing.T) {
	{
		path := "test_delete.txt"
		content := "Hello to delete"
		Write(path, content)
		Delete(path)
		_assert.False(t, Exists(path))
	}
	{
		path := "non_existent_delete.txt"
		Delete(path)
		_assert.True(t, true)
	}
}
func TestReadAll(t *testing.T) {
	{
		path := "test_read.txt"
		content := "Reading this content"
		Write(path, content)
		readContent := string(ReadAll(path))
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
	{
		path := "test_read_empty.txt"
		content := ""
		Write(path, content)
		readContent := string(ReadAll(path))
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
}
func TestReadAllAsString(t *testing.T) {
	{
		path := "test_read_string.txt"
		content := "This is a string"
		Write(path, content)
		readContent := ReadAllAsString(path)
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
}
func TestReadAllAsInt64(t *testing.T) {
	{
		path := "test_read_int64.txt"
		content := int64(123456789)
		Write(path, content)
		readContent := ReadAllAsInt64(path)
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
	{
		path := "test_read_empty_int64.txt"
		content := int64(0)
		Write(path, content)
		readContent := ReadAllAsInt64(path)
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
	{
		path := "test_read_invalid_int64.txt"
		content := "invalid int64 content"
		Write(path, content)
		readContent := ReadAllAsInt64(path)
		_assert.Equal(t, int64(0), readContent)
		Delete(path)
	}
	{
		path := "test_read_negative_int64.txt"
		content := int64(-987654321)
		Write(path, content)
		readContent := ReadAllAsInt64(path)
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
}
func TestCopy(t *testing.T) {
	{
		pathA := "TestCopy.a.txt"
		pathB := "TestCopy.b.txt"
		content := "TestCopy"
		Delete(pathA)
		Delete(pathB)
		Write(pathA, content)
		Copy(pathA, pathB)
		var expect bool = true
		get := Exists(pathB)
		_assert.Equal(t, expect, get)
		Delete(pathA)
		Delete(pathB)
	}
}
func TestRename(t *testing.T) {
	{
		pathA := "TestRename.a.txt"
		pathB := "TestRename.b.txt"
		content := "TestRename"
		Delete(pathA)
		Delete(pathB)
		Write(pathA, content)
		Rename(pathA, pathB)
		var expect bool = true
		get := Exists(pathB)
		_assert.Equal(t, expect, get)
		Delete(pathB)
	}
}
func TestOpen(t *testing.T) {
	{
		path := "test_open.txt"
		content := "Opened and written"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		readContent := string(fileObj.ReadAll())
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
}
func TestFile_Truncate(t *testing.T) {
	{
		path := "test_truncate.txt"
		content := "This content will be truncated"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.Truncate(5)
		_truncatedContent := string(fileObj.ReadAll())
		_assert.Equal(t, "This ", _truncatedContent)
		Delete(path)
	}
	{
		path := "test_truncate_empty.txt"
		content := "Content will be removed"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.Truncate(0)
		_truncatedContent := string(fileObj.ReadAll())
		_assert.Equal(t, "", _truncatedContent)
		Delete(path)
	}
	{
		path := "test_truncate_partial.txt"
		content := "This content has extra parts"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.Truncate(10)
		_truncatedContent := string(fileObj.ReadAll())
		_assert.Equal(t, "This conte", _truncatedContent)
		Delete(path)
	}
}
func TestFile_Write(t *testing.T) {
	{
		path := "test_write.txt"
		content := "This is a test write"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		writtenContent := string(fileObj.ReadAll())
		_assert.Equal(t, content, writtenContent)
		Delete(path)
	}
	{
		path := "test_write_empty.txt"
		content := ""
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		writtenContent := string(fileObj.ReadAll())
		_assert.Equal(t, content, writtenContent)
		Delete(path)
	}
}
func TestFile_WriteAt(t *testing.T) {
	{
		path := "test_writeat_basic.txt"
		content := "Hello World"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.WriteAt("Golang", 6)
		writtenContent := string(fileObj.ReadAll())
		_assert.Equal(t, "Hello Golang", writtenContent)
		fileObj.Close()
		Delete(path)
	}
	{
		path := "test_writeat_beginning.txt"
		content := "This is original text"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.WriteAt("New ", 0)
		writtenContent := string(fileObj.ReadAll())
		_assert.Equal(t, "New  is original text", writtenContent)
		fileObj.Close()
		Delete(path)
	}
	{
		path := "test_writeat_middle.txt"
		content := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.WriteAt("12345", 10)
		writtenContent := string(fileObj.ReadAll())
		expected := "ABCDEFGHIJ12345PQRSTUVWXYZ"
		_assert.Equal(t, expected, writtenContent)
		fileObj.Close()
		Delete(path)
	}
	{
		path := "test_writeat_end.txt"
		content := "Hello"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.WriteAt(" World", 5)
		writtenContent := string(fileObj.ReadAll())
		_assert.Equal(t, "Hello World", writtenContent)
		fileSize := len(writtenContent)
		_assert.Equal(t, 11, fileSize)
		fileObj.Close()
		Delete(path)
	}
	{
		path := "test_writeat_empty.txt"
		content := "Original content"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.WriteAt("", 5)
		writtenContent := string(fileObj.ReadAll())
		_assert.Equal(t, "Original content", writtenContent) // 应该保持不变
		fileObj.Close()
		Delete(path)
	}
	{
		path := "test_writeat_far_beyond.txt"
		content := "Short"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.WriteAt("End", 100)
		writtenContent := string(fileObj.ReadAll())
		_assert.Equal(t, 103, len(writtenContent))
		_assert.Equal(t, "Short", writtenContent[:5])
		fileObj.Close()
		fileContent := ReadAll(path)
		_assert.Equal(t, "End", string(fileContent[100:103]))
		Delete(path)
	}
}
func TestFile_ReadAll(t *testing.T) {
	{
		path := "test_readall.txt"
		content := "Read all content"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		readContent := string(fileObj.ReadAll())
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
}
func TestFile_ReadAllAsString(t *testing.T) {
	{
		path := "test_readall_string.txt"
		content := "Read content as string"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		readContent := fileObj.ReadAllAsString()
		_assert.Equal(t, content, readContent)
		Delete(path)
	}
}
func TestFile_ReadAllAsIne64(t *testing.T) {
	{
		path := "test_readall_int64.txt"
		content := "12345" // 用数字表示
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		readContent := fileObj.ReadAllAsInt64()
		_assert.Equal(t, int64(12345), readContent)
		Delete(path)
	}
	{
		path := "test_readall_nonint64.txt"
		content := "non-numeric content"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		readContent := fileObj.ReadAllAsInt64()
		_assert.Equal(t, int64(0), readContent)
		Delete(path)
	}
}
func TestFile_Close(t *testing.T) {
	{
		path := "test_close.txt"
		content := "File close test"
		perm := os.FileMode(0644)
		fileObj := Open(path, os.O_RDWR|os.O_CREATE, perm)
		fileObj.Write(content)
		fileObj.Close()
		defer func() {
			r := recover()
			_assert.NotNil(t, r)
		}()
		fileObj.Write("After close")
		Delete(path)
	}
}

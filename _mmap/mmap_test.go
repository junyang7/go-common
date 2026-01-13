package _mmap

import (
	"github.com/junyang7/go-common/_assert"
	"os"
	"syscall"
	"testing"
)

func TestOpen(t *testing.T) {
	{
		file, err := os.Create("testfile")
		_assert.NoError(t, err)
		defer os.Remove("testfile")
		mmap := Open(file, 0, 1024, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
		_assert.NotNil(t, mmap)
		_assert.Equal(t, file, mmap.file)
		_assert.Equal(t, 1024, len(mmap.data))
	}
}
func TestRead(t *testing.T) {
	{
		file, err := os.Create("testfile")
		_assert.NoError(t, err)
		defer os.Remove("testfile")
		data := []byte("Hello, mmap!")
		_, err = file.Write(data)
		_assert.NoError(t, err)
		mmap := Open(file, 0, len(data), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
		result := mmap.Read(0, 5)
		expected := []byte("Hello")
		_assert.Equal(t, expected, result)
	}
}
func TestAppend(t *testing.T) {
	{
		file, err := os.Create("testfile")
		_assert.NoError(t, err)
		defer os.Remove("testfile")
		mmap := Open(file, 0, 1024, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
		appendData := []byte("Hello")
		mmap.Append(appendData)
		result := mmap.Read(0, 5)
		expected := []byte("Hello")
		_assert.Equal(t, expected, result)
	}
}
func TestFlush(t *testing.T) {
	{
		file, err := os.Create("testfile")
		_assert.NoError(t, err)
		defer os.Remove("testfile")
		mmap := Open(file, 0, 1024, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
		data := []byte("Hello, mmap!")
		mmap.Append(data)
		mmap.Flush(syscall.MS_SYNC)
		file.Sync()
		content, err := os.ReadFile("testfile")
		_assert.NoError(t, err)
		_assert.Equal(t, data, content[:len(data)])
	}
}

package _directory

import (
	"github.com/junyang7/go-common/_interceptor"
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	f, err := os.Stat(path)
	if nil != err && os.IsNotExist(err) {
		return false
	}
	return f.IsDir()
}
func Create(path string) {
	if !Exists(path) {
		if err := os.MkdirAll(path, os.ModePerm); nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
	}
}
func Delete(path string) {
	if err := os.RemoveAll(path); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func Rename(old string, new string) {
	if err := os.Rename(old, new); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func Current() string {
	dir, err := os.Getwd()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return dir
}
func Name(path string, n int) string {
	if path == "" || n <= 0 {
		return path
	}
	parent := filepath.Dir(path)
	if parent == path {
		return path
	}
	return Name(parent, n-1)
}

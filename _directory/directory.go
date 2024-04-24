package _directory

import "os"

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
			panic(err)
		}
	}
}
func Delete(path string) {
	if err := os.RemoveAll(path); nil != err {
		panic(err)
	}
}
func Rename(old string, new string) {
	if err := os.Rename(old, new); nil != err {
		panic(err)
	}
}
func Current() string {
	dir, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	return dir
}

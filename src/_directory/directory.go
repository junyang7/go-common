package _directory

import "os"

func Exists(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}
func Create(path string) {
	if !Exists(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
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

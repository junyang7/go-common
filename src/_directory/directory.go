package _directory

import (
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"os"
)

func Exists(path string) bool {
	f, err := os.Stat(path)
	_interceptor.Insure(nil != err).
		CodeMessage(_codeMessage.ErrOsStat).
		Message(err.Error()).
		Do()
	return f.IsDir()
}
func Create(path string) {
	if !Exists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		_interceptor.Insure(nil != err).
			CodeMessage(_codeMessage.ErrOsMkdirAll).
			Message(err.Error()).
			Do()
	}
}
func Delete(path string) {
	err := os.RemoveAll(path)
	_interceptor.Insure(nil != err).
		CodeMessage(_codeMessage.ErrOsRemoveAll).
		Message(err.Error()).
		Do()
}
func Rename(old string, new string) {
	err := os.Rename(old, new)
	_interceptor.Insure(nil != err).
		CodeMessage(_codeMessage.ErrOsRename).
		Message(err.Error()).
		Do()
}

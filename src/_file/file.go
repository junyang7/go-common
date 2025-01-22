package _file

import (
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_interceptor"
	"io"
	"os"
)

func Write(path string, content interface{}, perm os.FileMode) {
	if err := os.WriteFile(path, _as.ByteList(content), perm); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func Exists(path string) bool {
	f, err := os.Stat(path)
	if nil != err && os.IsNotExist(err) {
		return false
	}
	return !f.IsDir()
}
func Delete(path string) {
	if err := os.RemoveAll(path); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func ReadAll(path string) []byte {
	b, err := os.ReadFile(path)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return b
}
func ReadAllAsString(filepath string) string {
	return string(ReadAll(filepath))
}
func ReadAllAsInt64(filepath string) int64 {
	return _as.Int64(ReadAll(filepath))
}
func Copy(pathA string, pathB string) {
	o, err := os.Open(pathA)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	defer o.Close()
	n, err := os.Create(pathB)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	defer n.Close()
	if _, err := io.Copy(n, o); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func Rename(old string, new string) {
	if err := os.Rename(old, new); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

type file struct {
	f *os.File
}

func Open(path string, flag int, perm os.FileMode) *file {
	this := &file{}
	f, err := os.OpenFile(path, flag, perm)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	this.f = f
	return this
}
func (this *file) Truncate(size int64) *file {
	if err := this.f.Truncate(1000); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return this
}
func (this *file) Write(content interface{}) *file {
	if _, err := this.f.Write(_as.ByteList(content)); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return this
}
func (this *file) WriteAt(content interface{}, offset int64) *file {
	if _, err := this.f.WriteAt(_as.ByteList(content), offset); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return this
}
func (this *file) ReadAll() []byte {
	if _, err := this.f.Seek(0, 0); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	b, err := io.ReadAll(this.f)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return b
}
func (this *file) ReadAllAsString() string {
	return string(this.ReadAll())
}
func (this *file) ReadAllAsInt64() int64 {
	return _as.Int64(this.ReadAll())
}
func (this *file) Close() {
	if err := this.f.Close(); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

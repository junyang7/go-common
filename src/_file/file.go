package _file

import (
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

type File struct {
	F      *os.File
	lock   *sync.Mutex
	closed bool
}

func Open(path string, flag int, perm os.FileMode) *File {
	this := &File{
		lock: &sync.Mutex{},
	}
	file, err := os.OpenFile(path, flag, perm)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsOpenFile).
		Message(err).
		Do()
	this.F = file
	return this
}
func (this *File) Close() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.closed {
		return
	}
	err := this.F.Close()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsFileClose).
		Message(err).
		Do()
	this.closed = true
}
func ContentType(filepath string) string {
	switch strings.ToLower(path.Ext(filepath)) {
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "text/javascript"
	case ".ico":
		return "image/x-icon"
	case ".jpe", ".jpeg":
		return "image/jpeg"
	case ".webp":
		return "image/webp"
	default:
		return "text/plain"
	}
}
func Copy(sourceFilepath string, targetFilePath string) {
	o := Open(sourceFilepath, os.O_RDONLY, os.ModePerm)
	defer o.Close()
	n := Open(targetFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer n.Close()
	_, err := io.Copy(n.F, o.F)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsIoCopy).
		Message(err).
		Do()
}
func Create(path string) {
	_, err := os.Create(path)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsCreate).
		Message(err).
		Do()
}
func Delete(path string) {
	err := os.RemoveAll(path)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsRemoveAll).
		Message(err).
		Do()
}
func Exists(path string) bool {
	f, err := os.Stat(path)
	if nil != err {
		return false
	}
	return !f.IsDir()
}
func (this *File) Read() []byte {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, err := this.F.Seek(0, 0)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsFileSeek).
		Message(err).
		Do()
	b, err := ioutil.ReadAll(this.F)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrIoUtilReadAll).
		Message(err).
		Do()
	return b
}
func Read(filepath string) []byte {
	b, err := os.ReadFile(filepath)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsReadFile).
		Message(err).
		Do()
	return b
}
func (this *File) ReadAsInt64() int64 {
	return _as.Int64(this.Read())
}
func ReadAsInt64(filepath string) int64 {
	return _as.Int64(Read(filepath))
}
func (this *File) ReadAsString() string {
	return string(this.Read())
}
func ReadAsString(filepath string) string {
	return string(Read(filepath))
}
func Rename(old string, new string) {
	err := os.Rename(old, new)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsRename).
		Message(err).
		Do()
}
func (this *File) Write(content interface{}) {
	this.WriteOffset(content, 0)
}
func (this *File) WriteOffset(content interface{}, offset int64) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, err := this.F.WriteAt([]byte(_as.String(content)), offset)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsFileWriteAt).
		Message(err).
		Do()
}

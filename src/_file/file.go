package _file

import (
	"github.com/junyang7/go-common/src/_as"
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
	if nil != err {
		panic(err)
	}
	this.F = file
	return this
}
func (this *File) Close() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.closed {
		return
	}
	if err := this.F.Close(); nil != err {
		panic(err)
	}
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
	if _, err := io.Copy(n.F, o.F); nil != err {
		panic(err)
	}
}
func Create(path string) {
	if _, err := os.Create(path); nil != err {
		panic(err)
	}
}
func Delete(path string) {
	if err := os.RemoveAll(path); nil != err {
		panic(err)
	}
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
	if _, err := this.F.Seek(0, 0); nil != err {
		panic(err)
	}
	b, err := ioutil.ReadAll(this.F)
	if nil != err {
		panic(err)
	}
	return b
}
func Read(filepath string) []byte {
	b, err := os.ReadFile(filepath)
	if nil != err {
		panic(err)
	}
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
	if err := os.Rename(old, new); nil != err {
		panic(err)
	}
}
func (this *File) Write(content interface{}) {
	this.WriteOffset(content, 0)
}
func (this *File) WriteOffset(content interface{}, offset int64) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, err := this.F.WriteAt([]byte(_as.String(content)), offset); err != nil {
		panic(err)
	}
}

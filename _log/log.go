package _log

import (
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_date"
	"github.com/junyang7/go-common/_datetimeMilli"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_is"
	"github.com/junyang7/go-common/_lock"
	"path/filepath"
	"strings"
)

func Debug(message ...interface{}) {
	write("debug", message...)
}
func Info(message ...interface{}) {
	write("info", message...)
}
func Warning(message ...interface{}) {
	write("warning", message...)
}
func Error(message ...interface{}) {
	write("error", message...)
}
func Write(name string, message ...interface{}) {
	write(name, message...)
}
func write(name string, messageList ...interface{}) {
	if _is.Empty(name) {
		return
	}
	file := ""
	if strings.HasPrefix(name, "/") {
		file = name + "." + _date.GetByYmd()
	} else {
		root := _conf.Get("log.root").String().Default("").Value()
		if _is.Empty(root) {
			return
		}
		if !strings.HasPrefix(root, "/") {
			return
		}
		file = root + "/" + name + "." + _date.GetByYmd()
	}
	path := filepath.Dir(file)
	if !_directory.Exists(path) {
		_directory.Create(path)
	}
	content := _datetimeMilli.Get()
	for _, message := range messageList {
		content += "\t" + _as.String(message)
	}
	content += "\n"
	lock := _lock.Get(file)
	lock.Lock()
	defer lock.Unlock()
	_file.Append(file, content)
}

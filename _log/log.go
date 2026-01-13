package _log

import (
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_date"
	"github.com/junyang7/go-common/_datetimeMilli"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_is"
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
	root := _conf.Get("log.root").String().Default("").Value()
	if _is.Empty(root) {
		return
	}
	if !strings.HasPrefix(root, "/") {
		return
	}
	if !_directory.Exists(root) {
		_directory.Create(root)
	}
	path := root + "/" + name + "." + _date.GetByYmd()
	content := _datetimeMilli.Get() + "\t" + name
	for _, message := range messageList {
		content += "\t" + _as.String(message)
	}
	content += "\n"
	_file.Append(path, content)
}

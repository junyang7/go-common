package _csv

import (
	"encoding/csv"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_list"
	"os"
)

type writer struct {
	f *os.File
	w *csv.Writer
}

func NewWriter(path string, utf8 bool) *writer {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	if utf8 {
		if _, err := f.Write([]byte("\xEF\xBB\xBF")); nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
	}
	w := csv.NewWriter(f)
	return &writer{
		f: f,
		w: w,
	}
}
func (this *writer) write(list []string) *writer {
	err := this.w.Write(list)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return this
}
func (this *writer) WriteList(list []interface{}, indexList []int) *writer {
	formatted := make([]string, 0)
	for k, v := range list {
		if _list.In(k, indexList) {
			formatted = append(formatted, _as.String(v)+"\t")
		} else {
			formatted = append(formatted, _as.String(v))
		}
	}
	return this.write(formatted)
}
func (this *writer) WriteListList(listList [][]interface{}, indexList []int) *writer {
	for _, list := range listList {
		this.WriteList(list, indexList)
	}
	return this
}
func (this *writer) WriteDict(dict map[string]interface{}, keyList []string, indexList []string) *writer {
	formatted := make([]string, 0)
	for _, k := range keyList {
		v := dict[k]
		if _list.In(k, indexList) {
			formatted = append(formatted, _as.String(v)+"\t")
		} else {
			formatted = append(formatted, _as.String(v))
		}
	}
	return this.write(formatted)
}
func (this *writer) WriteDictList(dictList []map[string]interface{}, keyList []string, indexList []string) *writer {
	for _, dict := range dictList {
		this.WriteDict(dict, keyList, indexList)
	}
	return this
}
func (this *writer) Close() {
	this.w.Flush()
	err := this.f.Close()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

type reader struct {
	f *os.File
	r *csv.Reader
}

func NewReader(path string) *reader {
	f, err := os.Open(path)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	r := csv.NewReader(f)
	return &reader{
		f: f,
		r: r,
	}
}
func (this *reader) Read() []string {
	row, err := this.r.Read()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return row
}
func (this *reader) ReadAll() [][]string {
	rowList, err := this.r.ReadAll()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return rowList
}
func (this *reader) Close() {
	if err := this.f.Close(); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

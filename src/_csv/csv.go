package _csv

import (
	"encoding/csv"
	"github.com/junyang7/go-common/src/_as"
	"io"
	"os"
)

type reader struct {
	f *os.File
	r *csv.Reader
}

func NewReader(path string) *reader {
	f, err := os.Open(path)
	if nil != err {
		panic(err)
	}
	r := csv.NewReader(f)
	return &reader{
		f: f,
		r: r,
	}
}
func (this *reader) Read() []string {
	row, err := this.r.Read()
	if nil != err && io.EOF != err {
		panic(err)
	}
	return row
}
func (this *reader) ReadAll() [][]string {
	rowList, err := this.r.ReadAll()
	if nil != err && io.EOF != err {
		panic(err)
	}
	return rowList
}
func (this *reader) Close() {
	if err := this.f.Close(); nil != err {
		panic(err)
	}
}

type writer struct {
	f *os.File
	w *csv.Writer
}

func NewWriter(path string) *writer {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if nil != err {
		panic(err)
	}
	w := csv.NewWriter(f)
	return &writer{
		f: f,
		w: w,
	}
}
func (this *writer) Write(row []interface{}) *writer {
	rowNew := make([]string, len(row))
	for k, v := range row {
		rowNew[k] = _as.String(v)
	}
	if err := this.w.Write(rowNew); nil != err {
		panic(err)
	}
	return this
}
func (this *writer) WriteAll(rowList [][]interface{}) *writer {
	for _, row := range rowList {
		this.Write(row)
	}
	return this
}
func (this *writer) Close() {
	this.w.Flush()
	if err := this.f.Close(); nil != err {
		panic(err)
	}
}

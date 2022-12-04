package _csv

import (
	"encoding/csv"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"io"
	"os"
)

type reader struct {
	f *os.File
	r *csv.Reader
}

func NewReader(path string) *reader {
	f, err := os.Open(path)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsOpen).
		Message(err.Error()).
		Do()
	r := csv.NewReader(f)
	return &reader{
		f: f,
		r: r,
	}
}
func (this *reader) Read() []string {
	row, err := this.r.Read()
	_interceptor.Insure(nil == err || io.EOF == err).
		CodeMessage(_codeMessage.ErrCsvNewReaderRead).
		Message(err.Error()).
		Do()
	return row
}
func (this *reader) ReadAll() [][]string {
	rowList, err := this.r.ReadAll()
	_interceptor.Insure(nil == err || io.EOF == err).
		CodeMessage(_codeMessage.ErrCsvNewReaderReadAll).
		Message(err.Error()).
		Do()
	return rowList
}
func (this *reader) Close() {
	err := this.f.Close()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsFileClose).
		Message(err.Error()).
		Do()
}

type writer struct {
	f *os.File
	w *csv.Writer
}

func NewWriter(path string) *writer {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsOpenFile).
		Message(err.Error()).
		Do()
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
	err := this.w.Write(rowNew)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrCsvNewWriterWrite).
		Message(err.Error()).
		Do()
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
	err := this.f.Close()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrOsFileClose).
		Message(err.Error()).
		Do()
}

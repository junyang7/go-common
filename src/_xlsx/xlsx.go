package _xlsx

import (
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_list"
	"github.com/xuri/excelize/v2"
)

type writer struct {
	path      string
	f         *excelize.File
	sheetName string
	num       int
}

func NewWriter(path string) *writer {
	return &writer{
		path:      path,
		f:         excelize.NewFile(),
		sheetName: "Sheet1",
		num:       1,
	}
}
func (this *writer) write(list []string) *writer {
	axis, err := excelize.CoordinatesToCellName(1, this.num)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	if err := this.f.SetSheetRow(this.sheetName, axis, &list); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	this.num++
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
	this.write(formatted)
	return this
}
func (this *writer) WriteDictList(dictList []map[string]interface{}, keyList []string, indexList []string) *writer {
	for _, dict := range dictList {
		this.WriteDict(dict, keyList, indexList)
	}
	return this
}
func (this *writer) Close() {
	if err := this.f.SaveAs(this.path); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	if err := this.f.Close(); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

type reader struct {
	f           *excelize.File
	sheetName   string
	rowIterator *excelize.Rows
}

func NewReader(path string) *reader {
	f, err := excelize.OpenFile(path)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return &reader{
		f:         f,
		sheetName: f.GetSheetName(0),
	}
}
func (this *reader) Read() []string {
	if nil == this.rowIterator {
		sheetIterator, err := this.f.Rows(this.sheetName)
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		this.rowIterator = sheetIterator
	}
	this.rowIterator.Next()
	row, err := this.rowIterator.Columns()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return row
}
func (this *reader) ReadAll() [][]string {
	rowList, err := this.f.GetRows(this.sheetName)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return rowList
}
func (this *reader) Close() {
	err := this.f.Close()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

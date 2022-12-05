package _xlsx

import (
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/xuri/excelize/v2"
)

type reader struct {
	f           *excelize.File
	sheetName   string
	rowIterator *excelize.Rows
}

func NewReader(path string) *reader {
	f, err := excelize.OpenFile(path)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrExcelizeOpenFile).
		Message(err).
		Do()
	return &reader{
		f:         f,
		sheetName: f.GetSheetName(0),
	}
}
func (this *reader) Read() []string {
	if nil == this.rowIterator {
		sheetIterator, err := this.f.Rows(this.sheetName)
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrExcelizeFileRows).
			Message(err).
			Do()
		this.rowIterator = sheetIterator
	}
	this.rowIterator.Next()
	row, err := this.rowIterator.Columns()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrExcelizeFileRowsRowIteratorColumns).
		Message(err).
		Do()
	return row
}
func (this *reader) ReadAll() [][]string {
	rowList, err := this.f.GetRows(this.sheetName)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrExcelizeFileGetRows).
		Message(err).
		Do()
	return rowList
}
func (this *reader) Close() {
	err := this.f.Close()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrExcelizeFileClose).
		Message(err).
		Do()
}

type writer struct {
	path      string
	f         *excelize.File
	sheetName string
	rowIndex  int
}

func NewWriter(path string) *writer {
	return &writer{
		path:      path,
		f:         excelize.NewFile(),
		sheetName: "Sheet1",
		rowIndex:  0,
	}
}
func (this *writer) Write(row []interface{}) *writer {
	this.rowIndex++
	axis, err := excelize.CoordinatesToCellName(1, this.rowIndex)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrExcelizeCoordinatesToCellName).
		Message(err).
		Do()
	{
		err := this.f.SetSheetRow(this.sheetName, axis, &row)
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrExcelizeFileCloseSetSheetRow).
			Message(err).
			Do()
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
	err := this.f.SaveAs(this.path)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrExcelizeFileSaveAs).
		Message(err).
		Do()
	{
		err := this.f.Close()
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrExcelizeFileClose).
			Message(err).
			Do()
	}
}

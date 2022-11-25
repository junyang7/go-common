package _xlsx

import "github.com/xuri/excelize/v2"

type reader struct {
	f           *excelize.File
	sheetName   string
	rowIterator *excelize.Rows
}

func NewReader(path string) *reader {
	f, err := excelize.OpenFile(path)
	if nil != err {
		panic(err)
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
			panic(err)
		}
		this.rowIterator = sheetIterator
	}
	this.rowIterator.Next()
	row, err := this.rowIterator.Columns()
	if nil != err {
		panic(err)
	}
	return row
}
func (this *reader) ReadAll() [][]string {
	rowList, err := this.f.GetRows(this.sheetName)
	if nil != err {
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
	if nil != err {
		panic(err)
	}
	if err := this.f.SetSheetRow(this.sheetName, axis, &row); nil != err {
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
	if err := this.f.SaveAs(this.path); nil != err {
		panic(err)
	}
	if err := this.f.Close(); nil != err {
		panic(err)
	}
}

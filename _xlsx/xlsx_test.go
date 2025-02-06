package _xlsx

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_hash"
	"github.com/junyang7/go-common/_list"
	"testing"
)

func TestNewWriter(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestWriter_WriteList(t *testing.T) {
	{
		path := "test.xlsx"
		_file.Delete(path)
		xlsxWriter := NewWriter(path)
		{
			list := []interface{}{1, "A", "b", 3.141592}
			xlsxWriter.WriteList(list, []int{})
		}
		xlsxWriter.Close()
		var expect string = "7d07677d7c7f42e128ffe20472029119"
		get := _hash.Md5(_file.ReadAllAsString(path))
		_assert.Equal(t, expect, get)
	}
}
func TestWriter_WriteListList(t *testing.T) {
	{
		path := "test.xlsx"
		_file.Delete(path)
		xlsxWriter := NewWriter(path)
		listList := [][]interface{}{{1, "A", "b", 3.141592}, {2, "a", "B", 3.141592}}
		xlsxWriter.WriteListList(listList, []int{})
		xlsxWriter.Close()
		var expect string = "658d84557c3a90826c7e7b033bd64a42"
		get := _hash.Md5(_file.ReadAllAsString(path))
		_assert.Equal(t, expect, get)
	}
}
func TestWriter_WriteDict(t *testing.T) {
	{
		path := "test.xlsx"
		_file.Delete(path)
		xlsxWriter := NewWriter(path)
		keyList := []string{"1", "2", "3", "4"}
		indexList := []string{}
		{
			dict := map[string]interface{}{"1": 1, "2": "A", "3": "b", "4": 3.141592}
			xlsxWriter.WriteDict(dict, keyList, indexList)
		}
		xlsxWriter.Close()
		var expect string = "7d07677d7c7f42e128ffe20472029119"
		get := _hash.Md5(_file.ReadAllAsString(path))
		_assert.Equal(t, expect, get)
	}
}
func TestWriter_WriteDictList(t *testing.T) {
	{
		path := "test.xlsx"
		_file.Delete(path)
		xlsxWriter := NewWriter(path)
		keyList := []string{"1", "2", "3", "4"}
		indexList := []string{}
		dictList := []map[string]interface{}{{"1": 1, "2": "A", "3": "b", "4": 3.141592}, {"1": 2, "2": "a", "3": "B", "4": 3.141592}}
		xlsxWriter.WriteDictList(dictList, keyList, indexList)
		xlsxWriter.Close()
		var expect string = "658d84557c3a90826c7e7b033bd64a42"
		get := _hash.Md5(_file.ReadAllAsString(path))
		_assert.Equal(t, expect, get)
	}
}
func TestWriter_Close(t *testing.T) {
	// no need to skip
	t.SkipNow()
}
func TestNewReader(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestReader_Read(t *testing.T) {
	{
		var expect string = "42c90625a6493ea636c92245650cd985"
		path := "test.xlsx"
		xlsxReader := NewReader(path)
		list := xlsxReader.Read()
		xlsxReader.Close()
		get := _hash.Md5(_list.Implode("", list))
		_assert.Equal(t, expect, get)
	}
}
func TestReader_ReadAll(t *testing.T) {
	{
		var expect string = "2d6d122d730a3d74a7769d3c491cb6c1"
		path := "test.xlsx"
		xlsxReader := NewReader(path)
		listList := xlsxReader.ReadAll()
		xlsxReader.Close()
		content := ""
		for _, list := range listList {
			content += _list.Implode("", list)
		}
		get := _hash.Md5(content)
		_assert.Equal(t, expect, get)
	}
}
func TestReader_Close(t *testing.T) {
	// no need to test
	t.SkipNow()
}

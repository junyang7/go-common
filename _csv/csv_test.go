package _csv

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_hash"
	"github.com/junyang7/go-common/_list"
	"strings"
	"testing"
)

func TestNewWriter(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestWriter_WriteList(t *testing.T) {
	{
		path := "test.csv"
		_file.Delete(path)
		csvWriter := NewWriter(path, false)
		{
			list := []interface{}{1, "A", "b", 3.141592}
			csvWriter.WriteList(list, []int{})
		}
		csvWriter.Close()
		var expect string = "289120e5deda605360daab537b538306"
		get := _hash.Md5(_file.ReadAllAsString(path))
		_assert.Equal(t, expect, get)
	}
}
func TestWriter_WriteListList(t *testing.T) {
	{
		path := "test.csv"
		_file.Delete(path)
		csvWriter := NewWriter(path, false)
		listList := [][]interface{}{{1, "A", "b", 3.141592}, {2, "a", "B", 3.141592}}
		csvWriter.WriteListList(listList, []int{})
		csvWriter.Close()
		var expect string = "e5c07306e9b0e8e6e214bf2b0d371d4f"
		get := _hash.Md5(_file.ReadAllAsString(path))
		_assert.Equal(t, expect, get)
	}
}
func TestWriter_WriteDict(t *testing.T) {
	{
		path := "test.csv"
		_file.Delete(path)
		csvWriter := NewWriter(path, false)
		keyList := []string{"1", "2", "3", "4"}
		indexList := []string{}
		{
			dict := map[string]interface{}{"1": 1, "2": "A", "3": "b", "4": 3.141592}
			csvWriter.WriteDict(dict, keyList, indexList)
		}
		csvWriter.Close()
		var expect string = "289120e5deda605360daab537b538306"
		get := _hash.Md5(_file.ReadAllAsString(path))
		_assert.Equal(t, expect, get)
	}
}
func TestWriter_WriteDictList(t *testing.T) {
	{
		path := "test.csv"
		_file.Delete(path)
		csvWriter := NewWriter(path, false)
		keyList := []string{"1", "2", "3", "4"}
		indexList := []string{}
		dictList := []map[string]interface{}{{"1": 1, "2": "A", "3": "b", "4": 3.141592}, {"1": 2, "2": "a", "3": "B", "4": 3.141592}}
		csvWriter.WriteDictList(dictList, keyList, indexList)
		csvWriter.Close()
		var expect string = "e5c07306e9b0e8e6e214bf2b0d371d4f"
		get := _hash.Md5(_file.ReadAllAsString(path))
		_assert.Equal(t, expect, get)
	}
}
func TestWriter_Close(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestNewReader(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestReader_Read(t *testing.T) {
	{
		var expect string = "42c90625a6493ea636c92245650cd985"
		path := "test.csv"
		csvReader := NewReader(path)
		list := csvReader.Read()
		csvReader.Close()
		fmt.Println(_list.Implode("", list))
		fmt.Println(len(strings.TrimSpace(_list.Implode("", list))))
		get := _hash.Md5(_list.Implode("", list))
		_assert.Equal(t, expect, get)
	}
}
func TestReader_ReadAll(t *testing.T) {
	{
		var expect string = "2d6d122d730a3d74a7769d3c491cb6c1"
		path := "test.csv"
		csvReader := NewReader(path)
		listList := csvReader.ReadAll()
		csvReader.Close()
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

package _sign

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_hash"
	"testing"
)

func TestMd5(t *testing.T) {
	{
		dataList := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		key := "secretKey"
		var filterLabelList []string
		var filterValueList []string
		expected := _hash.Md5("key1=value1&key2=value2" + key)
		sign := Md5(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
	{
		dataList := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		key := "secretKey"
		filterLabelList := []string{"key1"}
		filterValueList := []string{}
		expected := _hash.Md5("key2=value2&key3=value3" + key)
		sign := Md5(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
	{
		dataList := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		key := "secretKey"
		filterLabelList := []string{}
		filterValueList := []string{"value2"}
		expected := _hash.Md5("key1=value1&key3=value3" + key)
		sign := Md5(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
	{
		dataList := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		key := "secretKey"
		filterLabelList := []string{}
		filterValueList := []string{}
		expected := _hash.Md5("key1=value1&key2=value2" + key)
		sign := Md5(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
	{
		dataList := map[string]string{}
		key := "secretKey"
		filterLabelList := []string{}
		filterValueList := []string{}
		expected := _hash.Md5(key)
		sign := Md5(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
}
func TestHmacSha1(t *testing.T) {
	{
		dataList := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		key := "secretKey"
		var filterLabelList []string
		var filterValueList []string
		expected := _hash.HmacSha1("key1=value1&key2=value2", key)
		sign := HmacSha1(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
	{
		dataList := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		key := "secretKey"
		filterLabelList := []string{"key1"}
		filterValueList := []string{}
		expected := _hash.HmacSha1("key2=value2&key3=value3", key)
		sign := HmacSha1(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
	{
		dataList := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		key := "secretKey"
		filterLabelList := []string{}
		filterValueList := []string{"value2"}
		expected := _hash.HmacSha1("key1=value1&key3=value3", key)
		sign := HmacSha1(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
	{
		dataList := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		key := "secretKey"
		filterLabelList := []string{}
		filterValueList := []string{}
		expected := _hash.HmacSha1("key1=value1&key2=value2", key)
		sign := HmacSha1(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
	{
		dataList := map[string]string{}
		key := "secretKey"
		filterLabelList := []string{}
		filterValueList := []string{}
		expected := _hash.HmacSha1("", key)
		sign := HmacSha1(dataList, key, filterLabelList, filterValueList)
		_assert.Equal(t, expected, sign)
	}
}

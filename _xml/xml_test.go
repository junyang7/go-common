package _xml

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

type TestStruct struct {
	Name  string `xml:"name"`
	Value int    `xml:"value"`
}

func TestEncode(t *testing.T) {
	{
		data := TestStruct{Name: "Test", Value: 123}
		result := Encode(data)
		expected := `<TestStruct><name>Test</name><value>123</value></TestStruct>`
		_assert.Equal(t, string(result), expected)
	}
	{
		data := TestStruct{}
		result := Encode(data)
		expected := `<TestStruct><name></name><value>0</value></TestStruct>`
		_assert.Equal(t, string(result), expected)
	}
	{
		var data *TestStruct
		result := Encode(data)
		_assert.Equal(t, string(result), ``)
	}
	{
		result := Encode("invalid data type")
		_assert.Equal(t, string(result), `<string>invalid data type</string>`)
	}
}
func TestEncodeAsString(t *testing.T) {
	{
		data := TestStruct{Name: "Test", Value: 123}
		result := EncodeAsString(data)
		expected := `<TestStruct><name>Test</name><value>123</value></TestStruct>`
		_assert.Equal(t, result, expected)
	}
	{
		data := TestStruct{}
		result := EncodeAsString(data)
		expected := `<TestStruct><name></name><value>0</value></TestStruct>`
		_assert.Equal(t, result, expected)
	}
	{
		var data *TestStruct
		result := EncodeAsString(data)
		_assert.Equal(t, result, ``)
	}
	{
		result := EncodeAsString("invalid data type")
		_assert.Equal(t, result, `<string>invalid data type</string>`)
	}
}
func TestDecode(t *testing.T) {
	{
		xmlData := []byte(`<TestStruct><name>Test</name><value>123</value></TestStruct>`)
		var result TestStruct
		Decode(xmlData, &result)
		_assert.Equal(t, result.Name, "Test")
		_assert.Equal(t, result.Value, 123)
	}
	{
		xmlData := []byte(`<TestStruct><name>Test</name></TestStruct>`)
		var result TestStruct
		Decode(xmlData, &result)
		_assert.Equal(t, result.Name, "Test")
		_assert.Equal(t, result.Value, 0)
	}
	{
		xmlData := []byte(`<TestStruct><name>Test</name><value>123</value><extra>Extra</extra></TestStruct>`)
		var result TestStruct
		Decode(xmlData, &result)
		_assert.Equal(t, result.Name, "Test")
		_assert.Equal(t, result.Value, 123)
	}
	{
		xmlData := []byte(`<TestStruct><name>Test<value>123</value></TestStruct>`)
		var result TestStruct
		_assert.Panics(t, func() {
			Decode(xmlData, &result)
		})
	}
	{
		xmlData := []byte(`<TestStruct><name>Test</name><value>123</value></TestStruct>`)
		var result string
		Decode(xmlData, &result)
		_assert.Empty(t, result)
	}
}

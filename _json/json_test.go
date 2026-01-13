package _json

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_file"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestEncode(t *testing.T) {
	{
		person := Person{Name: "Alice", Age: 30}
		result := Encode(person)
		expected := `{"Name":"Alice","Age":30}`
		_assert.Equal(t, expected, string(result))
	}
	{
		data := map[string]interface{}{
			"Name": "Bob",
			"Age":  25,
		}
		result := Encode(data)
		expected := `{"Age":25,"Name":"Bob"}`
		_assert.Equal(t, expected, string(result))
	}
}
func TestEncodeAsString(t *testing.T) {
	{
		person := Person{Name: "Alice", Age: 30}
		result := EncodeAsString(person)
		expected := `{"Name":"Alice","Age":30}`
		_assert.Equal(t, expected, result)
	}
	{
		data := map[string]interface{}{
			"Name": "Bob",
			"Age":  25,
		}
		result := EncodeAsString(data)
		expected := `{"Age":25,"Name":"Bob"}`
		_assert.Equal(t, expected, result)
	}
}
func TestDecode(t *testing.T) {
	path := "test_decode.json"
	data := `{"Name":"Alice","Age":30}`
	_file.Write(path, data)
	_assert.True(t, _file.Exists(path))
	defer _file.Delete(path)
	{
		var result Person
		DecodeByFile(path, &result)
		_assert.Equal(t, "Alice", result.Name)
		_assert.Equal(t, 30, result.Age)
	}
	{
		var result map[string]interface{}
		DecodeByFile(path, &result)
		_assert.Equal(t, "Alice", result["Name"])
		_assert.Equal(t, 30.0, result["Age"])
	}
}
func TestDecodeByFile(t *testing.T) {
	path := "test_decode_file.json"
	data := `{"Name":"Bob","Age":25}`
	_file.Write(path, data)
	_assert.True(t, _file.Exists(path))
	defer _file.Delete(path)
	{
		var result map[string]interface{}
		DecodeByFile(path, &result)
		_assert.Equal(t, "Bob", result["Name"])
		_assert.Equal(t, 25.0, result["Age"])
	}
}
func TestDecodeByText(t *testing.T) {
	{
		var result map[string]interface{}
		DecodeByText(`{"Name":"Bob","Age":25}`, &result)
		_assert.Equal(t, "Bob", result["Name"])
		_assert.Equal(t, 25.0, result["Age"])
	}
}
func TestReaderByte(t *testing.T) {
	path := "test_reader_byte.json"
	data := `{"name":"Alice","age":30}`
	_file.Write(path, data)
	_assert.True(t, _file.Exists(path))
	defer _file.Delete(path)
	{
		conf := New().File(path)
		result := conf.Get("name")
		_assert.Equal(t, "Alice", result.Value())
		result = conf.Get("age")
		_assert.Equal(t, 30.0, result.Value())
	}
}
func TestReaderText(t *testing.T) {
	{
		data := `{"name":"Bob","age":25}`
		conf := New().Text(data)
		result := conf.Get("name")
		_assert.Equal(t, "Bob", result.Value())
	}
}
func TestReaderFile(t *testing.T) {
	path := "test_reader_file.json"
	data := `{"name":"Charlie","age":35}`
	_file.Write(path, data)
	_assert.True(t, _file.Exists(path))
	defer _file.Delete(path)
	{
		conf := New().File(path)
		result := conf.Get("name")
		_assert.Equal(t, "Charlie", result.Value())
		result = conf.Get("age")
		_assert.Equal(t, 35.0, result.Value())
	}
}
func TestReaderGet(t *testing.T) {
	path := "test_reader_get.json"
	data := `{"user": {"name": "Alice", "age": 30}}`
	_file.Write(path, data)
	_assert.True(t, _file.Exists(path))
	defer _file.Delete(path)
	{
		conf := New().File(path)
		result := conf.Get("user.name")
		_assert.Equal(t, "Alice", result.Value())
		result = conf.Get("user.age")
		_assert.Equal(t, 30.0, result.Value())
	}
	{
		data := `{"users": [{"name": "Alice"}, {"name": "Bob"}]}`
		_file.Write(path, data)
		conf := New().File(path)
		result := conf.Get("users.0.name")
		_assert.Equal(t, "Alice", result.Value())
		result = conf.Get("users.1.name")
		_assert.Equal(t, "Bob", result.Value())
	}
}

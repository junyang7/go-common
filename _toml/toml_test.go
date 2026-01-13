package _toml

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_hash"
	"testing"
)

func TestEncode(t *testing.T) {
	{
		data := map[string]interface{}{
			"Name":    "Tom",
			"Age":     30,
			"Country": "China",
		}
		encoded := Encode(data)
		_assert.NotNil(t, encoded)
		_assert.True(t, len(encoded) > 0)
	}
	{
		var data interface{}
		encoded := Encode(data)
		_assert.Nil(t, encoded)
		_assert.Equal(t, len(encoded), 0)
	}
}
func TestEncodeAsString(t *testing.T) {
	{
		data := map[string]interface{}{
			"Name":    "Tom",
			"Age":     30,
			"Country": "China",
		}
		expect := _hash.Md5(`Age = 30
Country = "China"
Name = "Tom"
`)
		get := _hash.Md5(EncodeAsString(data))
		_assert.Equal(t, expect, get)
	}
}
func TestDecode(t *testing.T) {
	{
		source := []byte(`
			[person]
			name = "Tom"
			age = 30
			`)
		var target map[string]interface{}
		Decode(source, &target)
		_assert.Equal(t, target["person"].(map[string]interface{})["name"], "Tom")
		_assert.Equal(t, target["person"].(map[string]interface{})["age"], int64(30))
	}
	{
		source := []byte(`invalid toml content`)
		var target map[string]interface{}
		defer func() {
			err := recover()
			_assert.NotNil(t, err)
		}()
		Decode(source, &target)
	}
}
func TestNew(t *testing.T) {
	{
		conf := New()
		_assert.NotNil(t, conf)
	}
}
func TestReaderByte(t *testing.T) {
	{
		data := []byte(`
			[person]
			name = "Tom"
			age = 30
		`)
		conf := New().Byte(data)
		param := conf.Get("person.name")
		_assert.NotNil(t, param)
		_assert.Equal(t, param.String().Value(), "Tom")
	}
	{
		defer func() {
			err := recover()
			_assert.NotNil(t, err)
		}()
		data := []byte(`invalid content`)
		New().Byte(data)
	}
}
func TestReaderText(t *testing.T) {
	{
		data := `
			[person]
			name = "Tom"
			age = 30
		`
		conf := New().Text(data)
		param := conf.Get("person.name")
		_assert.NotNil(t, param)
		_assert.Equal(t, param.String().Value(), "Tom")
	}
	{
		defer func() {
			err := recover()
			_assert.NotNil(t, err)
		}()
		data := `invalid text content`
		New().Text(data)
	}
}
func TestReaderFile(t *testing.T) {
	{
		path := "test.toml"
		_file.Write(path, `[person]
name = "Tom"
age = 30`)
		conf := New().File(path)
		param := conf.Get("person.name")
		_assert.NotNil(t, param)
		_assert.Equal(t, param.String().Value(), "Tom")
		_file.Delete(path)
	}
	{
		path := "test.toml"
		_file.Write(path, `invalid content`)
		defer func() {
			err := recover()
			_assert.NotNil(t, err)
			_file.Delete(path)
		}()
		New().File(path)
	}
}
func TestReaderGet(t *testing.T) {
	{
		data := []byte(`
			[person]
			name = "Tom"
			age = 30
		`)
		conf := New().Byte(data)
		param := conf.Get("person.name")
		_assert.NotNil(t, param)
		_assert.Equal(t, param.String().Value(), "Tom")
	}
	{
		data := []byte(`
			[person]
			name = "Tom"
		`)
		conf := New().Byte(data)
		param := conf.Get("age")
		_assert.Equal(t, param.Value(), nil)
	}
	{
		data := []byte(`
			[people]
			[[people.person]]
			name = "Tom"
			age = 30
			[[people.person]]
			name = "Jerry"
			age = 25
		`)
		conf := New().Byte(data)
		param := conf.Get("people.person.1.name")
		_assert.NotNil(t, param)
		_assert.Equal(t, param.String().Value(), "Jerry")
	}
	{
		data := []byte(`
			[people]
			[[people.person]]
			name = "Tom"
			age = 30
		`)
		conf := New().Byte(data)
		param := conf.Get("people.person.1.name")
		_assert.Equal(t, param.Value(), nil)
	}
}

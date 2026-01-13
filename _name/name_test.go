package _name

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestUpperCamelCase(t *testing.T) {
	{
		get := UpperCamelCase("my_name_is_ok")
		expect := "MyNameIsOk"
		_assert.Equal(t, expect, get)
	}
	{
		get := UpperCamelCase("MyNameIsOk")
		expect := "MyNameIsOk"
		_assert.Equal(t, expect, get)
	}
	{
		get := UpperCamelCase("myNameIsOk")
		expect := "MyNameIsOk"
		_assert.Equal(t, expect, get)
	}
	{
		get := UpperCamelCase("my-name-is-ok")
		expect := "MyNameIsOk"
		_assert.Equal(t, expect, get)
	}
}
func TestLowerCamelCase(t *testing.T) {
	{
		get := LowerCamelCase("my_name_is_ok")
		expect := "myNameIsOk"
		_assert.Equal(t, expect, get)
	}
	{
		get := LowerCamelCase("MyNameIsOk")
		expect := "myNameIsOk"
		_assert.Equal(t, expect, get)
	}
	{
		get := LowerCamelCase("myNameIsOk")
		expect := "myNameIsOk"
		_assert.Equal(t, expect, get)
	}
	{
		get := LowerCamelCase("my-name-is-ok")
		expect := "myNameIsOk"
		_assert.Equal(t, expect, get)
	}
}
func TestSnakeCase(t *testing.T) {
	{
		get := SnakeCase("my_name_is_ok")
		expect := "my_name_is_ok"
		_assert.Equal(t, expect, get)
	}
	{
		get := SnakeCase("MyNameIsOk")
		expect := "my_name_is_ok"
		_assert.Equal(t, expect, get)
	}
	{
		get := SnakeCase("myNameIsOk")
		expect := "my_name_is_ok"
		_assert.Equal(t, expect, get)
	}
	{
		get := SnakeCase("my-name-is-ok")
		expect := "my_name_is_ok"
		_assert.Equal(t, expect, get)
	}
}

package _string

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestPad(t *testing.T) {
	{
		var s = "hello"
		var length = 10
		var sp = "*"
		var expect = "*****hello"
		get := Pad(s, length, sp, L)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello"
		var length = 10
		var sp = "*"
		var expect = "hello*****"
		get := Pad(s, length, sp, R)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello"
		var length = 5
		var sp = "*"
		var expect = "hello"
		get := Pad(s, length, sp, L)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello"
		var length = 10
		var sp = ""
		var expect = "hello"
		get := Pad(s, length, sp, L)
		_assert.Equal(t, expect, get)
	}
}
func TestPadLeft(t *testing.T) {
	{
		var s = "hello"
		var length = 10
		var sp = "*"
		var expect = "*****hello"
		get := PadLeft(s, length, sp)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello"
		var length = 5
		var sp = "*"
		var expect = "hello"
		get := PadLeft(s, length, sp)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello"
		var length = 10
		var sp = ""
		var expect = "hello"
		get := PadLeft(s, length, sp)
		_assert.Equal(t, expect, get)
	}
}
func TestPadRight(t *testing.T) {
	{
		var s = "hello"
		var length = 10
		var sp = "*"
		var expect = "hello*****"
		get := PadRight(s, length, sp)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello"
		var length = 5
		var sp = "*"
		var expect = "hello"
		get := PadRight(s, length, sp)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello"
		var length = 10
		var sp = ""
		var expect = "hello"
		get := PadRight(s, length, sp)
		_assert.Equal(t, expect, get)
	}
}
func TestReplaceAll(t *testing.T) {
	{
		var s = "hello world"
		var old = "o"
		var newStr = "0"
		var expect = "hell0 w0rld"
		get := ReplaceAll(s, old, newStr)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello world"
		var old = "z"
		var newStr = "x"
		var expect = "hello world"
		get := ReplaceAll(s, old, newStr)
		_assert.Equal(t, expect, get)
	}
	{
		var s = "hello world"
		var old = " "
		var newStr = ""
		var expect = "helloworld"
		get := ReplaceAll(s, old, newStr)
		_assert.Equal(t, expect, get)
	}
}

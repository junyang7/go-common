package _codeMessage

import (
	"github.com/junyang7/go-common/src/_assert"
	"testing"
)

func TestNew(t *testing.T) {
	{
		var code int = 0
		var message string = "success"
		cm := New(code, message)
		{
			var expect int = 0
			get := cm.Code
			_assert.Equal(t, expect, get)
		}
		{
			var expect string = "success"
			get := cm.Message
			_assert.Equal(t, expect, get)
		}
	}
}
func TestErrNone(t *testing.T) {
	{
		var expect int = 0
		get := ErrNone.Code
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "success"
		get := ErrNone.Message
		_assert.Equal(t, expect, get)
	}
}
func TestErrDefault(t *testing.T) {
	{
		var expect int = 1
		get := ErrDefault.Code
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "something goes wrong...!!!"
		get := ErrDefault.Message
		_assert.Equal(t, expect, get)
	}
}

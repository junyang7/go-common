package _exception

import (
	"fmt"
	"github.com/junyang7/go-common/src/_assert"
	"github.com/junyang7/go-common/src/_codeMessage"
	"testing"
)

func TestNew(t *testing.T) {
	n := New()
	{
		var expect int = _codeMessage.ErrDefault.Code
		get := n.Code
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = _codeMessage.ErrDefault.Message
		get := n.Message
		_assert.Equal(t, expect, get)
	}
}
func TestException_Throw(t *testing.T) {
	defer func() {
		err := recover()
		{
			var expect string = "*_exception.Exception"
			get := fmt.Sprintf("%T", err)
			_assert.Equal(t, expect, get)
		}
		{
			e := err.(*Exception)
			{
				var expect int = _codeMessage.ErrDefault.Code
				get := e.Code
				_assert.Equal(t, expect, get)
			}
			{
				var expect string = _codeMessage.ErrDefault.Message
				get := e.Message
				_assert.Equal(t, expect, get)
			}
			{
				var expect string = "you know!"
				get := e.Data.(string)
				_assert.Equal(t, expect, get)
			}
		}
	}()
	e := New()
	e.Data = "you know!"
	e.Throw()
}

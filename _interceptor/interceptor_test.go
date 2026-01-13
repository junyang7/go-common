package _interceptor

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_codeMessage"
	"github.com/junyang7/go-common/_exception"
	"testing"
)

func TestInsure(t *testing.T) {
	{
		interceptor := Insure(true)
		_assert.Equal(t, interceptor.ok, true)
		_assert.Equal(t, interceptor.code, _codeMessage.ErrDefault.Code)
		_assert.Equal(t, interceptor.message, _codeMessage.ErrDefault.Message)
	}
	{
		interceptor := Insure(false)
		_assert.Equal(t, interceptor.ok, false)
		_assert.Equal(t, interceptor.code, _codeMessage.ErrDefault.Code)
		_assert.Equal(t, interceptor.message, _codeMessage.ErrDefault.Message)
	}
}
func TestInterceptor_Code(t *testing.T) {
	{
		interceptor := Insure(false).Code(404)
		_assert.Equal(t, interceptor.code, 404)
	}
	{
		interceptor := Insure(false).Code(500)
		_assert.Equal(t, interceptor.code, 500)
	}
}
func TestInterceptor_Message(t *testing.T) {
	{
		interceptor := Insure(false).Message("Error occurred")
		_assert.Equal(t, interceptor.message, "Error occurred")
	}
	{
		interceptor := Insure(false).Message(123)
		_assert.Equal(t, interceptor.message, "123")
	}
	{
		interceptor := Insure(false).Message(true)
		_assert.Equal(t, interceptor.message, "true")
	}
}
func TestInterceptor_Data(t *testing.T) {
	{
		interceptor := Insure(false).Data("Some data")
		_assert.Equal(t, interceptor.data, "Some data")
	}
	{
		interceptor := Insure(false).Data(123)
		_assert.Equal(t, interceptor.data, 123)
	}
	{
		interceptor := Insure(false).Data(struct{}{})
		_assert.Equal(t, interceptor.data, struct{}{})
	}
}
func TestInterceptor_CodeMessage(t *testing.T) {
	{
		codeMessage := &_codeMessage.CodeMessage{
			Code:    403,
			Message: "Forbidden",
		}

		interceptor := Insure(false).CodeMessage(codeMessage)
		_assert.Equal(t, interceptor.code, 403)
		_assert.Equal(t, interceptor.message, "Forbidden")
	}
	{
		codeMessage := &_codeMessage.CodeMessage{
			Code:    400,
			Message: "Bad Request",
		}

		interceptor := Insure(false).CodeMessage(codeMessage)
		_assert.Equal(t, interceptor.code, 400)
		_assert.Equal(t, interceptor.message, "Bad Request")
	}
}
func TestInterceptor_Do(t *testing.T) {
	{
		Insure(true).
			Code(100).
			Message("it's just a test!").
			Data(map[string]interface{}{"data": "something ..."}).
			Do()
	}
	{
		defer func() {
			err := recover()
			_assert.NotNil(t, err)
			_assert.Equal(t, "*_exception.Exception", fmt.Sprintf("%T", err))
			e, ok := err.(*_exception.Exception)
			_assert.True(t, ok)
			_assert.Equal(t, 100, e.Code)
			_assert.Equal(t, "It's just a test!", e.Message)
			_assert.Equal(t, map[string]interface{}{"data": "something ..."}, e.Data.(map[string]interface{}))
		}()
		Insure(false).
			Code(100).
			Message("It's just a test!").
			Data(map[string]interface{}{"data": "something ..."}).
			Do()
	}
}

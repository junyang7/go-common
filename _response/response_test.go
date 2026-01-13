package _response

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_codeMessage"
	"testing"
)

func TestResponse_New(t *testing.T) {
	{
		resp := New()
		_assert.Equal(t, _codeMessage.ErrNone.Code, resp.Code)
		_assert.Equal(t, _codeMessage.ErrNone.Message, resp.Message)
		_assert.NotNil(t, resp.Data)
		_assert.Equal(t, int64(0), resp.Time)
		_assert.Equal(t, int64(0), resp.Consume)
	}
	{
		resp := &Response{
			Code:    404,
			Message: "Not Found",
			Data:    nil,
		}
		_assert.Equal(t, 404, resp.Code)
		_assert.Equal(t, "Not Found", resp.Message)
		_assert.Nil(t, resp.Data)
	}
	{
		resp := &Response{
			Code:    500,
			Message: "Internal Server Error",
			File:    "test_file.go",
			Line:    42,
		}
		_assert.Equal(t, "test_file.go", resp.File)
		_assert.Equal(t, 42, resp.Line)
	}
}
func TestResponse_FieldVerification(t *testing.T) {
	{
		resp := New()
		resp.Code = 200
		resp.Message = "OK"
		resp.Data = map[string]interface{}{"key": "value"}
		resp.Time = 1234567890
		resp.Consume = 1500
		_assert.Equal(t, 200, resp.Code)
		_assert.Equal(t, "OK", resp.Message)
		_assert.Equal(t, map[string]interface{}{"key": "value"}, resp.Data)
		_assert.Equal(t, int64(1234567890), resp.Time)
		_assert.Equal(t, int64(1500), resp.Consume)
	}
}
func TestResponse_EmptyData(t *testing.T) {
	{
		resp := &Response{
			Code:    200,
			Message: "OK",
			Data:    nil,
		}
		_assert.Nil(t, resp.Data)
	}
}

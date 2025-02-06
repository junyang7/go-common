package _interceptor

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_exception"
	"testing"
)

func TestInsure(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestInterceptor_Code(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestInterceptor_Message(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestInterceptor_Data(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestInterceptor_CodeMessage(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestInterceptor_Do(t *testing.T) {
	(func() {
		Insure(true).
			Code(100).
			Message("it's just a test!").
			Data(map[string]interface{}{"data": "something ..."}).
			Do()
	})()
	(func() {
		defer func() {
			err := recover()
			{
				var expect string = "*_exception.Exception"
				get := fmt.Sprintf("%T", err)
				_assert.Equal(t, expect, get)
			}
			{
				e := err.(*_exception.Exception)
				{
					var expect int = 100
					get := e.Code
					_assert.Equal(t, expect, get)
				}
				{
					var expect string = "It's just a test!"
					get := e.Message
					_assert.Equal(t, expect, get)
				}
				{
					var expect map[string]interface{} = map[string]interface{}{"data": "something ..."}
					get := e.Data.(map[string]interface{})
					_assert.Equal(t, expect, get)
				}
			}
		}()
		Insure(false).
			Code(100).
			Message("It's just a test!").
			Data(map[string]interface{}{"data": "something ..."}).
			Do()
	})()
}

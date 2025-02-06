package _interceptor

import (
	"fmt"
	"github.com/junyang7/go-common/_codeMessage"
	"github.com/junyang7/go-common/_exception"
)

type interceptor struct {
	ok      bool
	code    int
	message string
	data    interface{}
}

func Insure(ok bool) *interceptor {
	return &interceptor{
		ok:      ok,
		code:    _codeMessage.ErrDefault.Code,
		message: _codeMessage.ErrDefault.Message,
		data:    struct{}{},
	}
}
func (this *interceptor) Code(code int) *interceptor {
	this.code = code
	return this
}
func (this *interceptor) Message(message interface{}) *interceptor {
	this.message = fmt.Sprintf("%v", message)
	return this
}
func (this *interceptor) Data(data interface{}) *interceptor {
	this.data = data
	return this
}
func (this *interceptor) CodeMessage(codeMessage *_codeMessage.CodeMessage) *interceptor {
	this.code = codeMessage.Code
	this.message = codeMessage.Message
	return this
}
func (this *interceptor) Do() {
	if this.ok {
		return
	}
	exception := _exception.New()
	exception.Code = this.code
	exception.Message = this.message
	exception.Data = this.data
	exception.Throw()
}

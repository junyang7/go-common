package _interceptor

import (
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_exception"
	"runtime"
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
		code:    -1,
		message: "failure",
		data:    struct{}{},
	}
}
func (this *interceptor) Code(code int) *interceptor {
	this.code = code
	return this
}
func (this *interceptor) Message(message string) *interceptor {
	this.message = message
	return this
}
func (this *interceptor) Data(data interface{}) *interceptor {
	this.data = data
	return this
}
func (this *interceptor) Exception(codeMessage *_codeMessage.CodeMessage) *interceptor {
	this.code = codeMessage.Code
	this.message = codeMessage.Message
	return this
}
func (this *interceptor) Do() {
	if this.ok {
		return
	}
	response := _exception.New()
	response.Code = this.code
	response.Message = this.message
	response.Data = this.data
	if _, file, line, ok := runtime.Caller(1); ok {
		response.File = file
		response.Line = line
	}
	panic(response)
}

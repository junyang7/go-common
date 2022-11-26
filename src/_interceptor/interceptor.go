package _interceptor

import (
	"github.com/junyang7/go-common/src/_exception"
	"github.com/junyang7/go-common/src/_response"
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
func (this *interceptor) Exception(exception _exception.Exception) *interceptor {
	this.code = exception.Code
	this.message = exception.Message
	return this
}
func (this *interceptor) Do() {
	if this.ok {
		return
	}
	response := _response.New()
	response.Code = this.code
	response.Message = this.message
	response.Data = this.data
	if _, file, line, ok := runtime.Caller(1); ok {
		response.File = file
		response.Line = line
	}
	panic(response)
}

package _exception

import "github.com/junyang7/go-common/_codeMessage"

type Exception struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New() *Exception {
	return &Exception{
		Code:    _codeMessage.ErrDefault.Code,
		Message: _codeMessage.ErrDefault.Message,
		Data:    struct{}{},
	}
}
func (this *Exception) Throw() {
	panic(this)
}

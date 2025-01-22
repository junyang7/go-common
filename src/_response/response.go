package _response

import "github.com/junyang7/go-common/src/_codeMessage"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	Consume int64       `json:"consume"`
	File    string      `json:"file,omitempty"`
	Line    int         `json:"line,omitempty"`
}

func New() *Response {
	return &Response{
		Code:    _codeMessage.ErrNone.Code,
		Message: _codeMessage.ErrNone.Message,
		Data:    struct{}{},
	}
}

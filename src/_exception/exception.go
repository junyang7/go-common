package _exception

import "github.com/junyang7/go-common/src/_codeMessage"

type Exception struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	Consume int64       `json:"consume"`
	Trace   string      `json:"trace"`
	File    string      `json:"file"`
	Line    int         `json:"line"`
}

func New() *Exception {
	return &Exception{
		Code:    _codeMessage.ErrDefault.Code,
		Message: _codeMessage.ErrDefault.Message,
	}
}

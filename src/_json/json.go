package _json

import (
	"encoding/json"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_parameter"
)

func Encode(data interface{}) []byte {
	b, err := json.Marshal(data)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return b
}
func EncodeAsString(data interface{}) string {
	return _as.String(Encode(data))
}
func Decode(source []byte, target interface{}) {
	if err := json.Unmarshal(source, target); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

type reader struct {
	path string `json:"path"`
	text string `json:"text"`
}

func New() *reader {
	return &reader{}
}
func (this *reader) Path(path string) *reader {
	this.path = path
	return this
}
func (this *reader) Text(text string) *reader {
	this.text = text
	return this
}
func (this *reader) Get(path string) *_parameter.Parameter {
	return _parameter.New("", "")
}

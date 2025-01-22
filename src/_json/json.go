package _json

import (
	"encoding/json"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_interceptor"
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

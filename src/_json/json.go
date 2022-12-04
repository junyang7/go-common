package _json

import (
	"encoding/json"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
)

func Encode(data interface{}) []byte {
	b, err := json.Marshal(data)
	_interceptor.Insure(nil != err).
		CodeMessage(_codeMessage.ErrJsonMarshal).
		Message(err.Error()).
		Do()
	return b
}
func EncodeAsString(data interface{}) string {
	return _as.String(Encode(data))
}
func Decode(source []byte, target interface{}) {
	err := json.Unmarshal(source, target)
	_interceptor.Insure(nil != err).
		CodeMessage(_codeMessage.ErrJsonUnmarshal).
		Message(err.Error()).
		Do()
}

package _xml

import (
	"encoding/xml"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_interceptor"
)

func Encode(data interface{}) []byte {
	b, err := xml.Marshal(data)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return b
}
func EncodeAsString(data interface{}) string {
	return _as.String(Encode(data))
}
func Decode(source []byte, target interface{}) {
	if err := xml.Unmarshal(source, target); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

package _xml

import (
	"encoding/xml"
	"github.com/junyang7/go-common/src/_interceptor"
)

func Decode(source []byte, target interface{}) {
	if err := xml.Unmarshal(source, target); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

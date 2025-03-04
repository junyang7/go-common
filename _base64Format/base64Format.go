package _base64Format

import (
	"encoding/base64"
	"github.com/junyang7/go-common/_interceptor"
)

func Encode(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}
func Decode(data string) string {
	b, err := base64.RawURLEncoding.DecodeString(data)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return string(b)
}

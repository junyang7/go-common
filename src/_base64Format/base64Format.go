package _base64Format

import (
	"encoding/base64"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
)

func Encode(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}
func Decode(data string) string {
	b, err := base64.RawURLEncoding.DecodeString(data)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrBase64RawURLEncodingDecodeString).
		Message(err).
		Do()
	return _as.String(b)
}

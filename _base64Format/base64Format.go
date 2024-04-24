package _base64Format

import (
	"encoding/base64"
	"git.ziji.fun/junyang/go-common/_as"
)

func Encode(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}
func Decode(data string) string {
	b, err := base64.RawURLEncoding.DecodeString(data)
	if nil != err {
		panic(err)
	}
	return _as.String(b)
}

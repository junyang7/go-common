package _json

import (
	"encoding/json"
	"git.ziji.fun/junyang/go-common/_as"
)

func Encode(data interface{}) []byte {
	b, err := json.Marshal(data)
	if nil != err {
		panic(err)
	}
	return b
}
func EncodeAsString(data interface{}) string {
	return _as.String(Encode(data))
}
func Decode(source []byte, target interface{}) {
	if err := json.Unmarshal(source, target); nil != err {
		panic(err)
	}
}

package _json

import (
	"encoding/json"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_file"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_parameter"
	"strings"
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
	json interface{}
}

func New() *reader {
	return &reader{}
}
func (this *reader) Byte(byte []byte) *reader {
	Decode(byte, &this.json)
	return this
}
func (this *reader) Text(text string) *reader {
	return this.Byte([]byte(text))
}
func (this *reader) Path(path string) *reader {
	return this.Byte(_file.ReadAll(path))
}
func (this *reader) Get(path string) *_parameter.Parameter {
	o := this.json
	pathList := strings.Split(path, ".")
	for _, path := range pathList {
		switch v := o.(type) {
		case map[string]interface{}:
			if v, exists := v[path]; exists {
				o = v
			} else {
				o = nil
				break
			}
		case []interface{}:
			index := _as.Int(path)
			if index >= 0 && index < len(v) {
				o = v[index]
			} else {
				o = nil
				break
			}
		default:
			o = nil
			break
		}
	}
	return _parameter.New(path, o)
}

package _json

import (
	"encoding/json"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_parameter"
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
	v interface{}
}

func New() _conf.Conf {
	return &reader{}
}
func (this *reader) Byte(byte []byte) _conf.Conf {
	Decode(byte, &this.v)
	return this
}
func (this *reader) Text(text string) _conf.Conf {
	return this.Byte([]byte(text))
}
func (this *reader) File(path string) _conf.Conf {
	return this.Byte(_file.ReadAll(path))
}
func (this *reader) Get(path string) *_parameter.Parameter {
	o := this.v
	pathList := strings.Split(path, ".")
	for _, path := range pathList {
		switch v := o.(type) {
		case map[string]interface{}:
			if v, exists := v[path]; exists {
				o = v
			} else {
				o = nil
			}
			break
		case []interface{}:
			index := _as.Int(path)
			if index >= 0 && index < len(v) {
				o = v[index]
			} else {
				o = nil
			}
			break
		default:
			o = nil
			break
		}
	}
	return _parameter.New(path, o)
}

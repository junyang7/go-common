package _us

import (
	"github.com/junyang7/go-common/src/_aes"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_hash"
	"github.com/junyang7/go-common/src/_json"
	"github.com/junyang7/go-common/src/_unixMilli"
	"math/rand"
)

type Conf struct {
	K32      string `json:"k32"`
	I16      string `json:"i16"`
	Expires  int64  `json:"expires"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	Secure   string `json:"secure"`
	HttpOnly string `json:"http_only"`
}

func Encode(data map[string]string, conf *Conf) string {
	_t := _unixMilli.Get()
	data["_r"] = _as.String(rand.Int())
	data["_t"] = _as.String(_t)
	if _, ok := data["_e"]; !ok {
		data["_e"] = _as.String(_t + conf.Expires)
	}
	u := _aes.Encode(_json.EncodeAsString(data), conf.K32, conf.I16)
	return u + _hash.Md5(u)
}
func Decode(u string, conf *Conf) map[string]string {
	var res map[string]string
	_json.Decode([]byte(_aes.Decode(u, conf.K32, conf.I16)), &res)
	return res
}

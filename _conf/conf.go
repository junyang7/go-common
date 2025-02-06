package _conf

import (
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_parameter"
	"github.com/junyang7/go-common/_toml"
	"sync"
)

type Conf interface {
	Byte(byte []byte) Conf
	Text(text string) Conf
	File(path string) Conf
	Get(path string) *_parameter.Parameter
}

var (
	conf  Conf
	mutex sync.Mutex
)

func Load(path string, format string) {
	mutex.Lock()
	defer mutex.Unlock()
	switch format {
	case "toml":
		conf = _toml.New().File(path)
	case "json":
		conf = _json.New().File(path)
	}
}
func Reload(path string, format string) {
	Load(path, format)
}
func Get(path string) *_parameter.Parameter {
	mutex.Lock()
	defer mutex.Unlock()
	return conf.Get(path)
}

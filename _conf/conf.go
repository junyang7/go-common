package _conf

import (
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_parameter"
	"sync"
)

type Conf interface {
	Byte(byte []byte) Conf
	Text(text string) Conf
	File(path string) Conf
	Get(path string) *_parameter.Parameter
}

var (
	this  Conf
	mutex sync.Mutex
)

func Load(conf Conf) {
	mutex.Lock()
	defer mutex.Unlock()
	this = conf
}
func Get(path string) *_parameter.Parameter {
	mutex.Lock()
	defer mutex.Unlock()
	if this == nil {
		_interceptor.Insure(false).Message("config not loaded, call _conf.Load() first").Do()
	}
	return this.Get(path)
}
func Reset() {
	mutex.Lock()
	defer mutex.Unlock()
	this = nil
}
func IsLoaded() bool {
	mutex.Lock()
	defer mutex.Unlock()
	return this != nil
}

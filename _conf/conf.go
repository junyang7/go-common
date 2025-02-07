package _conf

import (
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
	return this.Get(path)
}

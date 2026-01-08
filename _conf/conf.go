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

// Load 加载配置实例
func Load(conf Conf) {
	mutex.Lock()
	defer mutex.Unlock()
	this = conf
}

// Get 获取配置值（委托给已加载的配置实现）
func Get(path string) *_parameter.Parameter {
	mutex.Lock()
	defer mutex.Unlock()
	if this == nil {
		_interceptor.Insure(false).Message("config not loaded, call _conf.Load() first").Do()
	}
	return this.Get(path) // 调用具体实现（_json 或 _toml）的 Get 方法
}

// Reset 重置配置（清空）
func Reset() {
	mutex.Lock()
	defer mutex.Unlock()
	this = nil
}

// IsLoaded 检查配置是否已加载
func IsLoaded() bool {
	mutex.Lock()
	defer mutex.Unlock()
	return this != nil
}

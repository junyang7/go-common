package _server

import (
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_redis"
	"github.com/junyang7/go-common/_sql"
	"sync"
	"time"
)

// 优雅关闭超时时间
const shutdownTimeout = 30 * time.Second

// 全局配置加载锁（确保只加载一次）
var (
	loadOnce sync.Once
	loaded   bool
)

// load 加载全局配置（只执行一次）
func load(conf _conf.Conf) {
	loadOnce.Do(func() {
		_conf.Load(conf)
		_sql.Load()
		_redis.Load()
		loaded = true
	})
}

// IsLoaded 检查配置是否已加载
func IsLoaded() bool {
	return loaded
}

// ResetLoadState 重置加载状态（仅用于测试）
func ResetLoadState() {
	loadOnce = sync.Once{}
	loaded = false
}


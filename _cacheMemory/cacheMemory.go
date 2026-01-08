package _cacheMemory

import (
	"sync"
	"time"
)

type cache struct {
	v interface{}
	t time.Time
}

var cacheDict sync.Map

// 初始化时启动自动清理协程
func init() {
	go cleanExpired()
}

// 后台协程：定期清理过期缓存
func cleanExpired() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		now := time.Now()
		cacheDict.Range(func(key, value interface{}) bool {
			c := value.(*cache)
			if now.After(c.t) {
				cacheDict.Delete(key)
			}
			return true
		})
	}
}

// Set 设置缓存
func Set(k string, v interface{}, ttl time.Duration) {
	cacheDict.Store(k, &cache{
		v: v,
		t: time.Now().Add(ttl),
	})
}

// Get 获取缓存，如果不存在或已过期返回 nil
func Get(k string) interface{} {
	if v, ok := cacheDict.Load(k); ok {
		c := v.(*cache)
		if time.Now().After(c.t) {
			cacheDict.Delete(k)
			return nil
		}
		return c.v
	}
	return nil
}

// Del 删除缓存
func Del(k string) {
	cacheDict.Delete(k)
}

// Exists 检查缓存是否存在且未过期
func Exists(k string) bool {
	if v, ok := cacheDict.Load(k); ok {
		c := v.(*cache)
		if time.Now().After(c.t) {
			cacheDict.Delete(k)
			return false
		}
		return true
	}
	return false
}

// Clear 清空所有缓存
func Clear() {
	cacheDict.Range(func(key, value interface{}) bool {
		cacheDict.Delete(key)
		return true
	})
}

// GetOrSet 获取缓存，如果不存在则调用 fn 生成并缓存
func GetOrSet(k string, fn func() interface{}, ttl time.Duration) interface{} {
	// 先尝试获取
	if v := Get(k); v != nil {
		return v
	}
	
	// 不存在，生成新值
	v := fn()
	Set(k, v, ttl)
	return v
}

// Count 返回当前缓存数量（包含过期的）
func Count() int {
	count := 0
	cacheDict.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

// Keys 返回所有缓存的 key（包含过期的）
func Keys() []string {
	keys := []string{}
	cacheDict.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}


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

func init() {
	go cleanExpired()
}
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
func Set(k string, v interface{}, ttl time.Duration) {
	cacheDict.Store(k, &cache{
		v: v,
		t: time.Now().Add(ttl),
	})
}
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
func Del(k string) {
	cacheDict.Delete(k)
}
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
func Clear() {
	cacheDict.Range(func(key, value interface{}) bool {
		cacheDict.Delete(key)
		return true
	})
}
func GetOrSet(k string, fn func() interface{}, ttl time.Duration) interface{} {
	if v := Get(k); v != nil {
		return v
	}
	v := fn()
	Set(k, v, ttl)
	return v
}
func Count() int {
	count := 0
	cacheDict.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}
func GetAll() []string {
	keys := []string{}
	cacheDict.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}

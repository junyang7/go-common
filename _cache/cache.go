package _cache

import (
	"sync"
	"time"
)

type cache struct {
	v interface{}
	t time.Time
}

var cacheDict sync.Map
var m sync.RWMutex

func SetByMemory(k string, v interface{}, ttl time.Duration) bool {
	cacheDict.Store(k, &cache{
		v: v,
		t: time.Now().Add(ttl),
	})
	return true
}
func GetByMemory(k string) interface{} {
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
func DelByMemory(k string) {
	cacheDict.Delete(k)
}

package _cache

import (
	"sync"
	"time"
)

type cache struct {
	v interface{}
	t time.Duration
}

var cacheDict map[string]*cache = map[string]*cache{}
var m sync.RWMutex = sync.RWMutex{}

func SetByMemory(k string, v interface{}, t time.Duration) bool {
	if m.TryLock() {
		defer m.Unlock()
		cacheDict[k] = &cache{
			v: v,
			t: t,
		}
		return true
	}
	return false
}
func GetByMemory(k string) interface{} {
	if m.TryRLock() {
		defer m.RUnlock()
		v, ok := cacheDict[k]
		if ok {
			return v
		}
	}
	return nil
}

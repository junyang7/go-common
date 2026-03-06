package _lock

import "sync"

var lock sync.Map

func Get(key string) *sync.Mutex {
	v, _ := lock.LoadOrStore(key, &sync.Mutex{})
	return v.(*sync.Mutex)
}

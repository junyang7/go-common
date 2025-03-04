package _cache

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
	"time"
)

func TestSetByMemory(t *testing.T) {

	{
		var key string = "testKey"
		var ttl time.Duration = time.Second * 5
		var expect bool = true
		var give interface{} = "testValue"
		get := SetByMemory(key, give, ttl)
		_assert.Equal(t, expect, get)
	}

}
func TestGetByMemory(t *testing.T) {

	{
		var key string = "testKey"
		var ttl time.Duration = time.Second * 5
		var give interface{} = "testValue"
		SetByMemory(key, give, ttl)
		var expect interface{} = give
		get := GetByMemory(key)
		_assert.Equal(t, expect, get)
	}

	{
		var key string = "expiredKey"
		var ttl time.Duration = time.Millisecond * 500
		var give interface{} = "expiredValue"
		SetByMemory(key, give, ttl)
		time.Sleep(time.Second)
		var expect interface{} = nil
		get := GetByMemory(key)
		_assert.Equal(t, expect, get)
	}

}
func TestDelByMemory(t *testing.T) {

	{
		var key string = "testKey"
		var ttl time.Duration = time.Second * 5
		var give interface{} = "testValue"
		SetByMemory(key, give, ttl)
		var expect interface{} = give
		get := GetByMemory(key)
		_assert.Equal(t, expect, get)
		DelByMemory(key)
		expect = nil
		get = GetByMemory(key)
		_assert.Equal(t, expect, get)
	}

	{
		var key string = "nonExistentKey"
		DelByMemory(key)
		var expect interface{} = nil
		get := GetByMemory(key)
		_assert.Equal(t, expect, get)
	}

}

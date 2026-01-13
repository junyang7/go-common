package _cacheMemory

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"sync"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	{
		var key string = "testKey"
		var ttl time.Duration = time.Second * 5
		var give interface{} = "testValue"
		Set(key, give, ttl)
		get := Get(key)
		_assert.Equal(t, give, get)
	}
}
func TestGet(t *testing.T) {
	{
		var key string = "testKey"
		var ttl time.Duration = time.Second * 5
		var give interface{} = "testValue"
		Set(key, give, ttl)
		var expect interface{} = give
		get := Get(key)
		_assert.Equal(t, expect, get)
	}
	{
		var key string = "expiredKey"
		var ttl time.Duration = time.Millisecond * 500
		var give interface{} = "expiredValue"
		Set(key, give, ttl)
		time.Sleep(time.Second)
		var expect interface{} = nil
		get := Get(key)
		_assert.Equal(t, expect, get)
	}
	{
		var key string = "nonExistentKey"
		var expect interface{} = nil
		get := Get(key)
		_assert.Equal(t, expect, get)
	}
}
func TestDel(t *testing.T) {
	{
		var key string = "testKey"
		var ttl time.Duration = time.Second * 5
		var give interface{} = "testValue"
		Set(key, give, ttl)
		var expect interface{} = give
		get := Get(key)
		_assert.Equal(t, expect, get)
		Del(key)
		expect = nil
		get = Get(key)
		_assert.Equal(t, expect, get)
	}
	{
		var key string = "nonExistentKey"
		Del(key)
		var expect interface{} = nil
		get := Get(key)
		_assert.Equal(t, expect, get)
	}
}
func TestExists(t *testing.T) {
	{
		key := "existKey"
		Set(key, "value", time.Second*5)
		_assert.True(t, Exists(key))
	}
	{
		key := "nonExistKey"
		_assert.False(t, Exists(key))
	}
	{
		key := "expiredKey"
		Set(key, "value", time.Millisecond*100)
		time.Sleep(time.Millisecond * 200)
		_assert.False(t, Exists(key))
	}
}
func TestClear(t *testing.T) {
	Set("key1", "value1", time.Minute)
	Set("key2", "value2", time.Minute)
	Set("key3", "value3", time.Minute)
	_assert.True(t, Exists("key1"))
	_assert.True(t, Exists("key2"))
	_assert.True(t, Exists("key3"))
	Clear()
	_assert.False(t, Exists("key1"))
	_assert.False(t, Exists("key2"))
	_assert.False(t, Exists("key3"))
}
func TestGetOrSet(t *testing.T) {
	{
		key := "getOrSetKey"
		Del(key) // 确保不存在
		called := false
		result := GetOrSet(key, func() interface{} {
			called = true
			return "generated value"
		}, time.Minute)
		_assert.True(t, called)
		_assert.Equal(t, "generated value", result)
	}
	{
		key := "getOrSetKey"
		called := false
		result := GetOrSet(key, func() interface{} {
			called = true
			return "new value"
		}, time.Minute)
		_assert.False(t, called)
		_assert.Equal(t, "generated value", result)
	}
}
func TestCount(t *testing.T) {
	Clear()
	_assert.Equal(t, 0, Count())
	Set("key1", "value1", time.Minute)
	_assert.Equal(t, 1, Count())
	Set("key2", "value2", time.Minute)
	_assert.Equal(t, 2, Count())
	Set("key3", "value3", time.Minute)
	_assert.Equal(t, 3, Count())
	Del("key2")
	_assert.Equal(t, 2, Count())
	Clear()
	_assert.Equal(t, 0, Count())
}
func TestKeys(t *testing.T) {
	Clear()
	Set("key1", "value1", time.Minute)
	Set("key2", "value2", time.Minute)
	Set("key3", "value3", time.Minute)
	keys := GetAll()
	_assert.Len(t, keys, 3)
	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}
	_assert.True(t, keyMap["key1"])
	_assert.True(t, keyMap["key2"])
	_assert.True(t, keyMap["key3"])
}
func TestConcurrentAccess(t *testing.T) {
	Clear()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key_%d_%d", index, j)
				Set(key, j, time.Minute)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key_%d_%d", index, j)
				Get(key)
			}
		}(i)
	}
	wg.Wait()
	_assert.True(t, true)
}
func TestOverwrite(t *testing.T) {
	key := "overwriteKey"
	Set(key, "value1", time.Minute)
	_assert.Equal(t, "value1", Get(key))
	Set(key, "value2", time.Minute)
	_assert.Equal(t, "value2", Get(key))
	Set(key, "value3", time.Minute)
	_assert.Equal(t, "value3", Get(key))
}
func TestDifferentTypes(t *testing.T) {
	Clear()
	Set("string", "hello", time.Minute)
	_assert.Equal(t, "hello", Get("string"))
	Set("int", 123, time.Minute)
	_assert.Equal(t, 123, Get("int"))
	Set("bool", true, time.Minute)
	_assert.Equal(t, true, Get("bool"))
	type User struct {
		Name string
		Age  int
	}
	user := User{Name: "Alice", Age: 30}
	Set("struct", user, time.Minute)
	_assert.Equal(t, user, Get("struct"))
	slice := []int{1, 2, 3}
	Set("slice", slice, time.Minute)
	result := Get("slice").([]int)
	_assert.EqualByList(t, slice, result)
}

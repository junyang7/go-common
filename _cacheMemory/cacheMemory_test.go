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
		
		// 验证设置成功（通过 Get 验证）
		get := Get(key)
		_assert.Equal(t, give, get)
	}
}

func TestGet(t *testing.T) {
	// 获取存在的缓存
	{
		var key string = "testKey"
		var ttl time.Duration = time.Second * 5
		var give interface{} = "testValue"
		Set(key, give, ttl)
		var expect interface{} = give
		get := Get(key)
		_assert.Equal(t, expect, get)
	}

	// 获取过期的缓存
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
	
	// 获取不存在的缓存
	{
		var key string = "nonExistentKey"
		var expect interface{} = nil
		get := Get(key)
		_assert.Equal(t, expect, get)
	}
}

func TestDel(t *testing.T) {
	// 删除存在的缓存
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

	// 删除不存在的缓存（不应报错）
	{
		var key string = "nonExistentKey"
		Del(key)
		var expect interface{} = nil
		get := Get(key)
		_assert.Equal(t, expect, get)
	}
}

func TestExists(t *testing.T) {
	// 存在且未过期
	{
		key := "existKey"
		Set(key, "value", time.Second*5)
		_assert.True(t, Exists(key))
	}
	
	// 不存在
	{
		key := "nonExistKey"
		_assert.False(t, Exists(key))
	}
	
	// 已过期
	{
		key := "expiredKey"
		Set(key, "value", time.Millisecond*100)
		time.Sleep(time.Millisecond * 200)
		_assert.False(t, Exists(key))
	}
}

func TestClear(t *testing.T) {
	// 设置多个缓存
	Set("key1", "value1", time.Minute)
	Set("key2", "value2", time.Minute)
	Set("key3", "value3", time.Minute)
	
	// 验证都存在
	_assert.True(t, Exists("key1"))
	_assert.True(t, Exists("key2"))
	_assert.True(t, Exists("key3"))
	
	// 清空
	Clear()
	
	// 验证都不存在了
	_assert.False(t, Exists("key1"))
	_assert.False(t, Exists("key2"))
	_assert.False(t, Exists("key3"))
}

func TestGetOrSet(t *testing.T) {
	// 第一次调用，缓存不存在，应该调用 fn
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
	
	// 第二次调用，缓存存在，不应该调用 fn
	{
		key := "getOrSetKey"
		
		called := false
		result := GetOrSet(key, func() interface{} {
			called = true
			return "new value"
		}, time.Minute)
		
		_assert.False(t, called)
		_assert.Equal(t, "generated value", result) // 应该返回之前的值
	}
}

func TestCount(t *testing.T) {
	// 清空
	Clear()
	_assert.Equal(t, 0, Count())
	
	// 添加缓存
	Set("key1", "value1", time.Minute)
	_assert.Equal(t, 1, Count())
	
	Set("key2", "value2", time.Minute)
	_assert.Equal(t, 2, Count())
	
	Set("key3", "value3", time.Minute)
	_assert.Equal(t, 3, Count())
	
	// 删除一个
	Del("key2")
	_assert.Equal(t, 2, Count())
	
	// 清空
	Clear()
	_assert.Equal(t, 0, Count())
}

func TestKeys(t *testing.T) {
	// 清空
	Clear()
	
	// 设置缓存
	Set("key1", "value1", time.Minute)
	Set("key2", "value2", time.Minute)
	Set("key3", "value3", time.Minute)
	
	keys := Keys()
	_assert.Len(t, keys, 3)
	
	// 验证包含所有key（顺序可能不同）
	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}
	_assert.True(t, keyMap["key1"])
	_assert.True(t, keyMap["key2"])
	_assert.True(t, keyMap["key3"])
}

// 测试并发安全性
func TestConcurrentAccess(t *testing.T) {
	Clear()
	
	var wg sync.WaitGroup
	
	// 10个goroutine并发写入
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
	
	// 10个goroutine并发读取
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
	
	// 验证没有panic，测试通过即说明并发安全
	_assert.True(t, true)
}

// 测试覆盖更新（同一key多次Set）
func TestOverwrite(t *testing.T) {
	key := "overwriteKey"
	
	// 第一次设置
	Set(key, "value1", time.Minute)
	_assert.Equal(t, "value1", Get(key))
	
	// 覆盖
	Set(key, "value2", time.Minute)
	_assert.Equal(t, "value2", Get(key))
	
	// 再次覆盖
	Set(key, "value3", time.Minute)
	_assert.Equal(t, "value3", Get(key))
}

// 测试不同类型的值
func TestDifferentTypes(t *testing.T) {
	Clear()
	
	// 字符串
	Set("string", "hello", time.Minute)
	_assert.Equal(t, "hello", Get("string"))
	
	// 整数
	Set("int", 123, time.Minute)
	_assert.Equal(t, 123, Get("int"))
	
	// 布尔值
	Set("bool", true, time.Minute)
	_assert.Equal(t, true, Get("bool"))
	
	// 结构体
	type User struct {
		Name string
		Age  int
	}
	user := User{Name: "Alice", Age: 30}
	Set("struct", user, time.Minute)
	_assert.Equal(t, user, Get("struct"))
	
	// 切片
	slice := []int{1, 2, 3}
	Set("slice", slice, time.Minute)
	result := Get("slice").([]int)
	_assert.EqualByList(t, slice, result)
}


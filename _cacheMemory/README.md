# _cacheMemory - 内存缓存工具

提供基于内存的键值缓存功能，支持 TTL（过期时间）和自动清理。

---

## 核心特性

- ✅ **TTL 支持**：每个缓存项可设置独立的过期时间
- ✅ **自动清理**：后台协程定期清理过期缓存，防止内存泄漏
- ✅ **并发安全**：基于 `sync.Map` 实现，支持高并发读写
- ✅ **类型通用**：支持任意类型（interface{}）
- ✅ **简单易用**：API 简洁直观

---

## 函数列表

### 基础操作

| 函数 | 说明 |
|------|------|
| `Set(k, v, ttl)` | 设置缓存 |
| `Get(k) interface{}` | 获取缓存 |
| `Del(k)` | 删除缓存 |
| `Exists(k) bool` | 检查缓存是否存在 |

### 批量操作

| 函数 | 说明 |
|------|------|
| `Clear()` | 清空所有缓存 |
| `Count() int` | 获取缓存数量 |
| `Keys() []string` | 获取所有 key |

### 高级操作

| 函数 | 说明 |
|------|------|
| `GetOrSet(k, fn, ttl) interface{}` | 获取缓存，不存在则生成并缓存 |

---

## 使用示例

### 基础用法

```go
import (
    "_cacheMemory"
    "time"
)

func main() {
    // 设置缓存（5秒后过期）
    _cacheMemory.Set("user:123", "Alice", time.Second*5)
    
    // 获取缓存
    value := _cacheMemory.Get("user:123")
    if value != nil {
        username := value.(string)
        println(username) // "Alice"
    }
    
    // 删除缓存
    _cacheMemory.Del("user:123")
}
```

### 检查缓存是否存在

```go
if _cacheMemory.Exists("user:123") {
    // 缓存存在且未过期
    value := _cacheMemory.Get("user:123")
    // ...
} else {
    // 缓存不存在或已过期
    // 重新获取数据
}
```

### GetOrSet 模式（推荐）

```go
// 获取缓存，不存在则查询数据库并缓存
user := _cacheMemory.GetOrSet("user:123", func() interface{} {
    // 这个函数只在缓存不存在时才会调用
    return db.QueryUser(123)
}, time.Minute*10)

// 类型断言
userData := user.(*User)
```

### 批量操作

```go
// 获取所有缓存的 key
keys := _cacheMemory.Keys()
println("缓存数量:", len(keys))

// 获取缓存数量
count := _cacheMemory.Count()
println("当前有", count, "个缓存项")

// 清空所有缓存
_cacheMemory.Clear()
```

### 不同类型的值

```go
// 字符串
_cacheMemory.Set("name", "Alice", time.Minute)

// 整数
_cacheMemory.Set("age", 30, time.Minute)

// 布尔值
_cacheMemory.Set("active", true, time.Minute)

// 结构体
type User struct {
    Name string
    Age  int
}
user := User{Name: "Bob", Age: 25}
_cacheMemory.Set("user:456", user, time.Minute)

// 切片
data := []int{1, 2, 3}
_cacheMemory.Set("numbers", data, time.Minute)

// 获取时需要类型断言
username := _cacheMemory.Get("name").(string)
age := _cacheMemory.Get("age").(int)
userData := _cacheMemory.Get("user:456").(User)
```

---

## 典型使用场景

### 1. API 响应缓存

```go
func GetUserInfo(userID int) *User {
    key := fmt.Sprintf("user:%d", userID)
    
    // 使用 GetOrSet 模式
    result := _cacheMemory.GetOrSet(key, func() interface{} {
        // 只有缓存不存在时才会查询数据库
        return db.QueryUser(userID)
    }, time.Minute*5)
    
    return result.(*User)
}
```

### 2. 配置缓存

```go
func GetConfig(key string) string {
    cacheKey := "config:" + key
    
    if value := _cacheMemory.Get(cacheKey); value != nil {
        return value.(string)
    }
    
    // 从配置文件读取
    config := readConfigFromFile(key)
    _cacheMemory.Set(cacheKey, config, time.Hour)
    
    return config
}
```

### 3. 限流计数

```go
func CheckRateLimit(userID string) bool {
    key := "rate:" + userID
    
    count := 0
    if v := _cacheMemory.Get(key); v != nil {
        count = v.(int)
    }
    
    if count >= 100 {
        return false // 超过限制
    }
    
    _cacheMemory.Set(key, count+1, time.Minute)
    return true
}
```

### 4. 验证码缓存

```go
// 设置验证码（5分钟有效）
func SetVerifyCode(phone, code string) {
    key := "verify:" + phone
    _cacheMemory.Set(key, code, time.Minute*5)
}

// 验证验证码
func CheckVerifyCode(phone, code string) bool {
    key := "verify:" + phone
    
    if !_cacheMemory.Exists(key) {
        return false // 验证码不存在或已过期
    }
    
    cachedCode := _cacheMemory.Get(key).(string)
    if cachedCode == code {
        _cacheMemory.Del(key) // 验证成功后删除
        return true
    }
    
    return false
}
```

---

## 技术细节

### 实现原理

```go
// 底层存储
var cacheDict sync.Map  // Go 标准库的并发安全 Map

// 缓存结构
type cache struct {
    v interface{}  // 缓存的值
    t time.Time    // 过期时间
}

// 自动清理
func init() {
    go cleanExpired()  // 启动后台协程
}

func cleanExpired() {
    // 每分钟清理一次过期缓存
    ticker := time.NewTicker(time.Minute)
    for range ticker.C {
        // 遍历所有缓存，删除已过期的
    }
}
```

### 并发安全

- ✅ 使用 `sync.Map` 保证并发安全
- ✅ 无需额外的锁机制
- ✅ 支持高并发读写

### 过期机制

#### 1. 被动删除（Get 时）
```go
Get(k) {
    if 已过期 {
        删除缓存
        return nil
    }
}
```

#### 2. 主动清理（后台协程）
```go
每分钟扫描一次，删除所有过期缓存
→ 防止内存泄漏
```

---

## 注意事项

### 1. 类型断言

由于使用 `interface{}`，获取值时需要类型断言：

```go
// 设置
_cacheMemory.Set("key", "value", time.Minute)

// 获取
value := _cacheMemory.Get("key")
if value != nil {
    str := value.(string)  // 需要类型断言
}
```

### 2. 过期时间

```go
// TTL 从当前时间开始计算
Set("key", "value", time.Second*10)
// 10秒后过期

// 负数或0的TTL会立即过期
Set("key", "value", -time.Second)  // 立即过期
Set("key", "value", 0)             // 立即过期
```

### 3. 内存占用

```go
⚠️ 所有缓存数据存储在内存中
⚠️ 缓存过多可能占用大量内存
⚠️ 进程重启后缓存丢失

建议：
- 定期监控缓存数量（Count）
- 合理设置 TTL
- 避免缓存大对象
```

### 4. 并发竞争（非问题）

```go
// 场景：两个 goroutine 同时 GetOrSet 同一个 key
goroutine1: GetOrSet("key", fn1, ttl)
goroutine2: GetOrSet("key", fn2, ttl)

// 可能结果：
- fn1 和 fn2 可能都被调用
- 最后一个 Set 的值会保留
- 这是预期行为（sync.Map 的特性）
```

---

## 性能特点

### 优点

| 特点 | 说明 |
|------|------|
| **读性能高** | sync.Map 对读操作优化 |
| **写性能好** | 无全局锁，并发写入快 |
| **内存效率** | 自动清理过期缓存 |

### 局限

| 局限 | 说明 |
|------|------|
| **单机内存** | 不支持分布式 |
| **无持久化** | 进程重启丢失 |
| **无 LRU** | 只有 TTL，无淘汰策略 |

---

## 对比其他缓存方案

| 特性 | _cacheMemory | go-cache | Redis | Memcached |
|------|--------------|----------|-------|-----------|
| **部署** | 内置 | 内置 | 独立服务 | 独立服务 |
| **性能** | 极快 | 快 | 网络开销 | 网络开销 |
| **持久化** | ❌ | ❌ | ✅ | ❌ |
| **分布式** | ❌ | ❌ | ✅ | ✅ |
| **TTL** | ✅ | ✅ | ✅ | ✅ |
| **自动清理** | ✅ | ✅ | ✅ | ✅ |
| **复杂度** | ⭐ 简单 | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ |

---

## API 参考

### Set

```go
func Set(k string, v interface{}, ttl time.Duration)
```

设置缓存项。

**参数：**
- `k`: 缓存键
- `v`: 缓存值（任意类型）
- `ttl`: 过期时间（从当前时间开始计算）

### Get

```go
func Get(k string) interface{}
```

获取缓存项，如果不存在或已过期返回 `nil`。

### Del

```go
func Del(k string)
```

删除缓存项。删除不存在的 key 不会报错。

### Exists

```go
func Exists(k string) bool
```

检查缓存项是否存在且未过期。

### Clear

```go
func Clear()
```

清空所有缓存项。

### GetOrSet

```go
func GetOrSet(k string, fn func() interface{}, ttl time.Duration) interface{}
```

获取缓存，如果不存在则调用 `fn` 生成值并缓存。

**参数：**
- `k`: 缓存键
- `fn`: 生成值的函数（仅在缓存不存在时调用）
- `ttl`: 过期时间

**返回：**
- 缓存值或新生成的值

### Count

```go
func Count() int
```

返回当前缓存数量（包含已过期但未清理的）。

### Keys

```go
func Keys() []string
```

返回所有缓存的 key（包含已过期但未清理的）。

---

## 设计哲学

### 简洁性

```go
✅ 仅 8 个核心函数
✅ API 直观，易于理解
✅ 无需复杂配置
```

### 自动化

```go
✅ 自动清理过期缓存
✅ 自动处理并发
✅ 自动管理内存
```

### 性能优先

```go
✅ 基于 sync.Map（高性能并发）
✅ 内存存储（无网络开销）
✅ 后台清理（不影响主流程）
```

---

## 最佳实践

### 1. 合理设置 TTL

```go
✅ 推荐
- 热点数据：1-5 分钟
- 配置数据：10-30 分钟
- 临时数据：10-60 秒

❌ 避免
- 永久缓存（容易忘记清理）
- 过短 TTL（频繁失效，缓存意义不大）
```

### 2. 使用 GetOrSet 模式

```go
// ✅ 推荐（简洁）
result := _cacheMemory.GetOrSet(key, func() interface{} {
    return fetchFromDB()
}, time.Minute*5)

// ⚠️ 不推荐（冗长）
result := _cacheMemory.Get(key)
if result == nil {
    result = fetchFromDB()
    _cacheMemory.Set(key, result, time.Minute*5)
}
```

### 3. 定期监控

```go
// 定期检查缓存数量
go func() {
    ticker := time.NewTicker(time.Minute*10)
    for range ticker.C {
        count := _cacheMemory.Count()
        if count > 10000 {
            log.Warn("缓存数量过多", "count", count)
        }
    }
}()
```

### 4. 键命名规范

```go
// ✅ 推荐：使用前缀和分隔符
_cacheMemory.Set("user:123", ...)
_cacheMemory.Set("config:db:host", ...)
_cacheMemory.Set("session:abc123", ...)

// ❌ 避免：无规则的命名
_cacheMemory.Set("u123", ...)
_cacheMemory.Set("data", ...)
```

---

## 性能指标（参考）

```
操作类型    耗时（纳秒级）
─────────  ──────────
Set         ~100-200ns
Get (命中)   ~50-100ns
Get (未命中) ~30-50ns
Delete      ~50-100ns
Exists      ~50-100ns
```

---

## 限制与注意

### 1. 单机限制
- ⚠️ 只能在单个进程内使用
- ⚠️ 不支持分布式缓存
- 💡 如需分布式，使用 Redis

### 2. 内存限制
- ⚠️ 所有数据存储在内存中
- ⚠️ 缓存过多可能导致 OOM
- 💡 监控 `Count()` 控制数量

### 3. 无持久化
- ⚠️ 进程重启后缓存丢失
- ⚠️ 不适合存储重要数据
- 💡 只用于临时、可重建的数据

### 4. 清理延迟
- ℹ️ 后台每分钟清理一次
- ℹ️ 过期缓存可能存在最多1分钟
- ℹ️ Get 时会立即清理过期项

---

## 总结

**_cacheMemory 是一个：**
- ✅ 简单实用的内存缓存工具
- ✅ 并发安全、性能优秀
- ✅ 自动清理、防止泄漏
- ✅ 适合单机、高性能场景

**适用于：**
- Web 应用的热点数据缓存
- API 响应缓存
- 配置缓存
- 会话数据
- 限流计数
- 临时数据存储

**不适用于：**
- 分布式系统
- 需要持久化的场景
- 大量数据存储
- 需要复杂淘汰策略的场景



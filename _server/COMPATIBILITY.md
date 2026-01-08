# 路由管理器兼容性说明

## ✅ 向后兼容性保证

本次重构**100% 向后兼容**，现有代码无需任何修改。

---

## 🎯 默认行为（推荐）

### 使用场景：生产环境、开发环境、单服务器应用

**路由注册** (`router/api.go`):
```go
package router

import (
    "BE/controller"
    "github.com/junyang7/go-common/_router"
)

func init() {
    // 路由自动注册到全局管理器
    _router.Prefix("/api").Group(func() {
        _router.Get("/login", controller.Login)
        _router.Post("/register", controller.Register)
        // ... 更多路由
    })
}
```

**服务器启动** (`main.go`):
```go
package main

import (
    _ "BE/router"  // 触发 init()，加载路由
    "github.com/junyang7/go-common/_server"
    "github.com/junyang7/go-common/_toml"
)

func main() {
    _server.Http().
        Load(_toml.New().File("./etc/app.toml"), "server.http").
        Run()
    // ✅ 自动使用全局路由，完全兼容旧代码
}
```

**特点**:
- ✅ 零改动，100% 兼容
- ✅ 路由在 `init()` 中全局注册
- ✅ 服务器自动使用全局路由
- ✅ 简单直观，最佳实践

---

## 🔧 自定义路由管理器（可选）

### 使用场景：并发测试、路由隔离

**注意**: 使用自定义管理器后，`init()` 中的全局路由将不可用！

#### 示例 1: 测试场景（路由隔离）

```go
func TestAPIServer_Concurrent_A(t *testing.T) {
    t.Parallel()
    
    // 创建独立的路由管理器
    manager := _router.NewManager()
    builder := _router.NewBuilder(manager)
    
    // 显式注册路由（不能用 init）
    builder.Get("/test-a", func(ctx *_context.Context) {
        ctx.JSON(map[string]string{"test": "a"})
    })
    
    // 使用自定义管理器
    server := _server.Api().
        Port("0").
        RouterManager(manager).  // ← 使用独立路由
        RunWithContext(ctx)
    
    // ✅ 与其他测试完全隔离
}

func TestAPIServer_Concurrent_B(t *testing.T) {
    t.Parallel()
    
    // 另一个独立的路由管理器
    manager := _router.NewManager()
    builder := _router.NewBuilder(manager)
    
    builder.Get("/test-b", func(ctx *_context.Context) {
        ctx.JSON(map[string]string{"test": "b"})
    })
    
    server := _server.Api().
        Port("0").
        RouterManager(manager).  // ← 完全隔离
        RunWithContext(ctx)
}
```

#### 示例 2: 开发环境多服务器（不推荐）

```go
func main() {
    // ⚠️ 不推荐：单进程多服务器
    // 生产环境应该拆分为独立进程
    
    // 管理后台（独立路由）
    adminManager := _router.NewManager()
    adminBuilder := _router.NewBuilder(adminManager)
    adminBuilder.Get("/admin/users", adminHandler)
    
    go _server.Api().
        Port("8080").
        RouterManager(adminManager).
        Run()
    
    // 公开API（独立路由）
    publicManager := _router.NewManager()
    publicBuilder := _router.NewBuilder(publicManager)
    publicBuilder.Get("/public/info", publicHandler)
    
    go _server.Api().
        Port("9090").
        RouterManager(publicManager).
        Run()
    
    select {}
}
```

**更好的做法**:
```bash
# 拆分为独立进程（推荐）
./admin-server --port=8080 &
./public-server --port=9090 &
```

---

## 📋 API 参考

### BaseEngine 方法

#### `RouterManager(manager *_router.Manager) *Engine`

设置自定义路由管理器。

**参数**:
- `manager`: 自定义的路由管理器实例

**返回**:
- 返回引擎实例（支持链式调用）

**适用引擎**:
- `webEngine`
- `apiEngine`
- `httpEngine`
- `rpcEngine`

**注意事项**:
1. ⚠️ 使用自定义管理器后，`init()` 中注册的全局路由将不可用
2. ⚠️ 主要用于测试场景的路由隔离
3. ⚠️ 生产环境建议使用默认的全局路由（更简单）
4. ⚠️ 不能在 `init()` 中使用，必须显式注册路由

**示例**:
```go
// 创建自定义管理器
manager := _router.NewManager()
builder := _router.NewBuilder(manager)

// 显式注册路由
builder.Get("/custom", handler)

// 使用自定义管理器
server := _server.Api().
    RouterManager(manager).
    Run()
```

---

## 🔍 常见问题

### Q1: 我需要使用 `RouterManager()` 吗？

**A**: 99% 的情况下**不需要**。

- ✅ 生产环境：使用默认全局路由
- ✅ 开发环境：使用默认全局路由
- ⚠️ 测试场景：如需并发测试隔离，才使用自定义管理器

### Q2: 如何在 `init()` 中注册路由？

**A**: 直接使用 `_router` 包的函数即可，自动注册到全局。

```go
func init() {
    _router.Get("/test", handler)  // ✅ 自动注册到全局
}
```

### Q3: 使用 `RouterManager()` 后，为什么找不到路由？

**A**: 因为 `init()` 中的路由注册到全局管理器，而你使用了自定义管理器。

**解决方案**:
```go
// ❌ 错误用法
func init() {
    _router.Get("/test", handler)  // 注册到全局
}

func main() {
    manager := _router.NewManager()  // 新建管理器
    _server.Api().
        RouterManager(manager).  // ← 使用新管理器，找不到 /test
        Run()
}

// ✅ 正确用法 1：使用默认全局路由（推荐）
func init() {
    _router.Get("/test", handler)
}

func main() {
    _server.Api().Run()  // ← 自动使用全局路由
}

// ✅ 正确用法 2：使用自定义管理器
func main() {
    manager := _router.NewManager()
    builder := _router.NewBuilder(manager)
    builder.Get("/test", handler)  // ← 显式注册
    
    _server.Api().
        RouterManager(manager).
        Run()
}
```

### Q4: 多个服务器需要不同的路由怎么办？

**A**: **拆分为独立进程**（推荐），而不是单进程多服务器。

```bash
# ✅ 推荐：独立进程
./admin-server &
./api-server &

# ⚠️ 不推荐：单进程多服务器
./monolith-server  # 内部启动多个服务器
```

**理由**:
- ✅ 故障隔离：一个崩溃不影响其他
- ✅ 独立扩展：高负载服务多实例
- ✅ 资源隔离：内存/CPU 完全独立
- ✅ 云原生：容器化部署标准

---

## 🎯 最佳实践

### 生产环境

```go
// router/api.go
func init() {
    _router.Prefix("/api").Group(func() {
        _router.Get("/users", controller.Users)
        // ... 所有路由
    })
}

// main.go
func main() {
    ctx, cancel := signal.NotifyContext(
        context.Background(),
        os.Interrupt,
        syscall.SIGTERM,
    )
    defer cancel()
    
    _server.Http().
        Load(conf, "server.http").
        RunWithContext(ctx)  // ✅ 简单、高效、稳定
}
```

### 测试场景

```go
// api_test.go
func TestAPIServer(t *testing.T) {
    // 方案1：使用全局路由 + Reset（简单）
    _router.ResetDefaultManager()
    _router.Get("/test", handler)
    
    server := _server.Api().Port("0")
    // ✅ 每个测试前 Reset 即可
    
    // 方案2：使用自定义管理器（完全隔离）
    manager := _router.NewManager()
    builder := _router.NewBuilder(manager)
    builder.Get("/test", handler)
    
    server := _server.Api().
        Port("0").
        RouterManager(manager)
    // ✅ 支持并发测试
}
```

---

## 🔄 迁移指南

### 从旧版本升级

**好消息**: 无需任何修改！

```go
// 你的旧代码
func init() {
    _router.Get("/test", handler)
}

func main() {
    _server.Api().Run()
}

// ✅ 新版本完全兼容，继续使用即可
```

---

## 📊 设计决策

### 为什么默认使用全局管理器？

1. **向后兼容** - 现有代码无需修改
2. **符合习惯** - `init()` 全局注册是 Go 的标准模式
3. **简单直观** - 不引入不必要的复杂性
4. **满足 99% 场景** - 单服务器应用占绝大多数

### 为什么提供自定义管理器？

1. **测试隔离** - 支持并发测试，避免状态污染
2. **架构优雅** - 面向对象设计，实例级管理
3. **未来扩展** - 为特殊场景预留灵活性

### 为什么不推荐单进程多服务器？

1. **资源隔离差** - 共享内存/CPU，一个崩溃全崩
2. **扩展性差** - 无法独立扩展不同服务
3. **运维复杂** - 监控、日志、故障排查困难
4. **违反微服务原则** - 应该拆分为独立进程

---

## 📞 联系方式

如有问题或建议，请提交 Issue。

---

**版本**: v2.0.0  
**更新日期**: 2025-10-16  
**兼容性**: 100% 向后兼容


# _server - 高性能 Go 服务器引擎

企业级多功能服务器框架，提供Web、API、HTTP、RPC等多种服务器类型，专注于**安全**、**性能**、**易用性**。

---

## ✨ 特性亮点

### 🔒 安全性
- ✅ **修复 CORS 漏洞** - 严格的 Origin 验证，防止反射攻击
- ✅ **修复路径遍历** - 符号链接检查，防止目录逃逸
- ✅ **防御 XSS/注入** - 参数验证和清理

### ⚡ 性能
- ✅ **实例级路由** - 消除全局变量污染
- ✅ **精确匹配优先** - O(1) 精确匹配，O(n) 正则回退
- ✅ **路由冻结** - 启动后锁定，提升并发性能
- ✅ **连接池复用** - SQL/Redis 连接池

### 🎯 易用性
- ✅ **链式API** - 流畅的配置体验
- ✅ **优雅关闭** - Context 控制，平滑退出
- ✅ **平滑启动** - 生命周期钩子
- ✅ **配置验证** - 启动前检查，快速失败

### 🏗️ 架构
- ✅ **BaseEngine** - 消除200+行重复代码
- ✅ **模块化设计** - 独立文件，清晰职责
- ✅ **线程安全** - 并发保护，无竞态条件

---

## 📦 服务器类型

| 类型 | 描述 | 状态 | 使用场景 |
|------|------|------|----------|
| **Web()** | 静态文件服务器 | ✅ 完整 | 前端资源托管 |
| **Api()** | RESTful API 服务器 | ✅ 完整 | 后端 API 服务 |
| **Http()** | 混合服务器 | ✅ 完整 | API + SPA 一体化 |
| **Rpc()** | gRPC 服务器 | ⚠️ 部分 | 微服务通信 |
| **Cli()** | 命令行工具 | 📝 计划中 | 脚本和工具 |
| **Job()** | 定时任务 | 📝 计划中 | Cron 任务 |
| **File()** | 文件处理 | 📝 计划中 | 批量文件操作 |
| **Websocket()** | WebSocket | 📝 计划中 | 实时通信 |

---

## 🚀 快速开始

### 1. 最简单的 API 服务器（推荐方式）

**路由注册** (`router/api.go`):
```go
package router

import (
    "BE/controller"
    "github.com/junyang7/go-common/_router"
)

func init() {
    // 路由在 init() 中全局注册（推荐）
    _router.Prefix("/api").Group(func() {
        _router.Get("/login", controller.Login)
        _router.Get("/users", controller.Users)
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
    // ✅ 自动使用全局路由，简单高效
}
```

### 2. 优雅关闭（生产环境推荐）

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    _ "BE/router"
    "github.com/junyang7/go-common/_server"
    "github.com/junyang7/go-common/_toml"
)

func main() {
    // 监听中断信号（Ctrl+C, kill）
    ctx, cancel := signal.NotifyContext(
        context.Background(),
        os.Interrupt,
        syscall.SIGTERM,
    )
    defer cancel()

    // 启动服务器
    server := _server.Http().
        Load(_toml.New().File("./etc/app.toml"), "server.http")

    if err := server.RunWithContext(ctx); err != nil {
        log.Fatal(err)
    }
    // ✅ 收到信号后自动优雅关闭
}
```

### 3. 平滑启动（生命周期钩子）

```go
package main

import (
    "fmt"
    "github.com/junyang7/go-common/_server"
)

func main() {
    server := _server.Api().
        Host("0.0.0.0").
        Port("8080")

    // 启动前回调（可用于预热）
    server.SetBeforeStartCallback(func() error {
        fmt.Println("🔧 准备启动服务器...")
        // 预热缓存、检查依赖等
        return nil
    })

    // 启动后回调（可用于注册服务）
    server.SetAfterStartCallback(func() {
        fmt.Println("✅ 服务器已启动！")
        // 注册到服务发现、发送就绪通知等
    })

    // 关闭前回调（可用于清理）
    server.SetBeforeStopCallback(func() {
        fmt.Println("🛑 正在关闭服务器...")
        // 拒绝新请求、等待现有请求完成
    })

    // 关闭后回调（可用于资源释放）
    server.SetAfterStopCallback(func() {
        fmt.Println("💤 服务器已关闭！")
        // 关闭数据库连接、清理临时文件等
    })

    ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
    defer cancel()

    server.RunWithContext(ctx)
}
```

### 4. 自定义路由管理器（可选，仅测试场景）

```go
package main

import (
    "testing"
    "github.com/junyang7/go-common/_router"
    "github.com/junyang7/go-common/_server"
    "github.com/junyang7/go-common/_context"
)

func TestAPIServer_Concurrent(t *testing.T) {
    t.Parallel()
    
    // 创建独立的路由管理器
    manager := _router.NewManager()
    builder := _router.NewBuilder(manager)
    
    // 显式注册路由
    builder.Get("/test", func(ctx *_context.Context) {
        ctx.JSON(map[string]string{"test": "ok"})
    })
    
    // 使用自定义管理器（路由隔离）
    server := _server.Api().
        Port("0").
        RouterManager(manager).  // ← 可选：自定义路由
        RunWithContext(ctx)
    
    // ✅ 与其他测试完全隔离，支持并发
}
```

**注意**: 
- ⚠️ 生产环境不需要使用 `RouterManager()`
- ⚠️ 使用自定义管理器后，`init()` 中的全局路由将不可用
- ✅ 主要用于测试场景的路由隔离
- 📖 详细说明见 [COMPATIBILITY.md](COMPATIBILITY.md)

---

## 🎨 API 详解

### Web 服务器

静态文件托管，适用于前端资源。

```go
_server.Web().
    Root("/var/www/html").         // 静态文件根目录
    Host("0.0.0.0").               // 监听地址
    Port("80").                    // 端口
    Debug(true).                   // 调试模式
    Run()
```

**特性：**
- ✅ 路径穿越防护（符号链接检查）
- ✅ 目录列表禁用（安全策略）
- ✅ 自动 MIME 类型检测

---

### API 服务器

RESTful API 服务，支持路由、中间件、参数绑定。

```go
_server.Api().
    Host("0.0.0.0").               // 监听地址
    Port("8080").                  // 端口
    Origin([]string{               // CORS 白名单
        "localhost",               // 精确匹配
        ".example.com",            // 子域名匹配
        "*",                       // 通配符（不建议生产环境）
    }).
    Prefix("/api/").               // API 路径前缀
    CORSHeaders([]string{          // 自定义 CORS Headers
        "content-type",
        "authorization",
        "x-custom-header",
    }).
    Router(myRouter).              // 添加路由
    Run()
```

**特性：**
- ✅ **CORS 安全修复** - 严格 Origin 验证，防止反射攻击
- ✅ **实例级路由** - 多实例隔离，无全局污染
- ✅ **路由参数** - 支持 `:id` 动态参数和正则匹配
- ✅ **中间件** - Before/After 中间件链
- ✅ **异常处理** - 统一错误响应，调试模式显示堆栈

---

### HTTP 服务器

API + 静态文件 + SPA，适用于全栈应用。

```go
_server.Http().
    Root("/var/www/dist").         // 静态文件根目录
    Host("0.0.0.0").               // 监听地址
    Port("80").                    // 端口
    Origin([]string{"localhost"}). // CORS 白名单
    Prefix("/api/").               // API 路径前缀
    Router(myRouter).              // 添加路由
    Run()
```

**路由规则：**
1. `/api/*` → API 处理器（路由匹配）
2. `/static/*`, `*.js`, `*.css` → 静态文件（直接返回）
3. 其他路径 → 尝试文件，不存在则返回 `index.html`（支持 Vue/React History 模式）

**特性：**
- ✅ **SPA 支持** - History 模式自动 fallback 到 index.html
- ✅ **智能路由** - API 和静态文件自动区分
- ✅ **安全防护** - 路径穿越、符号链接检查

---

### RPC 服务器

gRPC 服务，适用于微服务通信。

```go
_server.Rpc().
    Network("tcp").                // 网络类型
    Addr("0.0.0.0:50051").        // 监听地址
    Router(myRpcRouter).           // 添加路由
    Debug(true).                   // 调试模式
    Run()
```

**状态：** ⚠️ 部分实现，路由匹配待完善

---

## 🔐 安全最佳实践

### 1. CORS 配置

```go
// ❌ 不安全（生产环境禁用）
.Origin([]string{"*"})

// ✅ 精确匹配
.Origin([]string{
    "example.com",
    "api.example.com",
})

// ✅ 子域名匹配
.Origin([]string{
    ".example.com",  // 匹配 *.example.com
})

// ⚠️ 注意：通配符 * 不支持 credentials
```

### 2. 路径穿越防护

```go
// ✅ 自动防护（无需配置）
_server.Web().Root("/var/www").Run()

// 以下攻击会被自动阻止：
// - /../etc/passwd
// - /../../etc/passwd
// - 符号链接逃逸
```

### 3. 调试模式

```go
// ✅ 开发环境
.Debug(true)   // 显示错误堆栈、文件路径

// ✅ 生产环境
.Debug(false)  // 隐藏敏感信息
```

---

## ⚡ 性能优化

### 1. 路由性能

```go
// ✅ 精确匹配优先（O(1)）
_router.Get("/api/users", handler)

// ⚠️ 正则匹配（O(n)）
_router.Get("/api/users/:id", handler)

// 💡 建议：将高频路由定义为精确匹配
```

### 2. 连接池

```go
// ✅ 自动管理（通过 _sql.Load() 和 _redis.Load()）
_server.Api().
    Load(conf, "server.api").  // 自动初始化连接池
    Run()
```

### 3. 路由冻结

```go
// ✅ 自动冻结（Run() 时）
// 启动后路由表只读，提升并发性能
```

---

## 🧪 测试

### 运行测试

```bash
# 路由测试
cd _router && go test -v

# 服务器测试
cd _server && go test -v

# 覆盖率
go test -cover ./...
```

### 测试覆盖率

- `_router`: ✅ 100% (所有测试通过)
- `_server`: 📝 待补充

---

## 📊 性能基准

```bash
# 路由匹配性能
BenchmarkRouterMatch_Exact    10000000    150 ns/op
BenchmarkRouterMatch_Regex     1000000   1500 ns/op

# 精确匹配比正则快 10 倍
```

---

## 🆚 对比旧版本

| 特性 | 旧版本 | 新版本 |
|------|--------|--------|
| **全局变量** | ❌ RouterList 全局共享 | ✅ 实例级路由管理 |
| **CORS 安全** | ❌ Origin 反射攻击 | ✅ 严格验证 |
| **路径遍历** | ⚠️ 部分防护 | ✅ 符号链接检查 |
| **优雅关闭** | ❌ 不支持 | ✅ Context 控制 |
| **代码重复** | ❌ 200+ 行重复 | ✅ BaseEngine 复用 |
| **配置验证** | ❌ 运行时失败 | ✅ 启动前检查 |
| **路由性能** | ⚠️ O(n) 遍历 | ✅ 精确匹配 O(1) |
| **并发安全** | ⚠️ 部分安全 | ✅ 完全安全 |

---

## 🛠️ 故障排查

### 1. 端口被占用

```bash
# 查看端口占用
lsof -i :8080

# 杀死进程
kill -9 <PID>
```

### 2. CORS 错误

```go
// ✅ 确保 Origin 配置正确
.Origin([]string{"localhost", ".example.com"})

// ✅ 检查浏览器 DevTools Network 面板
// 查看 Access-Control-Allow-Origin header
```

### 3. 路由不匹配

```go
// ✅ 检查路径是否包含前缀
// API 引擎默认前缀: /api/

// ❌ 错误
_router.Get("/users", handler)  // 实际路径: /api/users

// ✅ 正确
_router.Get("/users", handler)  // 访问: /api/users
```

---

## 📚 相关包

- [`_router`](../_router/README.md) - 路由管理器
- [`_context`](../_context/README.md) - HTTP 上下文
- [`_parameter`](../_parameter/README.md) - 参数处理
- [`_conf`](../_conf/README.md) - 配置管理

---

## 🔄 迁移指南

### 从旧版本迁移

#### 1. 路由注册

```go
// ❌ 旧版本（全局路由）
_router.Get("/api/users", handler)
_server.Api().Run()

// ✅ 新版本（实例级路由）
router := _router.Get("/users", handler)
_server.Api().Router(router).Run()
```

#### 2. 优雅关闭

```go
// ❌ 旧版本（无优雅关闭）
_server.Api().Run()  // 阻塞，Ctrl+C 强制退出

// ✅ 新版本（优雅关闭）
ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
defer cancel()
_server.Api().RunWithContext(ctx)  // 支持信号量控制
```

#### 3. 测试隔离

```go
// ❌ 旧版本（全局污染）
func TestA(t *testing.T) {
    _router.Get("/a", handler)
    RouterList = []*Router{}  // 手动清空
}

// ✅ 新版本（自动隔离）
func TestA(t *testing.T) {
    _router.ResetDefaultManager()  // 一次性重置
    _router.Get("/a", handler)
}
```

---

## 🤝 贡献

欢迎提交 Issue 和 PR！

---

## 📄 许可证

MIT License

---

## 📮 联系方式

- 作者: junyang7
- 项目: https://github.com/junyang7/go-common

---

**🎉 享受高性能、安全的 Go 服务器开发！**


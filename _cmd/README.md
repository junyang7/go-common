# _cmd - 命令执行工具

简单、安全、强大的系统命令执行工具，支持简单直调和复杂构建两种模式。

---

## 🚀 快速开始

### 简单方式（90% 场景）

```go
// 执行命令
output := _cmd.Execute("ls", "-l")

// 自动类型转换
count := _cmd.ExecuteAsInt64("wc", "-l", "file.txt")
version := _cmd.ExecuteAsString("git", "rev-parse", "HEAD")

// 安全执行（不 panic）⭐ 推荐
result := _cmd.ExecuteSafe("git", "status")
if result.Success() {
    fmt.Println(result.StdoutString())
} else {
    fmt.Println("Error:", result.StderrString())
}
```

### 复杂方式（10% 场景）

```go
// 构建器模式：链式配置
result := _cmd.New("sh", "-c", "echo $VAR").
    Dir("/tmp").                 // 工作目录
    AddEnv("VAR", "hello").     // 环境变量
    Timeout(time.Second * 10).  // 超时控制
    StdinString("input").       // 标准输入
    Output()                    // 执行
```

---

## 📦 两种模式对比

| 特性 | 简单模式 | 复杂模式 |
|------|---------|---------|
| **使用方式** | 直接函数调用 | New() + 链式配置 |
| **代码行数** | 1 行 | 2-5 行 |
| **学习成本** | ⭐ 极低 | ⭐⭐ 较低 |
| **适用场景** | 日常使用 | 高级需求 |

---

## 📚 API 文档

### 简单模式 API

#### 基础执行
```go
Execute(name, args...)          // 执行命令，返回输出
ExecuteAsInt64(name, args...)   // 执行并转为 int64
ExecuteAsString(name, args...)  // 执行并转为 string
ExecuteSafe(name, args...)      // 安全执行（不 panic）⭐
ExecuteInteractive(cmd)         // 交互式执行
```

#### 快捷方式
```go
ExecuteWithTimeout(timeout, name, args...)  // 带超时
ExecuteWithContext(ctx, name, args...)      // 使用 context
ExecuteWithDir(dir, name, args...)          // 指定目录
ExecuteWithEnv(env, name, args...)          // 环境变量
```

### 复杂模式 API

#### 创建构建器
```go
New(name, args...) *Command
```

#### 配置方法（链式调用）
```go
.Dir(dir)              // 工作目录
.Env(env)              // 环境变量（替换）
.AddEnv(key, value)    // 添加环境变量
.Stdin(reader)         // 标准输入
.StdinString(s)        // 标准输入（字符串）
.Stdout(writer)        // 重定向输出
.Stderr(writer)        // 重定向错误
.Timeout(duration)     // 超时时间
.Context(ctx)          // Context 控制
```

#### 执行方法
```go
.Output() *Result      // 执行并捕获输出
.Run() *Result         // 执行（不捕获输出）
.MustOutput() []byte   // 执行（失败 panic）
.MustRun()             // 执行（失败 panic）
```

#### 辅助功能
```go
Pipe(cmd1, cmd2) *Result  // 命令管道
```

### Result 方法
```go
.Success() bool           // 是否成功
.StdoutString() string    // 输出字符串
.StderrString() string    // 错误字符串
.StdoutInt64() int64      // 输出 int64
.StdoutLines() []string   // 按行分割
.StderrLines() []string   // 错误按行
```

---

## 💡 使用示例

### 1. 简单命令执行

```go
// 最简单的用法
output := _cmd.Execute("date")
fmt.Println(string(output))
```

### 2. 类型转换

```go
// 自动转换为 int64
lineCount := _cmd.ExecuteAsInt64("wc", "-l", "file.txt")

// 自动转换为 string
gitHash := _cmd.ExecuteAsString("git", "rev-parse", "HEAD")
branch := _cmd.ExecuteAsString("git", "branch", "--show-current")
```

### 3. 错误处理

```go
// 推荐：使用 ExecuteSafe
result := _cmd.ExecuteSafe("git", "status")
if result.Success() {
    fmt.Println("Output:", result.StdoutString())
} else {
    fmt.Printf("Error: %s\n", result.StderrString())
    fmt.Printf("Exit code: %d\n", result.ExitCode)
}
```

### 4. 超时控制

```go
// 简单方式
output := _cmd.ExecuteWithTimeout(
    time.Second * 10,
    "curl", "https://example.com",
)

// 复杂方式（需要更多配置）
result := _cmd.New("curl", "https://example.com").
    Timeout(time.Second * 10).
    Output()
```

### 5. 环境变量

```go
// 简单方式：完全替换
output := _cmd.ExecuteWithEnv(
    []string{"PATH=/bin"},
    "which", "ls",
)

// 复杂方式：添加变量
result := _cmd.New("sh", "-c", "echo $MY_VAR").
    AddEnv("MY_VAR", "hello").
    Output()
```

### 6. 标准输入

```go
// 字符串输入
result := _cmd.New("grep", "pattern").
    StdinString("line1\nline2\nline3").
    Output()

// Reader 输入
input := bytes.NewReader(data)
result := _cmd.New("cat").Stdin(input).Output()
```

### 7. 管道操作

```go
// 模拟: echo "hello\nworld" | grep "world"
cmd1 := _cmd.New("echo", "hello\nworld")
cmd2 := _cmd.New("grep", "world")
result := _cmd.Pipe(cmd1, cmd2)

fmt.Println(result.StdoutString())
```

### 8. 复杂组合

```go
// 多个配置组合
result := _cmd.New("deploy.sh").
    Dir("/app").
    AddEnv("ENV", "production").
    AddEnv("DEBUG", "false").
    Timeout(time.Minute * 5).
    Output()

if result.Success() {
    fmt.Println("Deploy success!")
} else {
    fmt.Printf("Deploy failed: %s\n", result.StderrString())
}
```

### 9. 逐行处理

```go
// 获取所有 Docker 容器
result := _cmd.ExecuteSafe("docker", "ps", "-a")
if result.Success() {
    for i, line := range result.StdoutLines() {
        fmt.Printf("Container %d: %s\n", i+1, line)
    }
}
```

---

## 🎯 使用决策树

```
需要执行命令？
│
├─ 命令很简单（无特殊配置）？
│  ├─ 是 → 用 Execute() / ExecuteAsString()
│  └─ 否 → 继续判断
│
├─ 命令可能失败？
│  ├─ 是 → 用 ExecuteSafe()
│  └─ 否 → 继续判断
│
├─ 需要 1-2 个配置？
│  ├─ 是 → 用 ExecuteWith* 系列
│  └─ 否 → 继续判断
│
└─ 需要 3+ 个配置或组合？
   └─ 是 → 用 New() 构建器
```

---

## 🔒 安全性

### ✅ 防命令注入

```go
// ✅ 安全：参数独立传递
func SafeGrep(userInput, filename string) {
    result := _cmd.ExecuteSafe("grep", userInput, filename)
    // userInput 作为参数，不会被解释为命令
}

// ❌ 危险：字符串拼接（不要这样做！）
func UnsafeGrep(userInput string) {
    // 如果 userInput = "; rm -rf /"，会执行删除命令！
    _cmd.ExecuteInteractive("grep " + userInput)
}
```

### ✅ 错误处理

```go
// ✅ 推荐：使用 ExecuteSafe
result := _cmd.ExecuteSafe("command")
if !result.Success() {
    // 优雅处理错误
    log.Println("Error:", result.StderrString())
}

// ❌ 避免：直接 Execute（会 panic）
output := _cmd.Execute("risky-command") // 失败会 panic
```

---

## 📊 性能

### 基准测试结果

```
BenchmarkExecute           626      1.86 ms/op    45 KB/op
BenchmarkExecuteSafe       465      2.35 ms/op    13 KB/op
BenchmarkCommandBuilder    464      2.43 ms/op    13 KB/op
```

**结论**：三种方式性能接近，可以放心使用任何一种。

---

## 🎓 最佳实践

### ✅ DO（推荐）

```go
// 1. 简单场景用简单方式
output := _cmd.Execute("date")

// 2. 可能失败用 Safe
result := _cmd.ExecuteSafe("risky-command")

// 3. 参数独立传递（防注入）
_cmd.ExecuteSafe("grep", userInput, filename)

// 4. 复杂场景用构建器
result := _cmd.New("cmd").
    Dir("/tmp").
    AddEnv("VAR", "value").
    Output()

// 5. 检查错误和退出码
if !result.Success() {
    log.Printf("Exit code: %d\n", result.ExitCode)
    log.Printf("Error: %s\n", result.StderrString())
}
```

### ❌ DON'T（避免）

```go
// 1. 不要拼接用户输入
_cmd.ExecuteInteractive("grep " + userInput) // 危险！

// 2. 不要忽略错误
_cmd.Execute("risky-command") // 可能 panic

// 3. 不要过度使用构建器
_cmd.New("echo", "hello").Output() // 简单命令直接用 Execute

// 4. 不要无超时执行长命令
_cmd.Execute("long-running-task") // 可能永久挂起
```

---

## 🌟 特点

- ✅ **简单优先**：90% 场景一行代码搞定
- ✅ **渐进增强**：复杂需求用构建器
- ✅ **安全可靠**：防注入、完整错误处理
- ✅ **性能优异**：~2ms 执行，13KB 内存
- ✅ **测试完善**：24 个测试用例，全部通过
- ✅ **生产就绪**：真实场景验证

---

## 📖 实际场景

### Git 操作

```go
// 获取当前分支
branch := _cmd.ExecuteAsString("git", "branch", "--show-current")

// 提交代码
result := _cmd.ExecuteSafe("git", "commit", "-m", message)
if !result.Success() {
    log.Printf("Commit failed: %s\n", result.StderrString())
}
```

### Docker 操作

```go
// 检查容器状态
result := _cmd.ExecuteSafe("docker", "inspect", "-f", 
    "{{.State.Running}}", containerName)
isRunning := strings.TrimSpace(result.StdoutString()) == "true"

// 获取日志
logs := _cmd.New("docker", "logs", "--tail", "100", containerName).
    Timeout(time.Second * 5).
    Output()
```

### 文件处理

```go
// 统计文件行数
count := _cmd.ExecuteAsInt64("wc", "-l", "file.txt")

// 搜索文件
result := _cmd.New("find", ".", "-name", "*.go").
    Dir("/app").
    Output()

for _, file := range result.StdoutLines() {
    fmt.Println("Found:", file)
}
```

---

## 📦 完整示例

```go
package main

import (
    "fmt"
    "_cmd"
    "time"
)

func main() {
    // 简单使用
    fmt.Println("=== Simple Usage ===")
    date := _cmd.ExecuteAsString("date")
    fmt.Println("Date:", date)
    
    // 安全执行
    fmt.Println("\n=== Safe Execution ===")
    result := _cmd.ExecuteSafe("ls", "-l")
    if result.Success() {
        fmt.Println("Files:")
        for _, line := range result.StdoutLines() {
            fmt.Println(" ", line)
        }
    }
    
    // 复杂配置
    fmt.Println("\n=== Complex Configuration ===")
    result = _cmd.New("sh", "-c", "echo Hello $NAME").
        AddEnv("NAME", "World").
        Timeout(time.Second * 5).
        Output()
    
    if result.Success() {
        fmt.Println("Output:", result.StdoutString())
    }
}
```

---

## 🆚 对比其他工具

| 特性 | _cmd | os/exec | sh/exec |
|------|------|---------|---------|
| **简单性** | ⭐⭐⭐⭐⭐ 一行搞定 | ⭐⭐⭐ 需要多行 | ⭐⭐⭐⭐ 较简单 |
| **类型转换** | ✅ 内置 | ❌ 需手动 | ❌ 需手动 |
| **错误处理** | ✅ Safe 模式 | ⚠️ 需手动 | ⚠️ 需手动 |
| **构建器** | ✅ 链式调用 | ❌ 无 | ❌ 无 |
| **管道** | ✅ Pipe 函数 | ⚠️ 需手动 | ✅ 有 |
| **学习成本** | ⭐ 极低 | ⭐⭐⭐ 较高 | ⭐⭐ 较低 |

---

## 📝 总结

`_cmd` 包是一个：
- ✅ 简单易用的命令执行工具
- ✅ 支持简单和复杂两种模式
- ✅ 安全可靠，防注入防崩溃
- ✅ 性能优异，生产就绪
- ✅ 覆盖 95%+ 的日常需求

**推荐指数：⭐⭐⭐⭐⭐**

---

## 📚 更多文档

- [EVALUATION.md](./EVALUATION.md) - 完整评估报告
- [cmd_test.go](./cmd_test.go) - 24 个测试示例

---

**License:** MIT  
**Version:** 2.0  
**Status:** Production Ready ✅

# _conf - 配置接口

配置管理的抽象接口，支持多种配置格式（JSON、TOML 等）。

---

## 🎯 作用

定义统一的配置接口，业务代码通过接口访问配置，不关心具体格式。

---

## 📦 接口定义

```go
type Conf interface {
    Byte(byte []byte) Conf              // 从字节加载
    Text(text string) Conf              // 从字符串加载
    File(path string) Conf              // 从文件加载
    Get(path string) *_parameter.Parameter  // 获取配置值
}
```

---

## 💡 使用方式

### 1. 加载配置

```go
import (
    "_conf"
    "_json"  // 或 "_toml"
)

func init() {
    // 从文件加载（JSON）
    _conf.Load(_json.New().File("config.json"))
    
    // 或从文件加载（TOML）
    _conf.Load(_toml.New().File("config.toml"))
    
    // 或从字符串加载
    _conf.Load(_json.New().Text(`{"host":"localhost","port":3306}`))
}
```

### 2. 获取配置

```go
// 获取字符串配置
host := _conf.Get("database.host").String().Value()

// 获取数字配置
port := _conf.Get("database.port").Int64().Value()

// 获取布尔配置
debug := _conf.Get("debug").Bool().Value()

// 使用点号访问嵌套配置
apiKey := _conf.Get("api.key").String().Value()
```

### 3. 设置默认值

```go
// 如果配置不存在，使用默认值
host := _conf.Get("host").Default("localhost").String().Value()
port := _conf.Get("port").Default(3306).Int64().Value()
```

---

## 📋 配置格式示例

### JSON 格式

```json
{
  "database": {
    "host": "localhost",
    "port": 3306
  },
  "debug": true
}
```

```go
_conf.Load(_json.New().File("config.json"))
host := _conf.Get("database.host").String().Value()  // "localhost"
```

### TOML 格式

```toml
debug = true

[database]
host = "localhost"
port = 3306
```

```go
_conf.Load(_toml.New().File("config.toml"))
port := _conf.Get("database.port").Int64().Value()  // 3306
```

---

## 🔧 API

### Load

```go
func Load(conf Conf)
```

加载配置实例（全局单例）。

**参数：**
- `conf`: 实现了 Conf 接口的配置对象

**示例：**
```go
_conf.Load(_json.New().File("config.json"))
```

### Get

```go
func Get(path string) *_parameter.Parameter
```

获取配置值。如果配置未加载会抛出异常。

**参数：**
- `path`: 配置路径，使用点号分隔（如 "database.host"）

**返回：**
- `*_parameter.Parameter`: 参数对象，可转换为具体类型

**示例：**
```go
value := _conf.Get("server.port").Int64().Value()
```

### Reset

```go
func Reset()
```

重置配置（清空），测试时可用。

**示例：**
```go
_conf.Reset()
```

### IsLoaded

```go
func IsLoaded() bool
```

检查配置是否已加载。

**返回：**
- `bool`: true 表示已加载，false 表示未加载

**示例：**
```go
if !_conf.IsLoaded() {
    _conf.Load(_json.New().File("config.json"))
}
```

---

## ⚡ 实现

框架提供的配置格式实现：

| 包 | 格式 | 说明 |
|-----|------|------|
| `_json` | JSON | `_json.New()` |
| `_toml` | TOML | `_toml.New()` |

---

## 📖 完整示例

```go
package main

import (
    "_conf"
    "_json"
)

func main() {
    // 1. 加载配置
    _conf.Load(_json.New().File("config.json"))
    
    // 2. 获取配置
    dbHost := _conf.Get("database.host").String().Value()
    dbPort := _conf.Get("database.port").Int64().Value()
    debug := _conf.Get("debug").Default(false).Bool().Value()
    
    // 3. 使用配置
    println("Database:", dbHost, dbPort)
    println("Debug:", debug)
}
```

**config.json:**
```json
{
  "database": {
    "host": "localhost",
    "port": 3306
  },
  "debug": true
}
```

---

## 💡 设计理念

```
接口抽象：
├─ _conf 定义接口（不关心格式）
├─ _json 实现 JSON 解析
└─ _toml 实现 TOML 解析

业务代码：
└─ 只依赖 _conf 接口，不依赖具体实现
```

---

## ⚠️ 注意

1. **线程安全**: 所有方法都是线程安全的
2. **全局单例**: 配置在整个应用中共享
3. **初始化顺序**: 在使用前必须先 Load 配置
4. **错误处理**: 如果未 Load 就 Get 会抛出异常
5. **测试隔离**: 测试时可使用 Reset() 清空配置

---

**License:** MIT  
**Version:** 1.0


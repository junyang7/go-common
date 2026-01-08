# _parameter - 参数包装器

HTTP 参数的统一包装器，支持默认值、类型转换、验证等功能。

---

## 🎯 作用

将 HTTP 请求参数包装为 `Parameter` 对象，提供统一的操作接口。

---

## 💡 使用

```go
import "_parameter"

// 创建参数
param := _parameter.New("age", 30)

// 获取原始值
value := param.Value()  // interface{}

// 类型转换（返回 Validator）
age := param.Int64().Value()  // int64

// 使用默认值
param := _parameter.New("age", nil)
age := param.Default(18).Int64().Value()  // 18

// 链式调用 + 验证
age := param.Int64().EnsureMin(0).EnsureMax(150).Value()

// 必填验证
name := param.Required().String().Value()
```

---

## 📦 API 文档

### 基础方法

| 方法 | 说明 | 返回 |
|------|------|------|
| `New(name, value)` | 创建参数 | `*Parameter` |
| `Default(value)` | 设置默认值（仅 nil 时） | `*Parameter` |
| `Required()` | 必填验证（nil 时抛出异常） | `*Parameter` |
| `Value()` | 获取原始值 | `interface{}` |
| `IsNil()` | 检查是否为 nil | `bool` |

### 类型转换（返回 Validator）

| 方法 | 说明 | 返回 |
|------|------|------|
| `Int64()` | 转为 Int64 Validator | `*_validator.Int64` |
| `String()` | 转为 String Validator | `*_validator.String` |
| `Bool()` | 转为 Bool Validator | `*_validator.Bool` |
| `Float64()` | 转为 Float64 Validator | `*_validator.Float64` |

### 文件处理

| 方法 | 说明 | 返回 |
|------|------|------|
| `File()` | 获取第一个文件 | `*multipart.FileHeader` |
| `FileList()` | 获取文件列表 | `[]*multipart.FileHeader` |

---

## 📖 使用示例

### 基础用法

```go
// 创建参数
param := _parameter.New("age", 30)

// 获取值
value := param.Value()  // interface{} = 30

// 类型转换
age := param.Int64().Value()  // int64(30)
```

### 默认值（重要）

```go
// 值为 nil，使用默认值
param := _parameter.New("age", nil)
age := param.Default(18).Int64().Value()  // 18

// 值不为 nil，保持原值
param := _parameter.New("age", 30)
age := param.Default(18).Int64().Value()  // 30（不使用默认值）

// ⚠️ 注意：0、空字符串、false 都是有效值
param := _parameter.New("count", 0)
count := param.Default(10).Int64().Value()  // 0（不是 nil，不使用默认值）
```

### 必填验证

```go
// 值存在，验证通过
param := _parameter.New("name", "alice")
name := param.Required().String().Value()  // "alice"

// 值为 nil，抛出异常
param := _parameter.New("name", nil)
param.Required()  // ❌ panic: "name is required"
```

### 类型转换 + 验证

```go
// 使用 Validator 进行验证
age := param.Int64().
    EnsureMin(0).
    EnsureMax(150).
    Value()

name := param.String().
    EnsureLengthMin(2).
    EnsureLengthMax(50).
    Value()

email := param.String().
    EnsureEmail().
    Value()
```

### 文件处理

```go
// 获取第一个文件（常用）
file := param.File()
if file != nil {
    filename := file.Filename
    size := file.Size
    
    f, err := file.Open()
    defer f.Close()
    // 处理文件...
}

// 获取所有文件
files := param.FileList()
if files != nil {
    for _, file := range files {
        // 处理每个文件...
    }
}

// 在 _context 中使用
file := ctx.File("avatar").File()
files := ctx.File("avatar").FileList()

// 必填验证
file := ctx.File("avatar").Required().File()
```

---

## 🎓 实际场景

### 场景 1：HTTP 参数（配合 _context）

```go
func Handler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 使用默认值
    page := ctx.Get("page").Default(1).Int64().Value()
    size := ctx.Get("size").Default(10).Int64().Value()
    
    // 必填验证
    name := ctx.Post("name").Required().String().Value()
    
    // 类型转换 + 验证
    age := ctx.Post("age").Int64().
        EnsureMin(0).
        EnsureMax(150).
        Value()
    
    // 文件上传
    file := ctx.File("avatar").File()
    if file != nil {
        // 处理文件...
    }
}
```

### 场景 2：配置读取（配合 _conf）

```go
// 使用默认值
timeout := _conf.Get("timeout").Default(30).Int64().Value()
debug := _conf.Get("debug").Default(false).Bool().Value()
host := _conf.Get("host").Default("localhost").String().Value()
```

### 场景 3：分页参数

```go
func ListHandler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 分页参数（带默认值和验证）
    page := ctx.Get("page").
        Default(1).
        Int64().
        EnsureMin(1).
        Value()
    
    size := ctx.Get("size").
        Default(10).
        Int64().
        EnsureMin(1).
        EnsureMax(100).
        Value()
    
    // 查询数据...
}
```

### 场景 4：文件上传

```go
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 单文件上传（必填）
    file := ctx.File("avatar").Required().File()
    filename := file.Filename
    
    // 多文件上传
    files := ctx.File("attachments").FileList()
    for _, file := range files {
        // 处理每个文件...
    }
}
```

---

## ⚠️ 重要说明

### Default 只对 nil 生效

```go
// ✅ 正确理解
New("age", nil).Default(18)    // 使用默认值 18
New("age", 30).Default(18)     // 保持 30（不是 nil）
New("age", 0).Default(18)      // 保持 0（0 不是 nil）
New("name", "").Default("x")   // 保持 ""（空字符串不是 nil）

// 核心：Default 只检查 nil，不检查空值
```

### Validator 提供验证能力

```go
// Parameter 本身不验证，只是包装
param := New("age", -1)
value := param.Value()  // -1（不验证）

// Validator 提供验证
age := param.Int64().EnsureMin(0).Value()  // ❌ panic（值小于 0）
```

### 文件方法的区别

```go
// File() - 获取第一个文件（90% 场景）
file := param.File()  // *multipart.FileHeader

// FileList() - 获取所有文件（多文件上传）
files := param.FileList()  // []*multipart.FileHeader
```

---

## 📊 性能

```
New()      0.67 ns/op    0 B/op    0 allocs/op
String()  25.79 ns/op   32 B/op    1 allocs/op
Int64()   25.79 ns/op   32 B/op    1 allocs/op
```

**评价：** ⭐⭐⭐⭐⭐ 性能优异

---

## 💡 设计理念

### 职责定位

```
Parameter（参数包装器）：
├─ 包装原始值
├─ 提供默认值
├─ 必填验证
└─ 转换为 Validator

Validator（类型验证器）：
├─ 类型转换
├─ 范围验证
├─ 格式验证
└─ 返回最终值
```

### 使用流程

```
1. Parameter 包装原始值
   ↓
2. 设置默认值/必填验证
   ↓
3. 转换为 Validator
   ↓
4. 执行验证规则
   ↓
5. 获取最终值
```

---

**License:** MIT  
**Version:** 2.0  
**Status:** Production Ready ✅

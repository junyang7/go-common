# _codeMessage - 错误码消息

框架内置的错误码消息结构。

---

## 📦 结构

```go
type CodeMessage struct {
    Code    int
    Message string
}
```

---

## 🔧 框架内置

```go
ErrNone    = New(0, "success")                     // 成功
ErrDefault = New(-1, "something goes wrong...!!!") // 默认错误
```

---

## 💡 使用

### 创建错误码

```go
import "_codeMessage"

// 自定义错误码
var ErrUserNotFound = _codeMessage.New(1001, "用户不存在")
var ErrOrderInvalid = _codeMessage.New(2001, "订单无效")
```

### 使用内置错误

```go
// 成功
response.Code = _codeMessage.ErrNone.Code
response.Message = _codeMessage.ErrNone.Message

// 默认错误
response.Code = _codeMessage.ErrDefault.Code
response.Message = _codeMessage.ErrDefault.Message
```

---

## 📖 API

### New

```go
func New(code int, message string) *CodeMessage
```

创建一个错误码消息。

**参数：**
- `code`: 错误码
- `message`: 错误消息

**返回：**
- `*CodeMessage`: 错误码消息指针

**示例：**

```go
err := _codeMessage.New(404, "not found")
fmt.Printf("Code: %d, Message: %s\n", err.Code, err.Message)
```

---

## 📊 内置错误码

| 错误码 | 名称 | 消息 | 说明 |
|--------|------|------|------|
| 0 | `ErrNone` | success | 成功 |
| -1 | `ErrDefault` | something goes wrong...!!! | 默认错误 |

---

**License:** MIT  
**Version:** 1.0

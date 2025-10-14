# _base64Format - URL 安全的 Base64 编码工具

## 功能说明

提供 **URL 安全** 且 **无填充** 的 Base64 编码和解码功能。

### 核心特性

- ✅ **URL 安全**：使用 `-` 和 `_` 替代 `+` 和 `/`
- ✅ **无填充字符**：不使用 `=` 填充，输出更简洁
- ✅ **适合 URL**：可以直接用在 URL 参数、Cookie、JWT、文件名中

---

## 函数

### _base64Format.Encode

#### 函数签名
```go
func Encode(data string) string
```

#### 功能说明
将字符串编码为 URL 安全的 Base64 格式。

#### 参数
- `data`: 待编码的字符串

#### 返回值
- URL 安全的 Base64 编码字符串（无 `=` 填充）

#### 示例
```go
result := _base64Format.Encode("hello world!")
// 输出: "aGVsbG8gd29ybGQh"
```

---

### _base64Format.Decode

#### 函数签名
```go
func Decode(data string) string
```

#### 功能说明
将 URL 安全的 Base64 字符串解码为原始字符串。

#### 参数
- `data`: URL 安全的 Base64 编码字符串

#### 返回值
- 解码后的原始字符串

#### 示例
```go
result := _base64Format.Decode("aGVsbG8gd29ybGQh")
// 输出: "hello world!"
```

---

## 使用示例

### 基本用法

```go
package main

import (
    "_base64Format"
    "fmt"
)

func main() {
    // 编码
    plaintext := "hello world!"
    encoded := _base64Format.Encode(plaintext)
    fmt.Println("编码:", encoded)
    // 输出: aGVsbG8gd29ybGQh
    
    // 解码
    decoded := _base64Format.Decode(encoded)
    fmt.Println("解码:", decoded)
    // 输出: hello world!
}
```

### 编码示例

| 原始字符串 | 编码结果 |
|-----------|---------|
| hello world! | aGVsbG8gd29ybGQh |
| 您好，中国！ | 5oKo5aW977yM5Lit5Zu977yB |
| hello 中国 | aGVsbG8g5Lit5Zu9 |

### 解码示例

| Base64 字符串 | 解码结果 |
|--------------|---------|
| aGVsbG8gd29ybGQh | hello world! |
| 5oKo5aW977yM5Lit5Zu977yB | 您好，中国！ |
| aGVsbG8g5Lit5Zu9 | hello 中国 |

---

## 🆚 对比标准编码

### 编码格式对比

| 输入 | 标准 Base64 | URL Base64 | Raw URL Base64（本包）|
|------|------------|------------|---------------------|
| test+data | dGVzdCtkYXRh | dGVzdCtkYXRh | dGVzdCtkYXRh |
| test/data | dGVzdC9kYXRh | dGVzdC9kYXRh | dGVzdC9kYXRh |
| hello!! | aGVsbG8hIQ== | aGVsbG8hIQ== | aGVsbG8hIQ ✅ 无 = |
| a+b/c | YStiL2M= | YStiL2M= | YStiL2M ✅ 无 = |

### 特性对比

| 特性 | 标准 Base64 | 本包（RawURLEncoding）|
|------|------------|---------------------|
| **字符集** | A-Z, a-z, 0-9, `+`, `/` | A-Z, a-z, 0-9, `-`, `_` ✅ |
| **填充** | 使用 `=` | 无 `=` ✅ |
| **URL 安全** | ❌ 否（`+` `/` 需要转义）| ✅ 是 |
| **文件名安全** | ❌ 否（`/` 是路径分隔符）| ✅ 是 |

---

## 使用场景

### ✅ 适合的场景

```go
✅ URL 参数
url := "https://example.com?token=" + encoded  // 不需要再 URL 编码

✅ Cookie
cookie := "session=" + encoded  // 可以直接使用

✅ JWT Token
header.payload.signature  // 所有部分都用 RawURLEncoding

✅ 文件名
filename := encoded + ".dat"  // 不包含 / 等非法字符

✅ HTML 表单
<input type="hidden" value="{{.Encoded}}">  // 无需 HTML 转义
```

### ⚠️ 不适合的场景

```go
⚠️ 与其他系统对接（如果对方期望标准 Base64）
   需要确认对方使用的编码类型

⚠️ 邮件传输
   某些邮件系统可能期望标准 Base64
```

---

## 技术细节

### 编码原理

```
输入: "hello"
二进制: 01101000 01100101 01101100 01101100 01101111
分组(6位): 011010 000110 010101 101100 011011 000110 1111(补0) → 0110 1111 00(补0)
编码: a G V s b G 8

标准: aGVsbG8=    (需要 = 填充)
Raw:  aGVsbG8     (无填充) ✅
```

### URL 安全字符替换

```
标准 Base64 → URL 安全 Base64
    +      →      -
    /      →      _
    =      →    (删除)
```

---

## 注意事项

### 1. 编码可逆性
- ✅ 任何通过 `Encode` 编码的数据，都可以用 `Decode` 完全还原
- ✅ 支持任意字符（包括中文、特殊字符、二进制数据）

### 2. 错误处理
- ⚠️ `Decode` 函数如果输入非法 Base64 字符串，会触发错误
- ⚠️ 确保只解码通过本包 `Encode` 编码的数据

### 3. 兼容性
- ✅ 与标准 `base64.RawURLEncoding` 完全兼容
- ✅ 可与其他语言的 RawURLEncoding 互操作
- ⚠️ 与标准 Base64（带填充）不兼容

---

## 依赖

- 标准库：`encoding/base64`
- 项目内：`_interceptor`（错误处理）

---

## 总评

**_base64Format 是一个：**
- ✅ 功能明确的工具包
- ✅ 实现正确且简洁
- ✅ 适合 Web 开发场景
- ⚠️ 缺少文档说明

**评分：** ⭐⭐⭐⭐ (4/5) - 补充文档后可达 5/5


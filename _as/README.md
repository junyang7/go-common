# _as 包 - 类型转换工具

提供简洁、安全的 Go 语言类型转换功能，支持常见类型之间的相互转换。

## 📑 目录

- [String](#_asstring---转换为字符串)
- [Bool](#_asbool---转换为布尔值)
- [Float64](#_asfloat64---转换为浮点数)
- [ByteList](#_asbytelist---转换为字节数组)
- [Int64](#_asint64---转换为-int64)
- [Int](#_asint---转换为-int)
- [Uint64](#_asuint64---转换为-uint64)
- [Uint](#_asuint---转换为-uint)

---

## _as.String - 转换为字符串

### 函数签名
```go
func String(value interface{}) string
```

### 功能说明
将任意类型的值转换为字符串类型。

### 转换规则

- ① **string 类型**：保持不变
- ② **[]byte 类型**：通过 string 进行类型转换
- ③ **nil 类型**：转换成空字符串
- ④ **number 类型**：直接转换成字面量
- ⑤ **bool 类型**：直接转换成字面量
- ⑥ **其他类型**：通过 `fmt.Sprintf("%v", v)` 转换成可视化字符串形式

### 转换示例

| 转换前类型                                                            | 转换后                    |
|------------------------------------------------------------------|------------------------|
| string("-1")                                                     | string("-1")           |
| string("0")                                                      | string("0")            |
| string("1")                                                      | string("1")            |
| string("-3.141592")                                              | string("-3.141592")    |
| string("3.141592")                                               | string("3.141592")     |
| string("")                                                       | string("")             |
| string("hell word!")                                             | string("hell word!")   |
| string("true")                                                   | string("true")         |
| string("false")                                                  | string("false")        |
| string("1A")                                                     | string("1A")           |
| string("3.141592b")                                              | string("3.141592b")    |
| []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33} | string("hello world!") |
| []byte{}                                                         | string("")             |
| []byte{51, 46, 49, 52, 49, 53, 57, 50}                           | string("3.141592")     |
| []byte{45, 51, 46, 49, 52, 49, 53, 57, 50}                       | string("-3.141592")    |
| []byte{49}                                                       | string("1")            |
| []byte{45, 49}                                                   | string("-1")           |
| nil                                                              | string("")             |
| int(-1)                                                          | string("-1")           |
| int(0)                                                           | string("0")            |
| int(1)                                                           | string("1")            |
| int8(-1)                                                         | string("-1")           |
| int8(0)                                                          | string("0")            |
| int8(1)                                                          | string("1")            |
| int16(-1)                                                        | string("-1")           |
| int16(0)                                                         | string("0")            |
| int16(1)                                                         | string("1")            |
| int32(-1)                                                        | string("-1")           |
| int32(0)                                                         | string("0")            |
| int32(1)                                                         | string("1")            |
| int64(-1)                                                        | string("-1")           |
| int64(0)                                                         | string("0")            |
| int64(1)                                                         | string("1")            |
| uint(0)                                                          | string("0")            |
| uint(1)                                                          | string("1")            |
| uint8(0)                                                         | string("0")            |
| uint8(1)                                                         | string("1")            |
| uint16(0)                                                        | string("0")            |
| uint16(1)                                                        | string("1")            |
| uint32(0)                                                        | string("0")            |
| uint32(1)                                                        | string("1")            |
| uint64(0)                                                        | string("0")            |
| uint64(1)                                                        | string("1")            |
| float32(-1)                                                      | string("-1")           |
| float32(0)                                                       | string("0")            |
| float32(1)                                                       | string("1")            |
| float32(-3.141592)                                               | string("-3.141592")    |
| float32(3.141592)                                                | string("3.141592")     |
| float64(-1)                                                      | string("-1")           |
| float64(0)                                                       | string("0")            |
| float64(1)                                                       | string("1")            |
| float64(-3.141592)                                               | string("-3.141592")    |
| float64(3.141592)                                                | string("3.141592")     |
| bool(false)                                                      | string("false")        |
| bool(true)                                                       | string("true")         |
| -                                                                | fmt.Sprintf("%v", v)   |

---

## _as.Bool - 转换为布尔值

### 函数签名
```go
func Bool(value interface{}) bool
```

### 功能说明
将任意类型的值转换为布尔类型。

### 转换规则

- ① **bool 类型**：保持不变
- ② **string 类型**：空字符串转换成 `false`，其他转换成 `true`
- ③ **[]byte 类型**：空数组转换成 `false`，其他转换成 `true`
- ④ **nil 类型**：转换成 `false`
- ⑤ **number 类型**：`0` 转换成 `false`，其他转换成 `true`
- ⑥ **其他类型**：转换成 `false`

### 转换示例

| 转换前类型                                                            | 转换后         |
|------------------------------------------------------------------|-------------|
| bool(false)                                                      | bool(false) |
| bool(true)                                                       | bool(true)  |
| string("-1")                                                     | bool(true)  |
| string("0")                                                      | bool(true)  |
| string("1")                                                      | bool(true)  |
| string("-3.141592")                                              | bool(true)  |
| string("3.141592")                                               | bool(true)  |
| string("")                                                       | bool(false) |
| string("hell word!")                                             | bool(true)  |
| string("true")                                                   | bool(true)  |
| string("false")                                                  | bool(true)  |
| string("1A")                                                     | bool(true)  |
| string("3.141592b")                                              | bool(true)  |
| []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33} | bool(true)  |
| []byte{}                                                         | bool(false) |
| []byte{51, 46, 49, 52, 49, 53, 57, 50}                           | bool(true)  |
| []byte{45, 51, 46, 49, 52, 49, 53, 57, 50}                       | bool(true)  |
| []byte{49}                                                       | bool(true)  |
| []byte{45, 49}                                                   | bool(true)  |
| nil                                                              | bool(false) |
| int(-1)                                                          | bool(true)  |
| int(0)                                                           | bool(false) |
| int(1)                                                           | bool(true)  |
| int8(-1)                                                         | bool(true)  |
| int8(0)                                                          | bool(false) |
| int8(1)                                                          | bool(true)  |
| int16(-1)                                                        | bool(true)  |
| int16(0)                                                         | bool(false) |
| int16(1)                                                         | bool(true)  |
| int32(-1)                                                        | bool(true)  |
| int32(0)                                                         | bool(false) |
| int32(1)                                                         | bool(true)  |
| int64(-1)                                                        | bool(true)  |
| int64(0)                                                         | bool(false) |
| int64(1)                                                         | bool(true)  |
| uint(0)                                                          | bool(false) |
| uint(1)                                                          | bool(true)  |
| uint8(0)                                                         | bool(false) |
| uint8(1)                                                         | bool(true)  |
| uint16(0)                                                        | bool(false) |
| uint16(1)                                                        | bool(true)  |
| uint32(0)                                                        | bool(false) |
| uint32(1)                                                        | bool(true)  |
| uint64(0)                                                        | bool(false) |
| uint64(1)                                                        | bool(true)  |
| float32(-1)                                                      | bool(true)  |
| float32(0)                                                       | bool(false) |
| float32(1)                                                       | bool(true)  |
| float32(-3.141592)                                               | bool(true)  |
| float32(3.141592)                                                | bool(true)  |
| float64(-1)                                                      | bool(true)  |
| float64(0)                                                       | bool(false) |
| float64(1)                                                       | bool(true)  |
| float64(-3.141592)                                               | bool(true)  |
| float64(3.141592)                                                | bool(true)  |
| -                                                                | bool(false) |

---

## _as.Float64 - 转换为浮点数

### 函数签名
```go
func Float64(value interface{}) float64
```

### 功能说明
将任意类型的值转换为 64 位浮点数类型。

### 转换规则

- ① **float64 类型**：保持不变
- ② **string 类型**：如果字面量是数字，转换成对应的数字，否则转换成 `0`
- ③ **[]byte 类型**：先通过 string 转换成字符串字面量，然后按字符串字面量解析规则转换
- ④ **number 类型**：直接通过 float64 转换
- ⑤ **bool 类型**：`true` 转换成 `1`，`false` 转换成 `0`
- ⑥ **其他类型**：转换成 `0`

### 转换示例

| 转换前类型                                                            | 转换后                |
|------------------------------------------------------------------|--------------------|
| float64(-1)                                                      | float64(-1)        |
| float64(0)                                                       | float64(0)         |
| float64(1)                                                       | float64(1)         |
| float64(-3.141592)                                               | float64(-3.141592) |
| float64(3.141592)                                                | float64(3.141592)  |
| string("-1")                                                     | float64(-1)        |
| string("0")                                                      | float64(0)         |
| string("1")                                                      | float64(1)         |
| string("-3.141592")                                              | float64(-3.141592) |
| string("3.141592")                                               | float64(3.141592)  |
| string("")                                                       | float64(0)         |
| string("hell word!")                                             | float64(0)         |
| string("true")                                                   | float64(0)         |
| string("false")                                                  | float64(0)         |
| string("1A")                                                     | float64(0)         |
| string("3.141592b")                                              | float64(0)         |
| []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33} | float64(0)         |
| []byte{}                                                         | float64(0)         |
| []byte{51, 46, 49, 52, 49, 53, 57, 50}                           | float64(3.141592)  |
| []byte{45, 51, 46, 49, 52, 49, 53, 57, 50}                       | float64(-3.141592) |
| []byte{49}                                                       | float64(1)         |
| []byte{45, 49}                                                   | float64(-1)        |
| nil                                                              | float64(0)         |
| int(-1)                                                          | float64(-1)        |
| int(0)                                                           | float64(0)         |
| int(1)                                                           | float64(1)         |
| int8(-1)                                                         | float64(-1)        |
| int8(0)                                                          | float64(0)         |
| int8(1)                                                          | float64(1)         |
| int16(-1)                                                        | float64(-1)        |
| int16(0)                                                         | float64(0)         |
| int16(1)                                                         | float64(1)         |
| int32(-1)                                                        | float64(-1)        |
| int32(0)                                                         | float64(0)         |
| int32(1)                                                         | float64(1)         |
| int64(-1)                                                        | float64(-1)        |
| int64(0)                                                         | float64(0)         |
| int64(1)                                                         | float64(1)         |
| uint(0)                                                          | float64(0)         |
| uint(1)                                                          | float64(1)         |
| uint8(0)                                                         | float64(0)         |
| uint8(1)                                                         | float64(1)         |
| uint16(0)                                                        | float64(0)         |
| uint16(1)                                                        | float64(1)         |
| uint32(0)                                                        | float64(0)         |
| uint32(1)                                                        | float64(1)         |
| uint64(0)                                                        | float64(0)         |
| uint64(1)                                                        | float64(1)         |
| float32(-1)                                                      | float64(-1)        |
| float32(0)                                                       | float64(0)         |
| float32(1)                                                       | float64(1)         |
| float32(-3.141592)                                               | float64(-3.141592) |
| float32(3.141592)                                                | float64(3.141592)  |
| bool(false)                                                      | float64(0)         |
| bool(true)                                                       | float64(1)         |
| -                                                                | float64(0)         |

---

## _as.ByteList - 转换为字节数组

### 函数签名
```go
func ByteList(value interface{}) []byte
```

### 功能说明
将任意类型的值转换为字节数组。

### 转换规则

- ① **[]byte 类型**：保持不变
- ② **其他类型**：先通过 `_as.String` 转换成字符串字面量，然后通过 `[]byte` 转化成字节数组

### 转换示例

| 转换前类型                                                            | 转换后                                                              |
|------------------------------------------------------------------|------------------------------------------------------------------|
| []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33} | []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33} |
| []byte{}                                                         | []byte{}                                                         |
| []byte{51, 46, 49, 52, 49, 53, 57, 50}                           | []byte{51, 46, 49, 52, 49, 53, 57, 50}                           |
| []byte{45, 51, 46, 49, 52, 49, 53, 57, 50}                       | []byte{45, 51, 46, 49, 52, 49, 53, 57, 50}                       |
| []byte{49}                                                       | []byte{49}                                                       |
| []byte{45, 49}                                                   | []byte{45, 49}                                                   |
| -                                                                | []byte(_as.String(v))                                            |

---

## _as.Int64 - 转换为 int64

### 函数签名
```go
func Int64(value interface{}) int64
```

### 功能说明
将任意类型的值转换为 64 位有符号整数。

### 转换规则

- 先通过 `_as.Float64` 转换成 `float64`，然后通过 `int64` 转换

### 实现逻辑

| 转换前类型 | 转换后                   |
|-------|-----------------------|
| -     | int64(_as.Float64(v)) |

---

## _as.Int - 转换为 int

### 函数签名
```go
func Int(value interface{}) int
```

### 功能说明
将任意类型的值转换为有符号整数。

### 转换规则

- 先通过 `_as.Float64` 转换成 `float64`，然后通过 `int` 转换

### 实现逻辑

| 转换前类型 | 转换后                 |
|-------|---------------------|
| -     | int(_as.Float64(v)) |

---

## _as.Uint64 - 转换为 uint64

### 函数签名
```go
func Uint64(value interface{}) uint64
```

### 功能说明
将任意类型的值转换为 64 位无符号整数。负数会自动转换为对应的正数。

### 转换规则

- ① 先通过 `_as.Int64` 转换成 `int64`，得到中间值 `m`
- ② 如果 `m >= 0`，返回 `uint64(m)`，否则返回 `uint64(m*-1)`

### 实现逻辑

| 转换前类型 | 转换后                                                                |
|-------|---------------------------------------------------------------------|
| -     | m := _as.Int64(v); if m >= 0 { uint64(m) } else { uint64(m * -1) } |

---

## _as.Uint - 转换为 uint

### 函数签名
```go
func Uint(value interface{}) uint
```

### 功能说明
将任意类型的值转换为无符号整数。负数会自动转换为对应的正数。

### 转换规则

- ① 先通过 `_as.Int` 转换成 `int`，得到中间值 `m`
- ② 如果 `m >= 0`，返回 `uint(m)`，否则返回 `uint(m*-1)`

### 实现逻辑

| 转换前类型 | 转换后                                                          |
|-------|-----------------------------------------------------------------|
| -     | m := _as.Int(v); if m >= 0 { uint(m) } else { uint(m * -1) } |

---

## 📝 使用建议

1. **零值安全**：所有转换函数失败时返回对应类型的零值，不会 panic
2. **精度注意**：通过 Float64 中转的整数转换可能会有精度损失（小数部分会被截断）
3. **Bool 特性**：字符串转 bool 时，只有空字符串返回 `false`，其他任何字符串（包括 `"false"`）都返回 `true`
4. **负数处理**：转换为无符号整数时，负数会自动取绝对值

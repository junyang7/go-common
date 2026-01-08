# _context - HTTP 请求上下文

专业的 HTTP 请求参数解析和处理工具，支持 GET、POST、文件上传、JSON、Cookie、Header 等。

---

## 🎯 核心功能

- ✅ **GET 参数解析**：URL Query 参数
- ✅ **POST 参数解析**：表单、JSON、文件上传
- ✅ **参数合并**：GET + POST 合并（POST 优先）
- ✅ **Cookie 解析**：获取 Cookie 值
- ✅ **Header 解析**：获取请求头
- ✅ **Server 信息**：请求元数据、客户端 IP
- ✅ **文件上传**：Multipart 文件处理
- ✅ **数据绑定**：自动绑定到结构体

---

## 💡 使用示例

### 创建上下文

```go
import "_context"

func Handler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 使用 ctx 处理请求...
}
```

### GET 参数

```go
// GET /api?name=alice&age=30

// 获取单个参数
name := ctx.Get("name").String().Value()     // "alice"
age := ctx.Get("age").Int64().Value()        // 30

// 使用默认值
city := ctx.Get("city").Default("beijing").String().Value()

// 获取所有 GET 参数
all := ctx.GetAll()  // map[string]interface{}
```

### POST 参数

```go
// POST /api
// Content-Type: application/x-www-form-urlencoded
// Body: username=bob&password=secret

// 获取单个参数
username := ctx.Post("username").String().Value()  // "bob"
password := ctx.Post("password").String().Value()  // "secret"

// 获取所有 POST 参数
all := ctx.PostAll()  // map[string]interface{}
```

### POST JSON

```go
// POST /api
// Content-Type: application/json
// Body: {"name":"charlie","age":35}

// 获取参数
name := ctx.Post("name").String().Value()  // "charlie"
age := ctx.Post("age").Int64().Value()     // 35

// 获取原始 Body
body := ctx.Body()  // []byte
```

### Request 合并参数（POST 优先）

```go
// POST /api?id=123&source=web
// Body: action=update&value=100

// GET 参数
id := ctx.Get("id").String().Value()           // "123"
source := ctx.Get("source").String().Value()   // "web"

// POST 参数
action := ctx.Post("action").String().Value()  // "update"
value := ctx.Post("value").String().Value()    // "100"

// 合并参数（POST 优先）
id := ctx.Request("id").String().Value()       // "123"（来自 GET）
action := ctx.Request("action").String().Value() // "update"（来自 POST）

// 如果 GET 和 POST 都有同名参数，POST 优先
// GET: name=alice, POST: name=bob
name := ctx.Request("name").String().Value()  // "bob"（POST 优先）
```

### Cookie

```go
// 获取 Cookie
session := ctx.Cookie("session").String().Value()
token := ctx.Cookie("token").String().Value()

// 获取所有 Cookie
all := ctx.CookieAll()  // map[string]string
```

### Header

```go
// 获取 Header（大小写不敏感）
auth := ctx.Header("Authorization").String().Value()
contentType := ctx.Header("content-type").String().Value()

// 获取所有 Header
all := ctx.HeaderAll()  // map[string]string（key 为小写）
```

### Server 信息

```go
// 请求信息
method := ctx.Server("method").String().Value()       // "GET" / "POST"
path := ctx.Server("path").String().Value()           // "/api/users"
host := ctx.Server("host").String().Value()           // "example.com"
protocol := ctx.Server("protocol").String().Value()   // "HTTP/1.1"
scheme := ctx.Server("scheme").String().Value()       // "http" / "https"
url := ctx.Server("url").String().Value()             // "http://example.com/api"

// 客户端信息
clientIP := ctx.Server("client-ip").String().Value()  // 真实 IP
userAgent := ctx.Server("user-agent").String().Value()
referer := ctx.Server("referer").String().Value()

// Content 信息
contentType := ctx.Server("content-type").String().Value()
contentLength := ctx.Server("content-length").String().Value()

// 获取所有 Server 信息
all := ctx.ServerAll()  // map[string]string
```

### 文件上传

```go
// POST /upload
// Content-Type: multipart/form-data

// 方式1：获取第一个文件（推荐）⭐
file := ctx.File("avatar").File()
if file != nil {
    filename := file.Filename  // "avatar.jpg"
    size := file.Size          // 文件大小
    
    // 打开文件
    f, err := file.Open()
    defer f.Close()
    // 读取文件...
}

// 方式2：获取所有文件（多文件上传）
files := ctx.File("avatar").FileList()
if files != nil && len(files) > 0 {
    for _, file := range files {
        // 处理每个文件...
    }
}

// 方式3：必填验证 + 获取
file := ctx.File("avatar").Required().File()
// 如果文件不存在会抛出异常

// 获取所有上传文件
all := ctx.FileAll()  // map[string][]*multipart.FileHeader
```

### 数据绑定

```go
type User struct {
    Name  string `json:"name"`
    Age   int64  `json:"age"`
    Email string `json:"email"`
}

// 自动绑定（智能优先级）⭐
// 优先级：GET < POST表单 < POST JSON
// 后面的会覆盖前面的同名字段

// 场景1：纯 JSON
// POST /api
// Content-Type: application/json
// Body: {"name":"alice","age":30,"email":"alice@example.com"}
var user User
ctx.Bind(&user)
// 从 JSON Body 解析

// 场景2：表单数据
// POST /api
// Content-Type: application/x-www-form-urlencoded
// Body: name=bob&age=25
var user User
ctx.Bind(&user)
// 从表单参数绑定

// 场景3：GET + POST 混合（自动合并）⭐
// POST /api?city=beijing&score=100
// Body: name=alice&age=30
var user User
ctx.Bind(&user)
// 自动合并：GET 的 city/score + POST 的 name/age

// 场景4：GET + POST JSON（自动合并）⭐
// POST /api?city=beijing&score=100
// Content-Type: application/json
// Body: {"name":"alice","age":30}
var user User
ctx.Bind(&user)
// 优先级：city/score 来自 GET，name/age 来自 JSON（JSON 优先级最高）

// 手动指定来源
ctx.BindGet(&user)   // 只从 GET 绑定
ctx.BindPost(&user)  // 只从 POST 绑定（自动识别 JSON/Form）⭐
```

---

## 🎓 三个 Bind 方法对比

### Bind - 全自动（推荐）⭐

```go
ctx.Bind(&user)
```

**特点：**
- ✅ 自动合并 GET + POST 参数
- ✅ 优先级：GET < POST表单 < POST JSON
- ✅ JSON 保留完整结构（嵌套对象）
- ✅ 90% 场景使用

**示例：**
```go
// POST /api?source=web
// Body: {"name":"alice","age":30}
ctx.Bind(&user)
// user.source = "web" (GET)
// user.name = "alice" (JSON)
// user.age = 30 (JSON)
```

### BindPost - 只要 POST（自动识别）⭐

```go
ctx.BindPost(&user)
```

**特点：**
- ✅ 只绑定 POST 数据，忽略 GET
- ✅ 自动识别 Content-Type
  - JSON → 从 Body 解析（保留嵌套）
  - Form → 从 post map 映射
- ✅ 不受 URL 参数影响

**示例：**
```go
// POST /api?source=web  ← 忽略
// Body: {"name":"alice"}
ctx.BindPost(&user)
// user.source = "" (忽略 GET)
// user.name = "alice" (POST)
```

### BindGet - 只要 GET

```go
ctx.BindGet(&query)
```

**特点：**
- ✅ 只绑定 GET 参数（URL Query）
- ✅ 忽略 POST 数据
- ✅ 适合分页、搜索等查询参数

**示例：**
```go
// GET /api?page=2&size=20
ctx.BindGet(&query)
// query.page = 2
// query.size = 20
```

---

## 📦 API 文档

### 参数获取

| 方法 | 说明 | 返回 |
|------|------|------|
| `Get(key)` | 获取 GET 参数 | `*Parameter` |
| `Post(key)` | 获取 POST 参数 | `*Parameter` |
| `Request(key)` | 获取合并参数（POST 优先） | `*Parameter` |
| `Cookie(key)` | 获取 Cookie | `*Parameter` |
| `Header(key)` | 获取 Header（小写） | `*Parameter` |
| `Server(key)` | 获取 Server 信息 | `*Parameter` |
| `File(key)` | 获取上传文件 | `*Parameter` ⭐ |
| `Body()` | 获取原始 Body | `[]byte` |

### 批量获取

| 方法 | 说明 | 返回 |
|------|------|------|
| `GetAll()` | 获取所有 GET 参数 | `map[string]interface{}` |
| `PostAll()` | 获取所有 POST 参数 | `map[string]interface{}` |
| `RequestAll()` | 获取所有合并参数 | `map[string]interface{}` |
| `CookieAll()` | 获取所有 Cookie | `map[string]string` |
| `HeaderAll()` | 获取所有 Header | `map[string]string` |
| `ServerAll()` | 获取所有 Server 信息 | `map[string]string` |
| `FileAll()` | 获取所有上传文件 | `map[string][]*FileHeader` |

### 数据绑定

| 方法 | 说明 |
|------|------|
| `Bind(v)` | 自动绑定（GET + POST，智能优先级） |
| `BindGet(v)` | 只从 GET 参数绑定 |
| `BindPost(v)` | 只从 POST 绑定（自动识别 JSON/Form）⭐ |

### 响应方法

| 方法 | 说明 |
|------|------|
| `SetHeader(k, v)` | 设置响应头 |
| `SetCookie(cookie)` | 设置 Cookie |
| `JSON(data)` | 返回 JSON 响应 |
| `REDIRECT(uri)` | 重定向 |

---

## 🔑 参数存储规则

### 存储位置

```
get:     GET 参数（URL Query）
post:    POST 参数（表单/JSON）
request: 合并参数（GET + POST，POST 优先）
cookie:  Cookie 值
header:  Header 值（小写 key）
server:  Server 信息（小写 key）
file:    上传文件
```

### 优先级规则

```
Request 合并规则：
1. 先添加所有 GET 参数
2. POST 参数覆盖同名 GET 参数
3. POST 优先级最高
```

**示例：**
```
GET:  name=alice, id=123
POST: name=bob, age=30

Request: 
  id=123    (来自 GET)
  age=30    (来自 POST)
  name=bob  (POST 覆盖 GET)
```

---

## ⭐ 自动绑定优先级

`ctx.Bind()` 会自动按优先级填充结构体：

### 优先级顺序

```
GET 参数 < POST 表单 < POST JSON
   ↓         ↓           ↓
  低      →  中  →      高
```

### 绑定流程

```
1. 先从 GET 参数填充
2. POST 表单参数覆盖同名字段
3. POST JSON 覆盖所有同名字段
```

### 实际示例

```go
type User struct {
    Name  string `json:"name"`
    Age   int64  `json:"age"`
    City  string `json:"city"`
}

// POST /api?city=beijing&age=18
// Content-Type: application/json
// Body: {"name":"alice","age":30}

var user User
ctx.Bind(&user)

// 结果：
user.Name  // "alice"   (来自 JSON)
user.Age   // 30        (来自 JSON，覆盖 GET 的 18)
user.City  // "beijing" (来自 GET，JSON 没有这个字段)
```

### 为什么这样设计？

```
✅ 最符合直觉：JSON Body 应该是主要数据
✅ 灵活性好：URL 参数可以提供额外字段
✅ 全自动：不需要判断 Content-Type

例如：
POST /api/users?source=web&version=2
Body: {"name":"alice","email":"alice@example.com"}

绑定后：
name, email 来自 JSON（主要数据）
source, version 来自 URL（元数据）
```

---

## 📋 Server 信息字段（36个）

### 基础信息（9个）

| 字段 | 说明 | 示例 |
|------|------|------|
| `method` | HTTP 方法 | GET, POST, PUT, DELETE |
| `path` | 请求路径 | /api/users |
| `query` | 查询字符串 | id=123&name=alice |
| `host` | 主机名 | example.com |
| `protocol` | 协议版本 | HTTP/1.1, HTTP/2.0 |
| `scheme` | 协议类型 | http, https |
| `url` | 完整 URL | http://example.com/api |
| `remote-addr` | 远程地址 | 192.168.1.100:12345 |
| `request-uri` | 请求 URI | /api?id=123 |

### Content 相关（3个）

| 字段 | 说明 |
|------|------|
| `content-type` | 内容类型 |
| `content-length` | 内容长度 |
| `content-encoding` | 内容编码 |

### Accept 相关（4个）

| 字段 | 说明 |
|------|------|
| `accept` | 接受类型 |
| `accept-encoding` | 接受编码 |
| `accept-language` | 接受语言 |
| `accept-charset` | 接受字符集 |

### 客户端信息（3个）

| 字段 | 说明 |
|------|------|
| `client-ip` | 真实 IP（智能获取）⭐ |
| `user-agent` | 用户代理 |
| `referer` | 来源页面 |

### 跨域 CORS（3个）

| 字段 | 说明 |
|------|------|
| `origin` | 源站 |
| `access-control-request-method` | 预检请求方法 |
| `access-control-request-headers` | 预检请求头 |

### 认证相关（1个）

| 字段 | 说明 |
|------|------|
| `authorization` | 认证令牌 |

### AJAX 标识（1个）

| 字段 | 说明 |
|------|------|
| `x-requested-with` | AJAX 请求标识 |

### 代理相关（4个）

| 字段 | 说明 |
|------|------|
| `x-forwarded-for` | 代理 IP 链 |
| `x-forwarded-host` | 原始 Host |
| `x-forwarded-proto` | 原始协议 |
| `x-real-ip` | 真实 IP |

### 缓存相关（4个）

| 字段 | 说明 |
|------|------|
| `cache-control` | 缓存控制 |
| `if-modified-since` | 条件请求（时间） |
| `if-none-match` | 条件请求（ETag） |
| `if-match` | 条件匹配 |

### 连接相关（2个）

| 字段 | 说明 |
|------|------|
| `connection` | 连接类型 |
| `upgrade` | 协议升级 |

### 范围请求（1个）

| 字段 | 说明 |
|------|------|
| `range` | 范围请求 |

### 其他（2个）

| 字段 | 说明 |
|------|------|
| `dnt` | Do Not Track |
| `upgrade-insecure-requests` | HTTPS 升级 |

---

## 🔒 客户端 IP 获取策略

按优先级依次尝试：

```
1. X-Forwarded-For（取第一个）
2. X-Real-IP
3. RemoteAddr（去除端口）
```

---

## 📖 完整示例

### 示例 1：用户登录

```go
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 获取登录参数
    username := ctx.Post("username").String().Value()
    password := ctx.Post("password").String().Value()
    
    // 验证...
    if valid {
        ctx.JSON(map[string]string{"token": "..."})
    } else {
        ctx.JSON(errorResponse)
    }
}
```

### 示例 2：分页查询

```go
func ListHandler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // GET /api/users?page=2&size=20&sort=name
    page := ctx.Get("page").Default(1).Int64().Value()
    size := ctx.Get("size").Default(10).Int64().Value()
    sort := ctx.Get("sort").Default("id").String().Value()
    
    // 查询数据...
    ctx.JSON(data)
}
```

### 示例 3：数据绑定

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int64  `json:"age"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 自动绑定（支持 JSON 和表单）
    var req CreateUserRequest
    ctx.Bind(&req)
    
    // 使用绑定后的数据
    user := createUser(req.Name, req.Email, req.Age)
    ctx.JSON(user)
}
```

### 示例 4：文件上传

```go
func UploadHandler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 获取上传的文件
    files := ctx.File("avatar")
    if files == nil || len(files) == 0 {
        ctx.JSON(map[string]string{"error": "no file uploaded"})
        return
    }
    
    file := files[0]
    filename := file.Filename
    
    // 打开并保存文件
    f, _ := file.Open()
    defer f.Close()
    // 保存文件...
    
    ctx.JSON(map[string]string{"filename": filename})
}
```

### 示例 5：API 认证

```go
func AuthHandler(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 从 Header 获取 Token
    token := ctx.Header("Authorization").String().Value()
    
    // 验证 Token...
    if !valid {
        ctx.JSON(map[string]string{"error": "unauthorized"})
        return
    }
    
    // 继续处理...
}
```

### 示例 6：访问日志

```go
func LogMiddleware(w http.ResponseWriter, r *http.Request) {
    ctx := _context.New(w, r, false)
    
    // 记录请求信息
    log.Printf(
        "[%s] %s %s from %s",
        ctx.Server("method").String().Value(),
        ctx.Server("path").String().Value(),
        ctx.Server("protocol").String().Value(),
        ctx.Server("client-ip").String().Value(),
    )
}
```

---

## ⚡ 性能

所有测试通过，代码覆盖率 **73.9%**

---

## ⚠️ 注意事项

### 1. Content-Type 支持

当前支持的 Content-Type：
- `application/x-www-form-urlencoded` - 表单
- `multipart/form-data` - 文件上传
- `application/json` - JSON

### 2. POST 和 GET 同时存在

```
POST /api?name=alice
Body: name=bob

ctx.Get("name")     // "alice" (GET)
ctx.Post("name")    // "bob"   (POST)
ctx.Request("name") // "bob"   (POST 优先)
```

### 3. 文件上传限制

默认最大 32MB，可在 `parseMultipartForm` 中调整。

### 4. JSON 解析

仅支持对象类型 `{}` 的 JSON，不支持数组 `[]`。

---

## 🎯 设计亮点

### 1. 清晰的参数分离

```go
ctx.Get()     // 明确来自 GET
ctx.Post()    // 明确来自 POST
ctx.Request() // 明确是合并的
```

### 2. 统一的返回类型

```go
// 所有参数方法都返回 *Parameter
ctx.Get("name")     // *Parameter
ctx.Post("age")     // *Parameter
ctx.Cookie("token") // *Parameter
ctx.Header("auth")  // *Parameter
ctx.Server("ip")    // *Parameter

// 支持链式调用和类型转换
value := ctx.Get("age").Default(18).Int64().Value()
```

### 3. 智能绑定

```go
// 自动识别 Content-Type
ctx.Bind(&user)

// JSON -> 从 Body 解析
// Form -> 从 Request 映射
```

### 4. 小写标准化

```go
// Header 和 Server 的 key 统一转小写
ctx.Header("Authorization")  // 内部存为 "authorization"
ctx.Server("Content-Type")   // 内部存为 "content-type"

// 避免大小写问题
```

---

**License:** MIT  
**Version:** 2.0  
**Status:** Production Ready ✅


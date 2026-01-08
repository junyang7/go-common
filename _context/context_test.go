package _context

import (
	"bytes"
	"github.com/junyang7/go-common/_assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ============================================================
// GET 请求测试
// ============================================================

func TestContext_GET(t *testing.T) {
	// 创建 GET 请求
	req := httptest.NewRequest("GET", "http://example.com/path?name=alice&age=30&active=true", nil)
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 测试 Get 方法
	{
		name := ctx.Get("name").String().Value()
		_assert.Equal(t, "alice", name)
	}

	{
		age := ctx.Get("age").Int64().Value()
		_assert.Equal(t, int64(30), age)
	}

	{
		active := ctx.Get("active").Bool().Value()
		_assert.Equal(t, true, active)
	}

	// 不存在的参数
	{
		param := ctx.Get("notexist")
		_assert.Nil(t, param.Value())
	}

	// 测试 GetAll
	{
		all := ctx.GetAll()
		_assert.Equal(t, "alice", all["name"])
		_assert.Equal(t, "30", all["age"])
	}
}

// ============================================================
// POST 表单请求测试
// ============================================================

func TestContext_POST_Form(t *testing.T) {
	// 创建 POST 表单请求
	body := "username=bob&password=secret&age=25"
	req := httptest.NewRequest("POST", "http://example.com/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 测试 Post 方法
	{
		username := ctx.Post("username").String().Value()
		_assert.Equal(t, "bob", username)
	}

	{
		password := ctx.Post("password").String().Value()
		_assert.Equal(t, "secret", password)
	}

	{
		age := ctx.Post("age").Int64().Value()
		_assert.Equal(t, int64(25), age)
	}

	// 测试 PostAll
	{
		all := ctx.PostAll()
		_assert.Equal(t, "bob", all["username"])
	}
}

// ============================================================
// POST JSON 请求测试
// ============================================================

func TestContext_POST_JSON(t *testing.T) {
	// 创建 POST JSON 请求
	jsonBody := `{"name":"charlie","age":35,"active":true,"score":98.5}`
	req := httptest.NewRequest("POST", "http://example.com/api", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 测试 Post 方法
	{
		name := ctx.Post("name").String().Value()
		_assert.Equal(t, "charlie", name)
	}

	{
		age := ctx.Post("age").Int64().Value()
		_assert.Equal(t, int64(35), age)
	}

	{
		active := ctx.Post("active").Bool().Value()
		_assert.Equal(t, true, active)
	}

	{
		score := ctx.Post("score").Float64().Value()
		_assert.Equal(t, 98.5, score)
	}

	// 测试 Body
	{
		body := ctx.Body()
		_assert.Equal(t, jsonBody, string(body))
	}
}

// ============================================================
// Request 合并参数测试
// ============================================================

func TestContext_Request_Merge(t *testing.T) {
	// 创建带 GET 和 POST 的请求
	body := "name=bob&age=25"
	req := httptest.NewRequest("POST", "http://example.com/api?name=alice&city=beijing", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// GET 参数
	{
		name := ctx.Get("name").String().Value()
		_assert.Equal(t, "alice", name)

		city := ctx.Get("city").String().Value()
		_assert.Equal(t, "beijing", city)
	}

	// POST 参数
	{
		name := ctx.Post("name").String().Value()
		_assert.Equal(t, "bob", name)

		age := ctx.Post("age").String().Value()
		_assert.Equal(t, "25", age)
	}

	// Request 合并（POST 优先）
	{
		name := ctx.Request("name").String().Value()
		_assert.Equal(t, "bob", name) // POST 覆盖 GET

		city := ctx.Request("city").String().Value()
		_assert.Equal(t, "beijing", city) // 只有 GET

		age := ctx.Request("age").String().Value()
		_assert.Equal(t, "25", age) // 只有 POST
	}
}

// ============================================================
// Cookie 测试
// ============================================================

func TestContext_Cookie(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/", nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: "abc123"})
	req.AddCookie(&http.Cookie{Name: "token", Value: "xyz789"})
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 测试 Cookie 方法
	{
		session := ctx.Cookie("session").String().Value()
		_assert.Equal(t, "abc123", session)
	}

	{
		token := ctx.Cookie("token").String().Value()
		_assert.Equal(t, "xyz789", token)
	}

	// 不存在的 Cookie
	{
		param := ctx.Cookie("notexist")
		_assert.Nil(t, param.Value())
	}

	// 测试 CookieAll
	{
		all := ctx.CookieAll()
		_assert.Equal(t, "abc123", all["session"])
		_assert.Equal(t, "xyz789", all["token"])
	}
}

// ============================================================
// Header 测试
// ============================================================

func TestContext_Header(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/", nil)
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Set("X-Custom-Header", "custom-value")
	req.Header.Set("User-Agent", "TestClient/1.0")
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 测试 Header 方法（大小写不敏感）
	{
		auth := ctx.Header("authorization").String().Value()
		_assert.Equal(t, "Bearer token123", auth)
	}

	{
		auth := ctx.Header("Authorization").String().Value() // 大写
		_assert.Equal(t, "Bearer token123", auth)
	}

	{
		custom := ctx.Header("x-custom-header").String().Value()
		_assert.Equal(t, "custom-value", custom)
	}

	// 测试 HeaderAll
	{
		all := ctx.HeaderAll()
		_assert.Equal(t, "Bearer token123", all["authorization"])
		_assert.Equal(t, "custom-value", all["x-custom-header"])
	}
}

// ============================================================
// Server 信息测试
// ============================================================

func TestContext_Server(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/api/users?id=123", nil)
	req.Header.Set("User-Agent", "TestClient/1.0")
	req.Header.Set("Referer", "http://example.com/home")
	req.RemoteAddr = "192.168.1.100:12345"
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 测试基础信息
	{
		method := ctx.Server("method").String().Value()
		_assert.Equal(t, "POST", method)
	}

	{
		path := ctx.Server("path").String().Value()
		_assert.Equal(t, "/api/users", path)
	}

	{
		host := ctx.Server("host").String().Value()
		_assert.Equal(t, "example.com", host)
	}

	{
		protocol := ctx.Server("protocol").String().Value()
		_assert.Contains(t, protocol, "HTTP")
	}

	{
		scheme := ctx.Server("scheme").String().Value()
		_assert.Equal(t, "http", scheme)
	}

	// 测试 Header 信息
	{
		ua := ctx.Server("user-agent").String().Value()
		_assert.Equal(t, "TestClient/1.0", ua)
	}

	{
		referer := ctx.Server("referer").String().Value()
		_assert.Equal(t, "http://example.com/home", referer)
	}

	// 测试客户端 IP
	{
		ip := ctx.Server("client-ip").String().Value()
		_assert.Equal(t, "192.168.1.100", ip)
	}
}

// ============================================================
// 文件上传测试
// ============================================================

func TestContext_File(t *testing.T) {
	// 创建 multipart 表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件字段
	fileWriter, _ := writer.CreateFormFile("avatar", "avatar.jpg")
	fileWriter.Write([]byte("fake image data"))

	// 添加普通字段
	writer.WriteField("username", "alice")
	writer.WriteField("age", "30")

	writer.Close()

	// 创建请求
	req := httptest.NewRequest("POST", "http://example.com/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 测试 File()（返回第一个文件）
	{
		file := ctx.File("avatar").File()
		_assert.NotNil(t, file)
		_assert.Equal(t, "avatar.jpg", file.Filename)
	}

	// 测试 FileList()（返回所有文件）
	{
		files := ctx.File("avatar").FileList()
		_assert.NotNil(t, files)
		_assert.Len(t, files, 1)
		_assert.Equal(t, "avatar.jpg", files[0].Filename)
	}

	// 测试表单字段
	{
		username := ctx.Post("username").String().Value()
		_assert.Equal(t, "alice", username)
	}

	{
		age := ctx.Post("age").Int64().Value()
		_assert.Equal(t, int64(30), age)
	}

	// 不存在的文件
	{
		file := ctx.File("notexist").File()
		_assert.Nil(t, file)
		
		files := ctx.File("notexist").FileList()
		_assert.Nil(t, files)
	}
}

// ============================================================
// Bind 绑定测试
// ============================================================

func TestContext_Bind_JSON(t *testing.T) {
	type User struct {
		Name   string  `json:"name"`
		Age    int64   `json:"age"`
		Active bool    `json:"active"`
		Score  float64 `json:"score"`
	}

	jsonBody := `{"name":"alice","age":30,"active":true,"score":95.5}`
	req := httptest.NewRequest("POST", "http://example.com/api", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 绑定到结构体
	var user User
	ctx.Bind(&user)

	_assert.Equal(t, "alice", user.Name)
	_assert.Equal(t, int64(30), user.Age)
	_assert.Equal(t, true, user.Active)
	_assert.Equal(t, 95.5, user.Score)
}

func TestContext_Bind_Form(t *testing.T) {
	type LoginForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Remember bool   `json:"remember"`
	}

	body := "username=bob&password=secret&remember=true"
	req := httptest.NewRequest("POST", "http://example.com/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 绑定到结构体
	var form LoginForm
	ctx.Bind(&form)

	_assert.Equal(t, "bob", form.Username)
	_assert.Equal(t, "secret", form.Password)
	_assert.Equal(t, true, form.Remember)
}

func TestContext_BindGet(t *testing.T) {
	type Query struct {
		Page  int64  `json:"page"`
		Size  int64  `json:"size"`
		Sort  string `json:"sort"`
	}

	req := httptest.NewRequest("GET", "http://example.com/list?page=2&size=20&sort=name", nil)
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	var query Query
	ctx.BindGet(&query)

	_assert.Equal(t, int64(2), query.Page)
	_assert.Equal(t, int64(20), query.Size)
	_assert.Equal(t, "name", query.Sort)
}

func TestContext_BindPost(t *testing.T) {
	// 测试表单
	{
		type CreateUser struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		body := "name=charlie&email=charlie@example.com"
		req := httptest.NewRequest("POST", "http://example.com/users?extra=value", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		var user CreateUser
		ctx.BindPost(&user)

		_assert.Equal(t, "charlie", user.Name)
		_assert.Equal(t, "charlie@example.com", user.Email)
	}

	// 测试 JSON
	{
		type CreateUser struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		jsonBody := `{"name":"alice","email":"alice@example.com"}`
		req := httptest.NewRequest("POST", "http://example.com/users", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		var user CreateUser
		ctx.BindPost(&user)

		_assert.Equal(t, "alice", user.Name)
		_assert.Equal(t, "alice@example.com", user.Email)
	}

	// 测试嵌套 JSON（BindPost 应该支持）⭐
	{
		type Config struct {
			Theme string `json:"theme"`
			Lang  string `json:"lang"`
		}
		type UserWithConfig struct {
			Name   string  `json:"name"`
			Config Config  `json:"config"`
		}

		jsonBody := `{"name":"bob","config":{"theme":"dark","lang":"zh"}}`
		req := httptest.NewRequest("POST", "http://example.com/users", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		var user UserWithConfig
		ctx.BindPost(&user)

		_assert.Equal(t, "bob", user.Name)
		_assert.Equal(t, "dark", user.Config.Theme)
		_assert.Equal(t, "zh", user.Config.Lang)
	}
}

func TestContext_Bind_Priority(t *testing.T) {
	// 测试自动优先级：GET < POST表单 < POST JSON
	type Data struct {
		Name  string `json:"name"`
		Age   int64  `json:"age"`
		City  string `json:"city"`
		Score int64  `json:"score"`
	}

	// 场景1：GET + POST 表单
	{
		body := "name=post-form&age=25"
		req := httptest.NewRequest("POST", "http://example.com/api?name=get-url&city=beijing&score=100", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		var data Data
		ctx.Bind(&data)

		// 优先级验证
		_assert.Equal(t, "post-form", data.Name) // POST 表单覆盖 GET
		_assert.Equal(t, int64(25), data.Age)    // 只在 POST
		_assert.Equal(t, "beijing", data.City)   // 只在 GET
		_assert.Equal(t, int64(100), data.Score) // 只在 GET
	}

	// 场景2：GET + POST JSON
	{
		jsonBody := `{"name":"json-body","age":30}`
		req := httptest.NewRequest("POST", "http://example.com/api?name=get-url&city=shanghai&score=200", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		var data Data
		ctx.Bind(&data)

		// 优先级验证
		_assert.Equal(t, "json-body", data.Name) // JSON 覆盖 GET
		_assert.Equal(t, int64(30), data.Age)    // 只在 JSON
		_assert.Equal(t, "shanghai", data.City)  // 只在 GET（JSON 没覆盖）
		_assert.Equal(t, int64(200), data.Score) // 只在 GET（JSON 没覆盖）
	}

	// 场景3：只有 GET
	{
		req := httptest.NewRequest("GET", "http://example.com/api?name=get-only&age=40", nil)
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		var data Data
		ctx.Bind(&data)

		_assert.Equal(t, "get-only", data.Name)
		_assert.Equal(t, int64(40), data.Age)
	}
}

// ============================================================
// 复杂场景测试
// ============================================================

func TestContext_Mixed_GET_POST(t *testing.T) {
	// POST 请求，同时有 URL 参数
	body := "action=update&value=100"
	req := httptest.NewRequest("POST", "http://example.com/api?id=123&source=web", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// GET 参数在 get 中
	{
		id := ctx.Get("id").String().Value()
		_assert.Equal(t, "123", id)

		source := ctx.Get("source").String().Value()
		_assert.Equal(t, "web", source)
	}

	// POST 参数在 post 中
	{
		action := ctx.Post("action").String().Value()
		_assert.Equal(t, "update", action)

		value := ctx.Post("value").String().Value()
		_assert.Equal(t, "100", value)
	}

	// Request 合并（POST 优先）
	{
		id := ctx.Request("id").String().Value()
		_assert.Equal(t, "123", id) // 只在 GET

		action := ctx.Request("action").String().Value()
		_assert.Equal(t, "update", action) // 只在 POST

		// 如果同名，POST 优先
		all := ctx.RequestAll()
		_assert.Equal(t, 4, len(all)) // id, source, action, value
	}
}

func TestContext_POST_Priority(t *testing.T) {
	// GET 和 POST 都有同名参数，POST 应该优先
	body := "name=post-value"
	req := httptest.NewRequest("POST", "http://example.com/api?name=get-value", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// GET 中是 get-value
	{
		name := ctx.Get("name").String().Value()
		_assert.Equal(t, "get-value", name)
	}

	// POST 中是 post-value
	{
		name := ctx.Post("name").String().Value()
		_assert.Equal(t, "post-value", name)
	}

	// Request 中应该是 POST 的值（优先）
	{
		name := ctx.Request("name").String().Value()
		_assert.Equal(t, "post-value", name)
	}
}

// ============================================================
// 默认值测试
// ============================================================

func TestContext_Default(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/?name=alice", nil)
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	// 存在的参数，不使用默认值
	{
		name := ctx.Get("name").String().Value()
		_assert.Equal(t, "alice", name)
	}

	// 不存在的参数，使用默认值
	{
		age := ctx.Get("age").Default(18).Int64().Value()
		_assert.Equal(t, int64(18), age)
	}

	{
		city := ctx.Request("city").Default("beijing").String().Value()
		_assert.Equal(t, "beijing", city)
	}
}

// ============================================================
// Client IP 测试
// ============================================================

func TestContext_ClientIP(t *testing.T) {
	// 测试 X-Forwarded-For
	{
		req := httptest.NewRequest("GET", "http://example.com/", nil)
		req.Header.Set("X-Forwarded-For", "203.0.113.1, 198.51.100.1")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		ip := ctx.Server("client-ip").String().Value()
		_assert.Equal(t, "203.0.113.1", ip) // 取第一个
	}

	// 测试 X-Real-IP
	{
		req := httptest.NewRequest("GET", "http://example.com/", nil)
		req.Header.Set("X-Real-IP", "203.0.113.2")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		ip := ctx.Server("client-ip").String().Value()
		_assert.Equal(t, "203.0.113.2", ip)
	}

	// 测试 RemoteAddr
	{
		req := httptest.NewRequest("GET", "http://example.com/", nil)
		req.RemoteAddr = "192.168.1.100:54321"
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		ip := ctx.Server("client-ip").String().Value()
		_assert.Equal(t, "192.168.1.100", ip) // 去除端口
	}
}

// ============================================================
// 边界情况测试
// ============================================================

func TestContext_EdgeCases(t *testing.T) {
	// 空 GET 请求
	{
		req := httptest.NewRequest("GET", "http://example.com/", nil)
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		param := ctx.Get("any")
		_assert.Nil(t, param.Value())
	}

	// 空 POST 请求
	{
		req := httptest.NewRequest("POST", "http://example.com/", nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		param := ctx.Post("any")
		_assert.Nil(t, param.Value())
	}

	// 空 JSON
	{
		req := httptest.NewRequest("POST", "http://example.com/", strings.NewReader("{}"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		body := ctx.Body()
		_assert.Equal(t, "{}", string(body))
	}

	// 无 Content-Type
	{
		req := httptest.NewRequest("POST", "http://example.com/", strings.NewReader("data"))
		w := httptest.NewRecorder()
		ctx := New(w, req, false)

		// 不应该 panic
		_assert.NotNil(t, ctx)
	}
}

// ============================================================
// 性能测试
// ============================================================

func BenchmarkContext_New_GET(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "http://example.com/?name=test&age=30", nil)
		w := httptest.NewRecorder()
		New(w, req, false)
	}
}

func BenchmarkContext_New_POST_Form(b *testing.B) {
	body := "username=test&password=secret"
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "http://example.com/", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		New(w, req, false)
	}
}

func BenchmarkContext_New_POST_JSON(b *testing.B) {
	jsonBody := `{"name":"test","age":30}`
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "http://example.com/", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		New(w, req, false)
	}
}

func BenchmarkContext_Get(b *testing.B) {
	req := httptest.NewRequest("GET", "http://example.com/?name=test", nil)
	w := httptest.NewRecorder()
	ctx := New(w, req, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.Get("name").String().Value()
	}
}


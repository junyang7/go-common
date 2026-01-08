package _context

import (
	"encoding/json"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_parameter"
	"github.com/junyang7/go-common/_render"
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_unixMilli"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

// Context HTTP 请求上下文
type Context struct {
	// 时间戳
	TimeS int64
	
	// 原始对象（私有）
	w     http.ResponseWriter
	r     *http.Request
	debug bool
	
	// 参数存储（私有，通过方法访问）
	get     map[string]interface{} // GET 参数
	post    map[string]interface{} // POST 参数
	request map[string]interface{} // 合并参数（POST 优先）
	cookie  map[string]string      // Cookie 值
	header  map[string]string      // Header 值
	server  map[string]string      // Server 信息
	file    map[string][]*multipart.FileHeader // 上传文件
	body    []byte                 // 原始 Body
	
	// 自定义存储
	STORE map[string]interface{}
}

// New 创建 HTTP 上下文
func New(w http.ResponseWriter, r *http.Request, debug bool) *Context {
	ctx := &Context{
		TimeS:   _unixMilli.Get(),
		w:       w,
		r:       r,
		debug:   debug,
		STORE:   make(map[string]interface{}),
		get:     make(map[string]interface{}),
		post:    make(map[string]interface{}),
		request: make(map[string]interface{}),
		cookie:  make(map[string]string),
		header:  make(map[string]string),
		server:  make(map[string]string),
		file:    make(map[string][]*multipart.FileHeader),
	}
	
	// 解析请求
	ctx.parseRequest()
	
	return ctx
}

// parseRequest 解析 HTTP 请求
func (c *Context) parseRequest() {
	// 1. 解析 GET 参数（URL Query）
	c.parseGET()
	
	// 2. 解析 Header
	c.parseHeader()
	
	// 3. 解析 Cookie
	c.parseCookie()
	
	// 4. 解析 Server 信息
	c.parseServer()
	
	// 5. 解析 POST 参数（根据 Content-Type）
	c.parsePOST()
	
	// 6. 合并参数（POST 优先）
	c.mergeRequest()
}

// parseGET 解析 GET 参数
func (c *Context) parseGET() {
	for k, v := range c.r.URL.Query() {
		if len(v) > 0 {
			c.get[k] = v[0]
		}
	}
}

// parseHeader 解析 Header
func (c *Context) parseHeader() {
	for k, v := range c.r.Header {
		if len(v) > 0 {
			c.header[strings.ToLower(k)] = v[0]
		}
	}
}

// parseCookie 解析 Cookie
func (c *Context) parseCookie() {
	for _, cookie := range c.r.Cookies() {
		c.cookie[cookie.Name] = cookie.Value
	}
}

// parseServer 解析 Server 信息
func (c *Context) parseServer() {
	// 基础信息
	c.server["method"] = c.r.Method
	c.server["path"] = c.r.URL.Path
	c.server["query"] = c.r.URL.RawQuery
	c.server["host"] = c.r.Host
	c.server["protocol"] = c.r.Proto
	c.server["scheme"] = c.scheme()
	c.server["url"] = c.fullURL()
	
	// 请求信息
	c.server["remote-addr"] = c.r.RemoteAddr
	c.server["request-uri"] = c.r.RequestURI
	
	// Content 相关
	c.server["content-type"] = c.contentType()
	c.server["content-length"] = c.r.Header.Get("Content-Length")
	c.server["content-encoding"] = c.r.Header.Get("Content-Encoding")
	
	// Accept 相关
	c.server["accept"] = c.r.Header.Get("Accept")
	c.server["accept-encoding"] = c.r.Header.Get("Accept-Encoding")
	c.server["accept-language"] = c.r.Header.Get("Accept-Language")
	c.server["accept-charset"] = c.r.Header.Get("Accept-Charset")
	
	// 客户端信息
	c.server["referer"] = c.r.Referer()
	c.server["user-agent"] = c.r.UserAgent()
	c.server["client-ip"] = c.clientIP()
	
	// 跨域相关
	c.server["origin"] = c.r.Header.Get("Origin")
	c.server["access-control-request-method"] = c.r.Header.Get("Access-Control-Request-Method")
	c.server["access-control-request-headers"] = c.r.Header.Get("Access-Control-Request-Headers")
	
	// 认证相关
	c.server["authorization"] = c.r.Header.Get("Authorization")
	
	// AJAX 标识
	c.server["x-requested-with"] = c.r.Header.Get("X-Requested-With")
	
	// 代理相关
	c.server["x-forwarded-for"] = c.r.Header.Get("X-Forwarded-For")
	c.server["x-forwarded-host"] = c.r.Header.Get("X-Forwarded-Host")
	c.server["x-forwarded-proto"] = c.r.Header.Get("X-Forwarded-Proto")
	c.server["x-real-ip"] = c.r.Header.Get("X-Real-IP")
	
	// 缓存相关
	c.server["cache-control"] = c.r.Header.Get("Cache-Control")
	c.server["if-modified-since"] = c.r.Header.Get("If-Modified-Since")
	c.server["if-none-match"] = c.r.Header.Get("If-None-Match")
	c.server["if-match"] = c.r.Header.Get("If-Match")
	
	// 连接相关
	c.server["connection"] = c.r.Header.Get("Connection")
	c.server["upgrade"] = c.r.Header.Get("Upgrade")
	
	// 范围请求
	c.server["range"] = c.r.Header.Get("Range")
	
	// 其他常用
	c.server["dnt"] = c.r.Header.Get("DNT")
	c.server["upgrade-insecure-requests"] = c.r.Header.Get("Upgrade-Insecure-Requests")
}

// parsePOST 解析 POST 参数（根据 Content-Type）
func (c *Context) parsePOST() {
	contentType := c.contentType()
	
	switch contentType {
	case "application/x-www-form-urlencoded":
		c.parseFormData()
	case "multipart/form-data":
		c.parseMultipartForm()
	case "application/json":
		c.parseJSON()
	}
}

// parseFormData 解析表单数据
func (c *Context) parseFormData() {
	if err := c.r.ParseForm(); err != nil {
			_interceptor.Insure(false).Message(err).Do()
		}
	
	for k, v := range c.r.PostForm {
		if len(v) > 0 {
			c.post[k] = v[0]
		}
	}
}

// parseMultipartForm 解析多部分表单（文件上传）
func (c *Context) parseMultipartForm() {
	// 32MB 最大内存
	if err := c.r.ParseMultipartForm(32 << 20); err != nil {
			_interceptor.Insure(false).Message(err).Do()
		}
	
	// 解析表单值
	if c.r.MultipartForm != nil {
		for k, v := range c.r.MultipartForm.Value {
			if len(v) > 0 {
				c.post[k] = v[0]
			}
		}
		
		// 解析上传文件
		for k, v := range c.r.MultipartForm.File {
			if len(v) > 0 {
				c.file[k] = v
			}
		}
	}
}

// parseJSON 解析 JSON Body
func (c *Context) parseJSON() {
	// 读取 Body
	b, err := io.ReadAll(c.r.Body)
	if err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
	defer c.r.Body.Close()
	
	c.body = b
	
	// 解析 JSON
	if len(b) > 0 {
		var data map[string]interface{}
		if err := json.Unmarshal(b, &data); err != nil {
			// 如果不是对象，尝试直接存储
			return
		}
		
		// 存储到 post
		for k, v := range data {
			c.post[k] = v
		}
	}
}

// mergeRequest 合并 GET 和 POST 参数（POST 优先）
func (c *Context) mergeRequest() {
	// 先复制 GET
	for k, v := range c.get {
		c.request[k] = v
	}
	
	// POST 覆盖 GET
	for k, v := range c.post {
		c.request[k] = v
	}
}

// ============================================================
// 参数获取方法
// ============================================================

// Get 获取 GET 参数
func (c *Context) Get(key string) *_parameter.Parameter {
	if v, ok := c.get[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}

// Post 获取 POST 参数
func (c *Context) Post(key string) *_parameter.Parameter {
	if v, ok := c.post[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}

// Request 获取合并参数（POST 优先）
func (c *Context) Request(key string) *_parameter.Parameter {
	if v, ok := c.request[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}

// Cookie 获取 Cookie 值
func (c *Context) Cookie(key string) *_parameter.Parameter {
	if v, ok := c.cookie[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}

// Header 获取 Header 值
func (c *Context) Header(key string) *_parameter.Parameter {
	// 统一转小写
	key = strings.ToLower(key)
	if v, ok := c.header[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}

// Server 获取 Server 信息
func (c *Context) Server(key string) *_parameter.Parameter {
	// 统一转小写
	key = strings.ToLower(key)
	if v, ok := c.server[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}

// File 获取上传文件（返回 Parameter，支持链式调用）
func (c *Context) File(key string) *_parameter.Parameter {
	if files, ok := c.file[key]; ok {
		return _parameter.New(key, files)
	}
	return _parameter.New(key, nil)
}

// Body 获取原始 Body
func (c *Context) Body() []byte {
	return c.body
}

// ============================================================
// 批量获取方法
// ============================================================

// GetAll 获取所有 GET 参数
func (c *Context) GetAll() map[string]interface{} {
	return c.get
}

// PostAll 获取所有 POST 参数
func (c *Context) PostAll() map[string]interface{} {
	return c.post
}

// RequestAll 获取所有合并参数
func (c *Context) RequestAll() map[string]interface{} {
	return c.request
}

// CookieAll 获取所有 Cookie
func (c *Context) CookieAll() map[string]string {
	return c.cookie
}

// HeaderAll 获取所有 Header
func (c *Context) HeaderAll() map[string]string {
	return c.header
}

// ServerAll 获取所有 Server 信息
func (c *Context) ServerAll() map[string]string {
	return c.server
}

// FileAll 获取所有上传文件
func (c *Context) FileAll() map[string][]*multipart.FileHeader {
	return c.file
}

// ============================================================
// 数据绑定
// ============================================================

// Bind 绑定参数到结构体（自动按优先级：GET < POST表单 < POST JSON）
func (c *Context) Bind(v interface{}) {
	// 第一步：从 request 绑定（已包含 GET + POST 表单，POST 优先）
	c.bindFromMap(v, c.request)
	
	// 第二步：如果是 JSON Content-Type，从 JSON Body 解析并覆盖
	if c.contentType() == "application/json" && len(c.body) > 0 {
		_json.Decode(c.body, v)
	}
}

// BindGet 从 GET 参数绑定到结构体
func (c *Context) BindGet(v interface{}) {
	c.bindFromMap(v, c.get)
}

// BindPost 从 POST 参数绑定到结构体（自动根据 Content-Type 选择）
func (c *Context) BindPost(v interface{}) {
	// 如果是 JSON Content-Type，直接从 Body 解析（保留完整结构）
	if c.contentType() == "application/json" && len(c.body) > 0 {
		_json.Decode(c.body, v)
		return
	}
	
	// 否则从 post map 绑定（表单数据）
	c.bindFromMap(v, c.post)
}

// bindFromMap 从 map 绑定到结构体
func (c *Context) bindFromMap(v interface{}, data map[string]interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return
	}
	
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	
	rt := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if !field.CanSet() {
			continue
		}
		
		fieldType := rt.Field(i)
		
		// 获取字段名（优先使用 json tag）
		fieldName := c.getFieldName(fieldType)
		if fieldName == "" || fieldName == "-" {
			continue
		}
		
		// 从数据中获取值
		value, ok := data[fieldName]
		if !ok {
			continue
		}
		
		// 设置字段值
		c.setFieldValue(field, value)
	}
}

// getFieldName 获取字段名（优先 json tag）
func (c *Context) getFieldName(field reflect.StructField) string {
	// 优先使用 json tag
	jsonTag := field.Tag.Get("json")
	if jsonTag != "" {
		// 处理 json tag，如 "name,omitempty"
		if idx := strings.Index(jsonTag, ","); idx != -1 {
			return jsonTag[:idx]
		}
		return jsonTag
	}
	
	// 其次使用 form tag
	formTag := field.Tag.Get("form")
	if formTag != "" {
		return formTag
	}
	
	// 最后使用字段名
	return field.Name
}

// setFieldValue 设置字段值（支持类型转换）
func (c *Context) setFieldValue(field reflect.Value, value interface{}) {
		switch field.Kind() {
		case reflect.String:
		field.SetString(_as.String(value))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			field.SetInt(_as.Int64(value))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			field.SetUint(_as.Uint64(value))
		case reflect.Float32, reflect.Float64:
			field.SetFloat(_as.Float64(value))
		case reflect.Bool:
			field.SetBool(_as.Bool(value))
	case reflect.Slice:
		// 处理切片类型
		c.setSliceValue(field, value)
	case reflect.Interface:
		// 直接设置 interface{}
		field.Set(reflect.ValueOf(value))
	}
}

// setSliceValue 设置切片值
func (c *Context) setSliceValue(field reflect.Value, value interface{}) {
	// 如果 value 本身就是切片，直接使用
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Slice {
		field.Set(rv)
		return
	}
	
	// 否则创建包含单个元素的切片
	slice := reflect.MakeSlice(field.Type(), 1, 1)
	elem := slice.Index(0)
	c.setFieldValue(elem, value)
	field.Set(slice)
}

// ============================================================
// 辅助方法
// ============================================================

// contentType 获取 Content-Type（去除 charset 等参数）
func (c *Context) contentType() string {
	ct := c.r.Header.Get("Content-Type")
	if ct == "" {
		return ""
	}
	
	// 分割并取第一部分，如 "application/json; charset=utf-8" -> "application/json"
	parts := strings.Split(strings.ToLower(ct), ";")
	return strings.TrimSpace(parts[0])
}

// scheme 获取协议（http 或 https）
func (c *Context) scheme() string {
	if c.r.TLS != nil {
		return "https"
	}
	
	// 检查代理 Header
	if scheme := c.r.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	
	return "http"
}

// fullURL 获取完整 URL
func (c *Context) fullURL() string {
	return c.scheme() + "://" + c.r.Host + c.r.RequestURI
}

// clientIP 获取客户端真实 IP
func (c *Context) clientIP() string {
	// 1. 尝试从 X-Forwarded-For 获取
	if ip := c.r.Header.Get("X-Forwarded-For"); ip != "" {
		// X-Forwarded-For 可能包含多个 IP，取第一个
		ips := strings.Split(ip, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}
	
	// 2. 尝试从 X-Real-IP 获取
	if ip := c.r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	
	// 3. 使用 RemoteAddr
	ip := c.r.RemoteAddr
	// 移除端口号
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	return ip
}

// ============================================================
// 响应方法（保持不变）
// ============================================================

// SetHeader 设置响应头
func (c *Context) SetHeader(k string, v string) *Context {
	c.w.Header().Set(k, v)
	return c
}

// SetCookie 设置 Cookie
func (c *Context) SetCookie(cookie *http.Cookie) *Context {
	http.SetCookie(c.w, cookie)
	return c
}

// JSON 返回 JSON 响应
func (c *Context) JSON(data any) {
	res := _response.New()
	if nil != data {
		switch data.(type) {
		case *_response.Response:
			res = data.(*_response.Response)
		default:
			res.Data = data
		}
	}
	if !c.debug {
		res.File = ""
		res.Line = 0
	}
	res.Time = _unixMilli.Get()
	res.Consume = res.Time - c.TimeS
	_render.JSON(c.w, res)
}

// REDIRECT 重定向
func (c *Context) REDIRECT(uri string) {
	c.w.Header().Set("Location", uri)
	c.w.WriteHeader(301)
}

package _context

import (
	"encoding/json"
	"github.com/junyang7/go-common/_bind"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_parameter"
	"github.com/junyang7/go-common/_render"
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_unixMilli"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type Context struct {
	render          *_render.Render                    // 渲染器
	TimeS           int64                              // 时间戳
	w               http.ResponseWriter                // http写对象
	r               *http.Request                      // http读对象
	routerParameter map[string]string                  // 路由参数
	debug           bool                               // 调试开关
	get             map[string]interface{}             // GET参数集合
	post            map[string]interface{}             // POST参数集合
	request         map[string]interface{}             // GET+POST集合，POST覆盖同名GET
	cookie          map[string]string                  // Cookie键值对集合
	header          map[string]string                  // Header键值对集合
	server          map[string]string                  // Server键值对集合
	file            map[string][]*multipart.FileHeader // File键值对集合
	body            []byte                             // Body原始字节
	STORE           map[string]interface{}             // 自定义存储
}

func New(render *_render.Render, timeS int64, w http.ResponseWriter, r *http.Request, routerParameter map[string]string, debug bool) *Context {
	ctx := &Context{
		render:          render,
		TimeS:           timeS,
		w:               w,
		r:               r,
		routerParameter: routerParameter,
		debug:           debug,
		get:             make(map[string]interface{}),
		post:            make(map[string]interface{}),
		request:         make(map[string]interface{}),
		cookie:          make(map[string]string),
		header:          make(map[string]string),
		server:          make(map[string]string),
		file:            make(map[string][]*multipart.FileHeader),
		body:            make([]byte, 0),
		STORE:           make(map[string]interface{}),
	}
	ctx.prepare()
	return ctx
}
func (c *Context) prepare() {
	if nil == c.r || nil == c.w {
		return
	}
	c.prepareCookie()
	c.prepareHeader()
	c.prepareServer()
	c.prepareBody()
	c.prepareGet()
	c.preparePost()
	c.prepareRequest()
}
func (c *Context) prepareCookie() {
	for _, cookie := range c.r.Cookies() {
		c.cookie[cookie.Name] = cookie.Value
	}
}
func (c *Context) prepareHeader() {
	for k, v := range c.r.Header {
		if len(v) > 0 {
			c.header[strings.ToLower(k)] = v[0]
		}
	}
}
func (c *Context) prepareServer() {
	c.server["method"] = c.r.Method
	c.server["path"] = c.r.URL.Path
	c.server["query"] = c.r.URL.RawQuery
	c.server["host"] = c.r.Host
	c.server["protocol"] = c.r.Proto
	c.server["scheme"] = c.helpScheme()
	c.server["url"] = c.helpFullURL()
	c.server["remote-addr"] = c.r.RemoteAddr
	c.server["request-uri"] = c.r.RequestURI
	c.server["content-type"] = c.helpContentType()
	c.server["content-length"] = c.r.Header.Get("Content-Length")
	c.server["content-encoding"] = c.r.Header.Get("Content-Encoding")
	c.server["accept"] = c.r.Header.Get("Accept")
	c.server["accept-encoding"] = c.r.Header.Get("Accept-Encoding")
	c.server["accept-language"] = c.r.Header.Get("Accept-Language")
	c.server["accept-charset"] = c.r.Header.Get("Accept-Charset")
	c.server["referer"] = c.r.Referer()
	c.server["user-agent"] = c.r.UserAgent()
	c.server["client-ip"] = c.helpClientIP()
	c.server["origin"] = c.r.Header.Get("Origin")
	c.server["access-control-request-method"] = c.r.Header.Get("Access-Control-Request-Method")
	c.server["access-control-request-headers"] = c.r.Header.Get("Access-Control-Request-Headers")
	c.server["authorization"] = c.r.Header.Get("Authorization")
	c.server["x-requested-with"] = c.r.Header.Get("X-Requested-With")
	c.server["x-forwarded-for"] = c.r.Header.Get("X-Forwarded-For")
	c.server["x-forwarded-host"] = c.r.Header.Get("X-Forwarded-Host")
	c.server["x-forwarded-proto"] = c.r.Header.Get("X-Forwarded-Proto")
	c.server["x-real-ip"] = c.r.Header.Get("X-Real-IP")
	c.server["cache-control"] = c.r.Header.Get("Cache-Control")
	c.server["if-modified-since"] = c.r.Header.Get("If-Modified-Since")
	c.server["if-none-match"] = c.r.Header.Get("If-None-Match")
	c.server["if-match"] = c.r.Header.Get("If-Match")
	c.server["connection"] = c.r.Header.Get("Connection")
	c.server["upgrade"] = c.r.Header.Get("Upgrade")
	c.server["range"] = c.r.Header.Get("Range")
	c.server["dnt"] = c.r.Header.Get("DNT")
	c.server["upgrade-insecure-requests"] = c.r.Header.Get("Upgrade-Insecure-Requests")
}
func (c *Context) prepareBody() {
	b, err := io.ReadAll(c.r.Body)
	if err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
	defer c.r.Body.Close()
	c.body = b
}
func (c *Context) prepareGet() {
	for k, v := range c.r.URL.Query() {
		if len(v) > 0 {
			c.get[k] = v[0]
		}
	}
	for k, v := range c.routerParameter {
		if len(v) > 0 {
			c.get[k] = v[0]
		}
	}
}
func (c *Context) preparePost() {
	contentType := c.helpContentType()
	switch contentType {
	case "application/x-www-form-urlencoded":
		c.preparePostXWwwFormUrlencoded()
	case "multipart/form-data":
		c.preparePostFormData()
	case "application/json":
		c.preparePostJson()
	default:
		// TODO 暂时不处理
	}
}
func (c *Context) preparePostXWwwFormUrlencoded() {
	if err := c.r.ParseForm(); err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
	for k, v := range c.r.PostForm {
		if len(v) > 0 {
			c.post[k] = v[0]
		}
	}
}
func (c *Context) preparePostFormData() {
	if err := c.r.ParseMultipartForm(32 << 20); err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
	if c.r.MultipartForm != nil {
		for k, v := range c.r.MultipartForm.Value {
			if len(v) > 0 {
				c.post[k] = v[0]
			}
		}
		for k, v := range c.r.MultipartForm.File {
			if len(v) > 0 {
				c.file[k] = v
			}
		}
	}
}
func (c *Context) preparePostJson() {
	if len(c.body) > 0 {
		var data map[string]interface{}
		if err := json.Unmarshal(c.body, &data); err != nil {
			return
		}
		for k, v := range data {
			c.post[k] = v
		}
	}
}
func (c *Context) prepareRequest() {
	for k, v := range c.get {
		c.request[k] = v
	}
	for k, v := range c.post {
		c.request[k] = v
	}
}
func (c *Context) Get(key string) *_parameter.Parameter {
	if v, ok := c.get[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}
func (c *Context) Post(key string) *_parameter.Parameter {
	if v, ok := c.post[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}
func (c *Context) Request(key string) *_parameter.Parameter {
	if v, ok := c.request[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}
func (c *Context) Cookie(key string) *_parameter.Parameter {
	if v, ok := c.cookie[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}
func (c *Context) Header(key string) *_parameter.Parameter {
	key = strings.ToLower(key)
	if v, ok := c.header[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}
func (c *Context) Server(key string) *_parameter.Parameter {
	key = strings.ToLower(key)
	if v, ok := c.server[key]; ok {
		return _parameter.New(key, v)
	}
	return _parameter.New(key, nil)
}
func (c *Context) File(key string) *_parameter.Parameter {
	if files, ok := c.file[key]; ok {
		return _parameter.New(key, files)
	}
	return _parameter.New(key, nil)
}
func (c *Context) Body() []byte {
	return c.body
}
func (c *Context) GetAll() map[string]interface{} {
	return c.get
}
func (c *Context) PostAll() map[string]interface{} {
	return c.post
}
func (c *Context) RequestAll() map[string]interface{} {
	return c.request
}
func (c *Context) CookieAll() map[string]string {
	return c.cookie
}
func (c *Context) HeaderAll() map[string]string {
	return c.header
}
func (c *Context) ServerAll() map[string]string {
	return c.server
}
func (c *Context) FileAll() map[string][]*multipart.FileHeader {
	return c.file
}
func (c *Context) Bind(v interface{}) {
	_bind.Do(v, c.request)
}
func (c *Context) BindGet(v interface{}) {
	_bind.Do(v, c.get)
}
func (c *Context) BindPost(v interface{}) {
	_bind.Do(v, c.post)
}
func (c *Context) BindBody(v interface{}) {
	_json.Decode(c.body, v)
}
func (c *Context) BindRequest(v interface{}) {
	_bind.Do(v, c.request)
}
func (c *Context) helpContentType() string {
	ct := c.r.Header.Get("Content-Type")
	if ct == "" {
		return ""
	}
	parts := strings.Split(strings.ToLower(ct), ";")
	return strings.TrimSpace(parts[0])
}
func (c *Context) helpScheme() string {
	if c.r.TLS != nil {
		return "https"
	}
	if scheme := c.r.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	return "http"
}
func (c *Context) helpFullURL() string {
	return c.helpScheme() + "://" + c.r.Host + c.r.RequestURI
}
func (c *Context) helpClientIP() string {
	if ip := c.r.Header.Get("X-Forwarded-For"); ip != "" {
		ips := strings.Split(ip, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}
	if ip := c.r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	ip := c.r.RemoteAddr
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	return ip
}
func (c *Context) SetHeader(k string, v string) *Context {
	c.w.Header().Set(k, v)
	return c
}
func (c *Context) SetCookie(cookie *http.Cookie) *Context {
	http.SetCookie(c.w, cookie)
	return c
}
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
	c.render.Format("JSON")
	c.render.Response(res)
}
func (c *Context) REDIRECT(uri string) {
	c.w.Header().Set("Location", uri)
	c.w.WriteHeader(301)
}

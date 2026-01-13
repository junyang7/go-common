package _server

import (
	"fmt"
	"github.com/junyang7/go-common/_codeMessage"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_context"
	"github.com/junyang7/go-common/_exception"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_list"
	"github.com/junyang7/go-common/_redis"
	"github.com/junyang7/go-common/_render"
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_router"
	"github.com/junyang7/go-common/_sql"
	"github.com/junyang7/go-common/_unixMilli"
	"net"
	"net/http"
	"regexp"
	"runtime"
	"strings"
)

type apiEngineConf struct {
	Debug   bool     `json:"debug"`   // 调试模式
	Network string   `json:"network"` // 服务协议
	Host    string   `json:"host"`    // 主机
	Port    string   `json:"port"`    // 端口
	Origin  []string `json:"origin"`  // 允许跨域配置列表
}
type apiEngine struct {
	debug   bool
	network string
	host    string
	port    string
	origin  []string
}

func Api() *apiEngine {
	return &apiEngine{
		debug:   false,
		network: "tcp",
		host:    "0.0.0.0",
		port:    "0",
		origin:  []string{},
	}
}
func (this *apiEngine) Load(conf _conf.Conf, business string) *apiEngine {
	_conf.Load(conf)
	_sql.Load()
	_redis.Load()
	var c apiEngineConf
	_json.Decode(_json.Encode(_conf.Get(business).Value()), &c)
	this.debug = c.Debug
	this.network = c.Network
	this.host = c.Host
	this.port = c.Port
	this.origin = c.Origin
	return this
}
func (this *apiEngine) Debug(debug bool) *apiEngine {
	this.debug = debug
	return this
}
func (this *apiEngine) Network(network string) *apiEngine {
	this.network = strings.TrimSpace(network)
	return this
}
func (this *apiEngine) Host(host string) *apiEngine {
	this.host = strings.TrimSpace(host)
	return this
}
func (this *apiEngine) Port(port string) *apiEngine {
	this.port = strings.TrimSpace(port)
	return this
}
func (this *apiEngine) Origin(origin []string) *apiEngine {
	this.origin = origin
	return this
}
func (this *apiEngine) Router(router *_router.Router) *apiEngine {
	_router.RouterList = append(_router.RouterList, router)
	return this
}
func (this *apiEngine) getDebug() bool {
	return this.debug
}
func (this *apiEngine) getNetwork() string {
	if len(this.network) > 0 {
		return this.network
	}
	return "tcp"
}
func (this *apiEngine) getHost() string {
	if len(this.host) > 0 {
		return this.host
	}
	return "0.0.0.0"
}
func (this *apiEngine) getPort() string {
	if len(this.port) > 0 {
		return this.port
	}
	return "0"
}
func (this *apiEngine) getOrigin() []string {
	return this.origin
}
func (this *apiEngine) getAddr() string {
	return fmt.Sprintf("%s:%s", this.getHost(), this.getPort())
}
func (this *apiEngine) check() *apiEngine {
	return this
}
func (this *apiEngine) Run() {
	this.check()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", this.ServeApi)
	server := &http.Server{
		Handler: mux,
	}
	listener, err := net.Listen(this.getNetwork(), this.getAddr())
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	fmt.Printf("Server is running on: %s\n", listener.Addr().String())
	if err := server.Serve(listener); nil != err && err != http.ErrServerClosed {
		_interceptor.Insure(false).Message(err).Do()
	}
	fmt.Println("Server stopped.")
}
func (this *apiEngine) ServeApi(w http.ResponseWriter, r *http.Request) {
	p := &apiEngineProcessor{
		debug:   this.getDebug(),
		network: this.getNetwork(),
		host:    this.getHost(),
		port:    this.getPort(),
		origin:  this.getOrigin(),
		w:       w,
		r:       r,
		render:  _render.New(),
	}
	p.do()
}

type apiEngineProcessor struct {
	debug   bool
	network string
	host    string
	port    string
	origin  []string
	w       http.ResponseWriter
	r       *http.Request
	ctx     *_context.Context
	router  *_router.Router
	render  *_render.Render
}

func (this *apiEngineProcessor) do() {
	defer func() {
		if err := recover(); nil != err {
			this.exception(err)
		}
	}()
	timeS := _unixMilli.Get()
	this.w.Header().Set("access-control-allow-credentials", "true")
	this.checkOrigin()
	if this.r.Method == http.MethodOptions {
		return
	}
	this.checkRouter()
	this.checkRouterMethod()
	this.ctx = _context.New(this.render, timeS, this.w, this.r, this.router.Parameter, this.debug)
	this.doMiddlewareBefore()
	this.doBusiness()
	defer func() {
		this.doMiddlewareAfter()
	}()
}
func (this *apiEngineProcessor) exception(err any) {
	res := _response.New()
	switch err.(type) {
	case *_exception.Exception:
		err := err.(*_exception.Exception)
		res.Code = err.Code
		res.Message = err.Message
		res.Data = err.Data
		break
	default:
		res.Code = _codeMessage.ErrDefault.Code
		res.Message = fmt.Sprintf("%v", err)
		break
	}
	if this.debug {
		if _, file, line, ok := runtime.Caller(5); ok {
			res.File = file
			res.Line = line
		}
	}
	r := _render.New()
	r.Response(res)
	r.Format("JSON")
	r.DoApi(this.w)
}
func (this *apiEngineProcessor) checkOrigin() {
	origin := this.r.Header.Get("Origin")
	matchedList := regexp.MustCompile("(\\S+)://([^:]+):?(\\d+)?").FindStringSubmatch(strings.Trim(origin, "/"))
	if 0 == len(matchedList) {
		return
	}
	for _, origin := range this.origin {
		if "*" == origin || (len(matchedList) >= 3 && matchedList[2] == origin) || (len(origin) > 0 && strings.HasPrefix(origin, ".") && len(matchedList[2]) >= len(origin) && strings.HasSuffix(matchedList[2], origin)) {
			headerValue := matchedList[1] + "://" + matchedList[2]
			if len(matchedList) >= 4 && matchedList[3] != "" {
				headerValue += ":" + matchedList[3]
			}
			this.w.Header().Set("access-control-allow-origin", headerValue)
			this.w.Header().Set("access-control-allow-methods", "GET,POST,OPTIONS")
			this.w.Header().Set("access-control-allow-headers", "content-type, authorization")
			this.w.Header().Set("access-control-expose-headers", "content-type, authorization")
			return
		}
	}
	_interceptor.Insure(false).
		Message(`不支持的跨域请求`).
		Data(map[string]interface{}{`origin`: origin}).
		Do()
}
func (this *apiEngineProcessor) checkRouter() {
	path := this.r.URL.Path
	for _, r := range _router.RouterList {
		if !r.IsRegexp {
			if path == r.Rule {
				this.router = r
				break
			}
			continue
		}
		matchedList := regexp.MustCompile(r.Rule).FindStringSubmatch(path)
		if 0 == len(matchedList) {
			continue
		}
		index := 0
		for k, _ := range r.Parameter {
			r.Parameter[k] = matchedList[index+1]
			index++
		}
		this.router = r
		break
	}
	_interceptor.
		Insure(nil != this.router).
		Message(`不支持的路由地址`).
		Data(map[string]interface{}{`path`: path}).
		Do()
	_interceptor.
		Insure(nil != this.router.Call).
		Message(`路由处理方法未定义`).
		Data(map[string]interface{}{`path`: path}).
		Do()
}
func (this *apiEngineProcessor) checkRouterMethod() {
	method := this.r.Method
	_interceptor.
		Insure(_list.In(`ANY`, this.router.MethodList) || _list.In(method, this.router.MethodList)).
		Message(`不支持的请求方法`).
		Data(map[string]interface{}{`method`: method}).
		Do()
}
func (this *apiEngineProcessor) doMiddlewareBefore() {
	for _, middleware := range this.router.MiddlewareBeforeList {
		middleware(this.ctx)
	}
}
func (this *apiEngineProcessor) doBusiness() {
	this.router.Call(this.ctx)
	this.render.DoApi(this.w)
}
func (this *apiEngineProcessor) doMiddlewareAfter() {
	for _, middleware := range this.router.MiddlewareAfterList {
		middleware(this.ctx)
	}
}

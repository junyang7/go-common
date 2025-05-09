package _server

import (
	"context"
	"fmt"
	"github.com/junyang7/go-common/_codeMessage"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_context"
	"github.com/junyang7/go-common/_exception"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_list"
	"github.com/junyang7/go-common/_pb"
	"github.com/junyang7/go-common/_redis"
	"github.com/junyang7/go-common/_render"
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_router"
	"github.com/junyang7/go-common/_sql"
	"github.com/junyang7/go-common/_structure"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func load(conf _conf.Conf) {
	_conf.Load(conf)
	_sql.Load()
	_redis.Load()
}

type webEngine struct {
	debug   bool
	network string
	host    string
	port    string
	origin  []string
	root    string
}

func Web() *webEngine {
	return &webEngine{}
}
func (this *webEngine) Load(conf _conf.Conf, business string) *webEngine {
	load(conf)
	raw := _conf.Get(business).Value()
	var serverWeb _structure.ServerWeb
	_json.Decode(_json.Encode(raw), &serverWeb)
	this.host = serverWeb.Host
	this.port = serverWeb.Port
	this.debug = serverWeb.Debug
	this.origin = serverWeb.Origin
	this.root = serverWeb.Root
	return this
}
func (this *webEngine) Debug(debug bool) *webEngine {
	this.debug = debug
	return this
}
func (this *webEngine) Network(network string) *webEngine {
	this.network = network
	return this
}
func (this *webEngine) Host(host string) *webEngine {
	this.host = host
	return this
}
func (this *webEngine) Port(port string) *webEngine {
	this.port = port
	return this
}
func (this *webEngine) Origin(origin []string) *webEngine {
	this.origin = origin
	return this
}
func (this *webEngine) Root(root string) *webEngine {
	this.root = root
	return this
}
func (this *webEngine) getDebug() bool {
	return this.debug
}
func (this *webEngine) getNetwork() string {
	if len(this.network) > 0 {
		return this.network
	}
	return "tcp"
}
func (this *webEngine) getHost() string {
	if len(this.host) > 0 {
		return this.host
	}
	return "0.0.0.0"
}
func (this *webEngine) getPort() string {
	if len(this.port) > 0 {
		return this.port
	}
	return "0"
}
func (this *webEngine) getOrigin() []string {
	return this.origin
}
func (this *webEngine) getRoot() string {
	return this.root
}
func (this *webEngine) getAddr() string {
	return fmt.Sprintf("%s:%s", this.host, this.port)
}
func (this *webEngine) Run() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(this.getRoot())))
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

type apiEngine struct {
	debug   bool
	network string
	host    string
	port    string
	origin  []string
	handler *http.Server
}

func Api() *apiEngine {
	return &apiEngine{}
}
func (this *apiEngine) Load(conf _conf.Conf, business string) *apiEngine {
	load(conf)
	raw := _conf.Get(business).Value()
	var serverApi _structure.ServerApi
	_json.Decode(_json.Encode(raw), &serverApi)
	this.host = serverApi.Host
	this.port = serverApi.Port
	this.debug = serverApi.Debug
	this.origin = serverApi.Origin
	return this
}
func (this *apiEngine) Debug(debug bool) *apiEngine {
	this.debug = debug
	return this
}
func (this *apiEngine) Network(network string) *apiEngine {
	this.network = network
	return this
}
func (this *apiEngine) Host(host string) *apiEngine {
	this.host = host
	return this
}
func (this *apiEngine) Port(port string) *apiEngine {
	this.port = port
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
	return fmt.Sprintf("%s:%s", this.host, this.port)
}
func (this *apiEngine) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", this.ServeHTTP)
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
func (this *apiEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := &apiProcessor{
		w:      w,
		r:      r,
		origin: this.origin,
		debug:  this.getDebug(),
	}
	p.do()
}

type apiProcessor struct {
	w      http.ResponseWriter
	r      *http.Request
	ctx    *_context.Context
	origin []string
	router *_router.Router
	debug  bool
}

func (this *apiProcessor) do() {
	defer func() {
		if err := recover(); nil != err {
			this.exception(err)
		}
	}()
	this.ctx = _context.New(this.w, this.r, this.debug)
	this.w.Header().Set("access-control-allow-credentials", "true")
	this.checkOrigin()
	if this.ctx.SERVER["method"] == "OPTIONS" {
		return
	}
	this.checkRouter()
	this.checkRouterMethod()
	this.middlewareBefore()
	this.business()
}
func (this *apiProcessor) checkOrigin() {
	origin := this.ctx.ServerParameter("origin").String().Value()
	matchedList := regexp.MustCompile("(\\S+)://([^:]+):?(\\d+)?").FindStringSubmatch(strings.Trim(origin, "/"))
	if 0 == len(matchedList) {
		return
	}
	for _, origin := range this.origin {
		if "*" == origin || matchedList[2] == origin || "." == origin[0:1] && matchedList[2][len(matchedList[2])-len(origin):] == origin {
			headerValue := matchedList[1] + "://" + matchedList[2]
			if 4 == len(matchedList) {
				headerValue += ":" + matchedList[3]
			}
			this.w.Header().Set("access-control-allow-origin", headerValue)
			this.w.Header().Set("access-control-allow-Headers", "content-type, authorization")
			return
		}
	}
	_interceptor.Insure(false).
		Message(`不支持的跨域请求`).
		Data(map[string]interface{}{`origin`: origin}).
		Do()
}
func (this *apiProcessor) checkRouter() {
	path := this.ctx.ServerParameter(`path`).String().Value()
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
		this.router = r
		for index, parameter := range r.ParameterList {
			this.ctx.GET[parameter] = matchedList[index+1]
			this.ctx.POST[parameter] = this.ctx.GET[parameter]
		}
	}
	_interceptor.Insure(nil != this.router).
		Message(`不支持的路由地址`).
		Data(map[string]interface{}{`path`: path}).
		Do()
	_interceptor.Insure(nil != this.router.Call).
		Message(`路由处理方法未定义`).
		Data(map[string]interface{}{`path`: path}).
		Do()
}
func (this *apiProcessor) checkRouterMethod() {
	method := this.ctx.ServerParameter(`method`).String().Value()
	_interceptor.Insure(_list.In(`ANY`, this.router.MethodList) || _list.In(method, this.router.MethodList)).
		Message(`不支持的请求方法`).
		Data(map[string]interface{}{`method`: method}).
		Do()
}
func (this *apiProcessor) middlewareBefore() {
	for _, middleware := range this.router.MiddlewareBeforeList {
		middleware(this.ctx)
	}
}
func (this *apiProcessor) business() {
	this.router.Call(this.ctx)
}
func (this *apiProcessor) middlewareAfter() {
	for _, middleware := range this.router.MiddlewareAfterList {
		middleware(this.ctx)
	}
}
func (this *apiProcessor) exception(err any) {
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
	if _, file, line, ok := runtime.Caller(5); ok {
		res.File = file
		res.Line = line
	}
	_render.JSON(this.w, res)
}

type httpEngine struct {
	debug   bool
	network string
	host    string
	port    string
	origin  []string
	root    string
}

func Http() *httpEngine {
	return &httpEngine{}
}
func (this *httpEngine) Load(conf _conf.Conf, business string) *httpEngine {
	load(conf)
	raw := _conf.Get(business).Value()
	var serverHttp _structure.ServerHttp
	_json.Decode(_json.Encode(raw), &serverHttp)
	this.host = serverHttp.Host
	this.port = serverHttp.Port
	this.debug = serverHttp.Debug
	this.origin = serverHttp.Origin
	this.root = serverHttp.Root
	return this
}
func (this *httpEngine) Debug(debug bool) *httpEngine {
	this.debug = debug
	return this
}
func (this *httpEngine) Network(network string) *httpEngine {
	this.network = network
	return this
}
func (this *httpEngine) Host(host string) *httpEngine {
	this.host = host
	return this
}
func (this *httpEngine) Port(port string) *httpEngine {
	this.port = port
	return this
}
func (this *httpEngine) Origin(origin []string) *httpEngine {
	this.origin = origin
	return this
}
func (this *httpEngine) Root(root string) *httpEngine {
	this.root = root
	return this
}
func (this *httpEngine) Router(router *_router.Router) *httpEngine {
	_router.RouterList = append(_router.RouterList, router)
	return this
}
func (this *httpEngine) getDebug() bool {
	return this.debug
}
func (this *httpEngine) getNetwork() string {
	if len(this.network) > 0 {
		return this.network
	}
	return "tcp"
}
func (this *httpEngine) getHost() string {
	if len(this.host) > 0 {
		return this.host
	}
	return "0.0.0.0"
}
func (this *httpEngine) getPort() string {
	if len(this.port) > 0 {
		return this.port
	}
	return "0"
}
func (this *httpEngine) getOrigin() []string {
	return this.origin
}
func (this *httpEngine) getRoot() string {
	return this.root
}
func (this *httpEngine) getAddr() string {
	return fmt.Sprintf("%s:%s", this.host, this.port)
}
func (this *httpEngine) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", Api().Debug(this.debug).Host(this.host).Port(this.port).Origin(this.origin).ServeHTTP)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		root := this.getRoot()

		// 拼接用户请求路径对应的文件路径
		requestPath := r.URL.Path
		fullPath := filepath.Join(root, filepath.Clean("/"+requestPath)) // 避免裸路径拼错

		// 防止路径穿越攻击（确保 fullPath 在 root 内）
		rel, err := filepath.Rel(root, fullPath)
		if err != nil || strings.HasPrefix(rel, "..") {
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}

		// 检查文件是否存在且不是目录
		info, err := os.Stat(fullPath)
		if err == nil && !info.IsDir() {
			http.ServeFile(w, r, fullPath)
			return
		}

		// 不存在的路径，返回 index.html 支持 Vue history 模式
		http.ServeFile(w, r, filepath.Join(root, "index.html"))
	})

	listener, err := net.Listen(this.getNetwork(), this.getAddr())
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	fmt.Printf("Server is running on: %s\n", listener.Addr().String())
	server := &http.Server{
		Handler: mux,
	}
	if err := server.Serve(listener); nil != err && err != http.ErrServerClosed {
		_interceptor.Insure(false).Message(err).Do()
	}
}

type cliEngine struct{}

func Cli() *cliEngine {
	return &cliEngine{}
}

type fileEngine struct{}

func File() *fileEngine {
	return &fileEngine{}
}

type jobEngine struct{}

func Job() *jobEngine {
	return &jobEngine{}
}

type rpcEngine struct {
	conf    _conf.Conf
	network string
	addr    string
	debug   bool
}

func Rpc() *rpcEngine {
	return &rpcEngine{}
}
func (this *rpcEngine) Conf(conf _conf.Conf) *rpcEngine {
	this.conf = conf
	return this
}
func (this *rpcEngine) Network(network string) *rpcEngine {
	this.network = network
	return this
}
func (this *rpcEngine) Addr(addr string) *rpcEngine {
	this.addr = addr
	return this
}
func (this *rpcEngine) Router(router *_router.Router) *rpcEngine {
	_router.RouterList = append(_router.RouterList, router)
	return this
}
func (this *rpcEngine) Debug(debug bool) *rpcEngine {
	this.debug = debug
	return this
}
func (this *rpcEngine) Run() {
	listener, err := net.Listen(this.network, this.addr)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	fmt.Printf("Server is running on: %s\n", listener.Addr().String())
	s := grpc.NewServer()
	_pb.RegisterServiceServer(s, &rpcCall{})
	if err := s.Serve(listener); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

type rpcCall struct {
	_pb.UnimplementedServiceServer
}

func (this *rpcCall) Call(c context.Context, r *_pb.Request) (oRes *_pb.Response, oErr error) {

	// 接受请求数据
	// 处理业务逻辑
	//fmt.Println("<====")
	//fmt.Println(r.Header)
	//fmt.Println(r.Body)
	//var a map[string]interface{}
	//_json.Decode(r.Body, &a)
	//fmt.Println(a)

	res := _response.New()
	defer func() {
		if err := recover(); nil != err {

			// 同api处理方法
			// 判断异常类型，拼接返回

			res.Code = -1
			res.Message = fmt.Sprintf("%v", err)
			oRes = &_pb.Response{Response: _json.Encode(res)}
		}
	}()

	//// 业务逻辑返回数据
	//// 需要处理异常
	//res.Code = 0
	//res.Message = "success"
	//res.Data = map[string]string{"test": "Hello World!"}
	//oRes = &pb2.Response{Response: _json.Encode(res)}

	return oRes, oErr

}

type rpcCallProcessor struct {
}

func (this *rpcCallProcessor) do() (body []byte, header map[string]string) {
	//this.checkRouter()
	return nil, nil
}

type websocketEngine struct{}

func Websocket() *websocketEngine {
	return &websocketEngine{}
}

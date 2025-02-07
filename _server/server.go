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
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_router"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"regexp"
	"runtime"
	"strings"
)

type webEngine struct {
	addr  string
	root  string
	debug bool
}

func Web() *webEngine {
	return &webEngine{}
}
func (this *webEngine) Addr(addr string) *webEngine {
	this.addr = addr
	return this
}
func (this *webEngine) Root(root string) *webEngine {
	this.root = root
	return this
}
func (this *webEngine) Debug(debug bool) *webEngine {
	this.debug = debug
	return this
}
func (this *webEngine) Run() {
	http.Handle("/", http.FileServer(http.Dir(this.root)))
	if err := http.ListenAndServe(this.addr, nil); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

type apiEngine struct {
	addr    string
	debug   bool
	origin  []string
	handler *http.Server
}

func Api() *apiEngine {
	return &apiEngine{}
}
func (this *apiEngine) Addr(addr string) *apiEngine {
	this.addr = addr
	return this
}
func (this *apiEngine) Debug(debug bool) *apiEngine {
	this.debug = debug
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
func (this *apiEngine) Run() {
	this.handler = &http.Server{
		Addr:    this.addr,
		Handler: http.HandlerFunc(this.ServeHTTP),
	}
	if err := this.handler.ListenAndServe(); nil != err && http.ErrServerClosed != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func (this *apiEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := &apiProcessor{
		w:      w,
		r:      r,
		ctx:    _context.New(w, r, this.debug),
		origin: this.origin,
	}
	p.do()
}

type apiProcessor struct {
	w      http.ResponseWriter
	r      *http.Request
	ctx    *_context.Context
	origin []string
	router *_router.Router
}

func (this *apiProcessor) do() {
	defer func() {
		if err := recover(); nil != err {
			this.exception(err)
		}
	}()
	this.w.Header().Add("access-control-allow-credentials", "true")
	this.checkOrigin()
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
	this.ctx.JSON(res)
}

type httpEngine struct {
	addr   string
	root   string
	debug  bool
	origin []string
}

func Http() *httpEngine {
	return &httpEngine{}
}
func (this *httpEngine) Addr(addr string) *httpEngine {
	this.addr = addr
	return this
}
func (this *httpEngine) Root(root string) *httpEngine {
	this.root = root
	return this
}
func (this *httpEngine) Debug(debug bool) *httpEngine {
	this.debug = debug
	return this
}
func (this *httpEngine) Origin(origin []string) *httpEngine {
	this.origin = origin
	return this
}
func (this *httpEngine) Router(router *_router.Router) *httpEngine {
	_router.RouterList = append(_router.RouterList, router)
	return this
}
func (this *httpEngine) Run() {
	http.HandleFunc("/api/", Api().Addr(this.addr).Debug(this.debug).Origin(this.origin).ServeHTTP)
	Web().Addr(this.addr).Root(this.root).Debug(this.debug).Run()
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
	l, err := net.Listen(this.network, this.addr)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	s := grpc.NewServer()
	_pb.RegisterServiceServer(s, &rpcCall{})
	if err := s.Serve(l); nil != err {
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

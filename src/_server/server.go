package _server

import (
	"context"
	"encoding/json"
	"fmt"
	pb2 "github.com/junyang7/go-common/src/_client/pb"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_context"
	"github.com/junyang7/go-common/src/_exception"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_list"
	"github.com/junyang7/go-common/src/_response"
	"github.com/junyang7/go-common/src/_router"
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
		ctx:    _context.New(w, r, true),
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
	network string
	addr    string
	debug   bool
}

func Rpc() *rpcEngine {
	return &rpcEngine{}
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
	pb2.RegisterServiceServer(s, &rpcProcessor{})
	if err := s.Serve(l); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}

type rpcProcessor struct {
	pb2.UnimplementedServiceServer
}

func (this *rpcProcessor) Call(c context.Context, r *pb2.Request) (*pb2.Response, error) {
	defer func() {
		if err := recover(); nil != err {
			fmt.Println(err)
		}
	}()
	b, _ := json.Marshal(struct {
		Header interface{} `json:"header"`
		Body   interface{} `json:"body"`
	}{
		Header: r.Header,
		Body:   r.Body,
	})
	res := &pb2.Response{Response: b}
	return res, nil
}

type websocketEngine struct{}

func Websocket() *websocketEngine {
	return &websocketEngine{}
}

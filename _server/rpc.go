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
	"github.com/junyang7/go-common/_pb"
	"github.com/junyang7/go-common/_redis"
	"github.com/junyang7/go-common/_render"
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_router"
	"github.com/junyang7/go-common/_sql"
	"github.com/junyang7/go-common/_unixMilli"
	"google.golang.org/grpc"
	"net"
	"regexp"
	"runtime"
	"strings"
)

type rpcEngineConf struct {
	Debug   bool   `json:"debug"`   // 调试模式
	Network string `json:"network"` // 服务协议
	Host    string `json:"host"`    // 主机
	Port    string `json:"port"`    // 端口
}
type rpcEngine struct {
	debug   bool
	network string
	host    string
	port    string
	_pb.UnimplementedServiceServer
}

func Rpc() *rpcEngine {
	return &rpcEngine{
		debug:   false,
		network: "tcp",
		host:    "0.0.0.0",
		port:    "0",
	}
}
func (this *rpcEngine) Load(conf _conf.Conf, business string) *rpcEngine {
	_conf.Load(conf)
	_sql.Load()
	_redis.Load()
	var c rpcEngineConf
	_json.Decode(_json.Encode(_conf.Get(business).Value()), &c)
	this.debug = c.Debug
	this.network = c.Network
	this.host = c.Host
	this.port = c.Port
	return this
}
func (this *rpcEngine) Debug(debug bool) *rpcEngine {
	this.debug = debug
	return this
}
func (this *rpcEngine) Network(network string) *rpcEngine {
	this.network = strings.TrimSpace(network)
	return this
}
func (this *rpcEngine) Host(host string) *rpcEngine {
	this.host = strings.TrimSpace(host)
	return this
}
func (this *rpcEngine) Port(port string) *rpcEngine {
	this.port = strings.TrimSpace(port)
	return this
}
func (this *rpcEngine) Router(router *_router.Router) *rpcEngine {
	_router.RouterList = append(_router.RouterList, router)
	return this
}
func (this *rpcEngine) getDebug() bool {
	return this.debug
}
func (this *rpcEngine) getNetwork() string {
	if len(this.network) > 0 {
		return this.network
	}
	return "tcp"
}
func (this *rpcEngine) getHost() string {
	if len(this.host) > 0 {
		return this.host
	}
	return "0.0.0.0"
}
func (this *rpcEngine) getPort() string {
	if len(this.port) > 0 {
		return this.port
	}
	return "0"
}
func (this *rpcEngine) getAddr() string {
	return fmt.Sprintf("%s:%s", this.getHost(), this.getPort())
}
func (this *rpcEngine) check() *rpcEngine {
	return this
}
func (this *rpcEngine) Run() {
	this.check()
	listener, err := net.Listen(this.getNetwork(), this.getAddr())
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	fmt.Printf("Server is running on: %s\n", listener.Addr().String())
	s := grpc.NewServer()
	_pb.RegisterServiceServer(s, this)
	if err := s.Serve(listener); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	fmt.Println("Server stopped.")
}
func (this *rpcEngine) Call(c context.Context, r *_pb.Request) (w *_pb.Response, err error) {
	p := &rpcEngineProcessor{
		debug:   this.debug,
		network: this.network,
		host:    this.host,
		port:    this.port,
		w:       w,
		r:       r,
		render:  _render.New(),
	}
	return p.do()
}

type rpcEngineProcessor struct {
	debug   bool
	network string
	host    string
	port    string
	w       *_pb.Response
	r       *_pb.Request
	ctx     *_context.Context
	router  *_router.Router
	render  *_render.Render
}

func (this *rpcEngineProcessor) do() (o *_pb.Response, err error) {
	defer func() {
		if err := recover(); nil != err {
			o = this.exception(err)
		}
	}()
	timeS := _unixMilli.Get()
	this.checkRouter()
	this.ctx = _context.New(this.render, timeS, nil, nil, this.router.Parameter, this.debug)
	this.doMiddlewareBefore()
	o = this.doBusiness()
	defer func() {
		this.doMiddlewareAfter()
	}()
	return o, nil
}
func (this *rpcEngineProcessor) exception(err any) (o *_pb.Response) {
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
	o = r.DoRpc()
	return o
}
func (this *rpcEngineProcessor) checkRouter() {
	path, ok := this.r.Header["path"]
	_interceptor.
		Insure(ok && path != "").
		Message(`没有设置path`).
		Data(map[string]interface{}{`path`: path}).
		Do()
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
func (this *rpcEngineProcessor) doMiddlewareBefore() {
	for _, middleware := range this.router.MiddlewareBeforeList {
		middleware(this.ctx)
	}
}
func (this *rpcEngineProcessor) doBusiness() (o *_pb.Response) {
	this.router.Call(this.ctx)
	o = this.render.DoRpc()
	return o
}
func (this *rpcEngineProcessor) doMiddlewareAfter() {
	for _, middleware := range this.router.MiddlewareAfterList {
		middleware(this.ctx)
	}
}

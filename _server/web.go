package _server

import (
	"fmt"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"net"
	"net/http"
	"regexp"
	"strings"
)

type webEngineConf struct {
	Debug   bool     `json:"debug"`   // 调试模式
	Network string   `json:"network"` // 服务协议
	Host    string   `json:"host"`    // 主机
	Port    string   `json:"port"`    // 端口
	Origin  []string `json:"origin"`  // 允许跨域配置列表
	Root    string   `json:"root"`    // 文件根目录
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
	return &webEngine{
		debug:   false,
		network: "tcp",
		host:    "0.0.0.0",
		port:    "0",
		origin:  []string{},
		root:    "",
	}
}
func (this *webEngine) Load(conf _conf.Conf, business string) *webEngine {
	_conf.Load(conf)
	var c webEngineConf
	_json.Decode(_json.Encode(_conf.Get(business).Value()), &c)
	this.debug = c.Debug
	this.network = c.Network
	this.host = c.Host
	this.port = c.Port
	this.origin = c.Origin
	this.root = c.Root
	return this
}
func (this *webEngine) Debug(debug bool) *webEngine {
	this.debug = debug
	return this
}
func (this *webEngine) Network(network string) *webEngine {
	this.network = strings.TrimSpace(network)
	return this
}
func (this *webEngine) Host(host string) *webEngine {
	this.host = strings.TrimSpace(host)
	return this
}
func (this *webEngine) Port(port string) *webEngine {
	this.port = strings.TrimSpace(port)
	return this
}
func (this *webEngine) Origin(origin []string) *webEngine {
	this.origin = origin
	return this
}
func (this *webEngine) Root(root string) *webEngine {
	this.root = strings.TrimSpace(root)
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
	return fmt.Sprintf("%s:%s", this.getHost(), this.getPort())
}
func (this *webEngine) check() *webEngine {
	this.checkRoot()
	return this
}
func (this *webEngine) checkRoot() *webEngine {
	_interceptor.
		Insure(this.root != "").
		Data(map[string]string{"root": this.root}).
		Message("请设置web根目录").
		Do()
	_interceptor.
		Insure(strings.HasPrefix(this.root, "/")).
		Data(map[string]string{"root": this.root}).
		Message("web根目录必须是绝对路径").
		Do()
	_interceptor.
		Insure(_directory.Exists(this.root)).
		Data(map[string]string{"root": this.root}).
		Message("web根目录不存在").
		Do()
	return this
}
func (this *webEngine) Run() {
	this.check()
	mux := http.NewServeMux()
	mux.HandleFunc("/", this.ServeWeb)
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
func (this *webEngine) ServeWeb(w http.ResponseWriter, r *http.Request) {
	p := &webEngineProcessor{
		debug:   this.getDebug(),
		network: this.getNetwork(),
		host:    this.getHost(),
		port:    this.getPort(),
		origin:  this.getOrigin(),
		root:    this.getRoot(),
		w:       w,
		r:       r,
	}
	p.do()
}

type webEngineProcessor struct {
	debug   bool
	network string
	host    string
	port    string
	origin  []string
	root    string
	w       http.ResponseWriter
	r       *http.Request
}

func (this *webEngineProcessor) do() {
	defer func() {
		if err := recover(); nil != err {
			this.exception(err)
		}
	}()
	this.checkOrigin()
	this.w.Header().Set("access-control-allow-credentials", "true")
	if this.r.Method == http.MethodOptions {
		return
	}
	this.doMiddlewareBefore()
	this.doBusiness()
	this.doMiddlewareAfter()
}
func (this *webEngineProcessor) exception(err any) {
	// TODO 预留，记录日志，返回404或者500
}
func (this *webEngineProcessor) checkOrigin() {
	if len(this.origin) == 0 {
		return
	}
	origin := this.r.Header.Get("Origin")
	matchedList := regexp.MustCompile("(\\S+)://([^:]+):?(\\d+)?").FindStringSubmatch(strings.Trim(origin, "/"))
	if len(matchedList) == 0 {
		return
	}
	for _, o := range this.origin {
		if o == "*" || (len(matchedList) >= 3 && matchedList[2] == o) || (strings.HasPrefix(o, ".") && strings.HasSuffix(matchedList[2], o)) {
			headerValue := matchedList[1] + "://" + matchedList[2]
			if len(matchedList) >= 4 && matchedList[3] != "" {
				headerValue += ":" + matchedList[3]
			}
			this.w.Header().Set("access-control-allow-origin", headerValue)
			this.w.Header().Set("access-control-allow-methods", "GET,POST,OPTIONS")
			this.w.Header().Set("access-control-allow-headers", "content-type,authorization")
			this.w.Header().Set("access-control-expose-headers", "content-type,authorization")
			return
		}
	}
}
func (this *webEngineProcessor) doMiddlewareBefore() {
	// TODO 预留
}
func (this *webEngineProcessor) doBusiness() {
	http.FileServer(http.Dir(this.root)).ServeHTTP(this.w, this.r)
}
func (this *webEngineProcessor) doMiddlewareAfter() {
	// TODO 预留
}

package _server

import (
	"fmt"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_redis"
	"github.com/junyang7/go-common/_sql"
	"net"
	"net/http"
	"sync"
)

type httpEngineConf struct {
	Web webEngineConf `json:"web"`
	Api apiEngineConf `json:"api"`
}
type httpEngine struct {
	web *webEngine
	api *apiEngine
}

func Http() *httpEngine {
	return &httpEngine{
		web: Web(),
		api: Api(),
	}
}
func (this *httpEngine) Load(conf _conf.Conf, business string) *httpEngine {
	_conf.Load(conf)
	_sql.Load()
	_redis.Load()
	var c httpEngineConf
	_json.Decode(_json.Encode(_conf.Get(business).Value()), &c)
	this.web.debug = c.Web.Debug
	this.web.network = c.Web.Network
	this.web.host = c.Web.Host
	this.web.port = c.Web.Port
	this.web.origin = c.Web.Origin
	this.web.root = c.Web.Root
	this.api.debug = c.Api.Debug
	this.api.network = c.Api.Network
	this.api.host = c.Api.Host
	this.api.port = c.Api.Port
	this.api.origin = c.Api.Origin
	return this
}
func (this *httpEngine) WebEngine() *webEngine {
	return this.web
}
func (this *httpEngine) ApiEngine() *apiEngine {
	return this.api
}
func (this *httpEngine) Run() {
	webNetwork := this.web.getNetwork()
	apiNetwork := this.api.getNetwork()
	webAddr := this.web.getAddr()
	apiAddr := this.api.getAddr()
	wg := sync.WaitGroup{}
	if webNetwork == apiNetwork && webAddr == apiAddr {
		this.api.check()
		this.web.check()
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			this.start(&httpEngineProcessor{
				network: webNetwork,
				addr:    webAddr,
				handleList: map[string]func(w http.ResponseWriter, r *http.Request){
					"/api/": this.api.check().ServeApi, // 注意顺序
					"/":     this.web.check().ServeWeb,
				}})
		}()
	} else {
		this.web.check()
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			this.start(&httpEngineProcessor{
				network: webNetwork,
				addr:    webAddr,
				handleList: map[string]func(w http.ResponseWriter, r *http.Request){
					"/": this.web.check().ServeWeb,
				}})
		}()
		this.api.check()
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			this.start(&httpEngineProcessor{
				network: apiNetwork,
				addr:    apiAddr,
				handleList: map[string]func(w http.ResponseWriter, r *http.Request){
					"/api/": this.api.check().ServeApi,
				}})
		}()
	}
	wg.Wait()
}

type httpEngineProcessor struct {
	network    string
	addr       string
	handleList map[string]func(w http.ResponseWriter, r *http.Request)
}

func (this *httpEngine) start(p *httpEngineProcessor) {
	mux := http.NewServeMux()
	for r, f := range p.handleList {
		mux.HandleFunc(r, f)
	}
	server := &http.Server{
		Handler: mux,
	}
	listener, err := net.Listen(p.network, p.addr)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	fmt.Printf("Server is running on: %s\n", listener.Addr().String())
	if err := server.Serve(listener); nil != err && err != http.ErrServerClosed {
		_interceptor.Insure(false).Message(err).Do()
	}
	fmt.Println("Server stopped.")
}

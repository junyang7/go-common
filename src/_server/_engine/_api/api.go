package _api

import (
	"github.com/junyang7/go-common/src/_context"
	"github.com/junyang7/go-common/src/_exception"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_render"
	"github.com/junyang7/go-common/src/_response"
	"github.com/junyang7/go-common/src/_server/_router"
	"net/http"
	"regexp"
	"strings"
)

type Conf struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
	Ipv4 struct {
		Black []string
		White []string
	}
	Method struct {
		Black []string
		White []string
	}
	Origin []string
	Header map[string]string
}

type engine struct {
	conf *Conf
}

func Initialize(conf *Conf) {
	this := &engine{conf: conf}
	err := http.ListenAndServe(this.conf.Ip+":"+this.conf.Port, this)
	_interceptor.Insure(nil != err).Do()
}

func (this *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := &processor{
		ctx: &_context.Context{W: w, R: r},
		w:   w,
		r:   r,
	}
	p.do()
}

type processor struct {
	conf   *Conf
	ctx    *_context.Context
	w      http.ResponseWriter
	r      *http.Request
	router *_router.Router
}

func (this *processor) do() {
	defer func() {
		if err := recover(); nil != err {
			this.exception(err)
		}
	}()
	this.checkIp()
	this.checkMethod()
	this.checkOrigin()
	this.checkRouter()
	this.checkRouterMethod()
	this.middlewareBefore()
	this.business()
	this.middlewareAfter()
}

func (this *processor) business() {
	this.router.Handler(this.ctx)
}
func (this *processor) checkIp() {

}
func (this *processor) checkMethod() {

}
func (this *processor) checkOrigin() {
	origin := this.ctx.Server("origin").String()
	matchedList := regexp.MustCompile("(\\S+)://([^:]+):?(\\d+)?").FindStringSubmatch(strings.Trim(origin, "/"))
	if 0 == len(matchedList) {
		return
	}
	for _, origin := range this.conf.Origin {
		if "*" == origin || matchedList[2] == origin || "." == origin[0:1] && matchedList[2][len(matchedList[2])-len(origin):] == origin {
			headerValue := matchedList[1] + "://" + matchedList[2]
			if 4 == len(matchedList) {
				headerValue += ":" + matchedList[3]
			}
			this.conf.Header["access-control-allow-origin"] = headerValue
			return
		}
	}
	panic("跨域阻止")
}
func (this *processor) checkRouter() {
	path := this.ctx.Server("path").String()
	if router, ok := _router.RouterMap[path]; ok {
		this.router = router
		return
	}
	panic("资源不存在")
}
func (this *processor) checkRouterMethod() {

}
func (this *processor) middlewareAfter() {

}
func (this *processor) middlewareBefore() {

}
func (this *processor) exception(err interface{}) {
	res := _response.New()
	switch err.(type) {
	case *_exception.Exception:
		err := err.(*_exception.Exception)
		res.Code = err.Code
		res.Message = err.Message
		res.Data = err.Data
	case string:
		res.Code = -1
		res.Message = err.(string)
	default:
		res.Code = -1
		res.Message = "failure"
	}
	_render.New(this.w, this.r).Json(res)
}

package _server

import (
	"github.com/junyang7/go-common/src/_context"
	"github.com/junyang7/go-common/src/_response"
	"github.com/junyang7/go-common/src/_server/_conf"
	"github.com/junyang7/go-common/src/_server/_router"
	"net/http"
	"regexp"
	"strings"
)

type engine struct {
	mode     string
	conf     *_conf.Conf
	response *_response.Response
	ctx      *_context.Context
	w        http.ResponseWriter
	r        *http.Request
	router   *_router.Router
}

func (this *engine) do() {
	defer func() {
		if err := recover(); nil != err {
			this.ctx.Json(err)
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

func (this *engine) business() {
	this.router.Handler(this.ctx)
}
func (this *engine) checkIp() {

}
func (this *engine) checkMethod() {

}
func (this *engine) checkOrigin() {
	origin := this.ctx.Server("origin")
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
func (this *engine) checkRouter() {
	path := this.ctx.Server("path")
	if router, ok := _router.RouterMap[path]; ok {
		this.router = router
		return
	}
	panic("资源不存在")
}
func (this *engine) checkRouterMethod() {

}
func (this *engine) middlewareAfter() {

}
func (this *engine) middlewareBefore() {

}

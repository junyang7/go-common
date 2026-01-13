package _router

import (
	"github.com/junyang7/go-common/_context"
	"github.com/junyang7/go-common/_is"
	"github.com/junyang7/go-common/_list"
	"regexp"
	"strings"
)

type router struct {
	business             string                        // 业务标识
	prefix               string                        // 前缀
	middlewareBeforeList []func(ctx *_context.Context) // 中间件：前
	middlewareAfterList  []func(ctx *_context.Context) // 中间件：后
	methodList           []string                      // 方法列表
}

func Any(rule string, call func(ctx *_context.Context)) {
	MethodList([]string{"ANY"}, rule, call)
}
func Get(rule string, call func(ctx *_context.Context)) {
	MethodList([]string{"GET"}, rule, call)
}
func Post(rule string, call func(ctx *_context.Context)) {
	MethodList([]string{"POST"}, rule, call)
}
func Put(rule string, call func(ctx *_context.Context)) {
	MethodList([]string{"PUT"}, rule, call)
}
func Delete(rule string, call func(ctx *_context.Context)) {
	MethodList([]string{"DELETE"}, rule, call)
}
func Options(rule string, call func(ctx *_context.Context)) {
	MethodList([]string{"OPTIONS"}, rule, call)
}
func Head(rule string, call func(ctx *_context.Context)) {
	MethodList([]string{"HEAD"}, rule, call)
}
func Patch(rule string, call func(ctx *_context.Context)) {
	MethodList([]string{"PATCH"}, rule, call)
}
func Method(method string, rule string, call func(ctx *_context.Context)) {
	MethodList([]string{method}, rule, call)
}
func MethodList(methodList []string, rule string, call func(ctx *_context.Context)) {
	r := &router{}
	r.MethodList(methodList, rule, call)
}
func MiddlewareBefore(middleware func(ctx *_context.Context)) *router {
	return MiddlewareBeforeList([]func(ctx *_context.Context){middleware})
}
func MiddlewareBeforeList(middlewareList []func(ctx *_context.Context)) *router {
	r := &router{}
	r.MiddlewareBeforeList(middlewareList)
	return r
}
func MiddlewareAfter(middleware func(ctx *_context.Context)) *router {
	return MiddlewareAfterList([]func(ctx *_context.Context){middleware})
}
func MiddlewareAfterList(middlewareList []func(ctx *_context.Context)) *router {
	r := &router{}
	r.MiddlewareAfterList(middlewareList)
	return r
}
func Prefix(prefix string) *router {
	r := &router{}
	r.Prefix(prefix)
	return r
}
func Business(business string) *router {
	r := &router{}
	r.Business(business)
	return r
}
func Group(group func()) {
	group()
}

type Router struct {
	Business             string                        // 业务标识
	Rule                 string                        // 路径
	Call                 func(ctx *_context.Context)   // 处理方法
	MethodList           []string                      // 方法列表
	MiddlewareBeforeList []func(ctx *_context.Context) // 前中间件列表
	MiddlewareAfterList  []func(ctx *_context.Context) // 后中间件列表
	Parameter            map[string]string             // 正则路径解析参数
	IsRegexp             bool                          // 是否使用正则匹配
}

func (this *router) Any(rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{"ANY"}, rule, call)
}
func (this *router) Get(rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{"GET"}, rule, call)
}
func (this *router) Post(rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{"POST"}, rule, call)
}
func (this *router) Put(rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{"PUT"}, rule, call)
}
func (this *router) Delete(rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{"DELETE"}, rule, call)
}
func (this *router) Options(rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{"OPTIONS"}, rule, call)
}
func (this *router) Head(rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{"HEAD"}, rule, call)
}
func (this *router) Patch(rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{"PATCH"}, rule, call)
}
func (this *router) Method(method string, rule string, call func(ctx *_context.Context)) {
	this.MethodList([]string{method}, rule, call)
}
func (this *router) MethodList(methodList []string, rule string, call func(ctx *_context.Context)) {
	var groupMethodList []string = methodList
	var groupMiddlewareBeforeList []func(ctx *_context.Context) = []func(ctx *_context.Context){}
	var groupMiddlewareAfterList []func(ctx *_context.Context) = []func(ctx *_context.Context){}
	var groupPrefix string = ``
	for _, g := range groupList {
		groupMethodList = append(groupMethodList, g.methodList...)
		groupMiddlewareBeforeList = append(groupMiddlewareBeforeList, g.middlewareBeforeList...)
		groupMiddlewareAfterList = append(groupMiddlewareAfterList, g.middlewareAfterList...)
		groupPrefix += g.prefix
	}
	groupMethodList = append(groupMethodList, this.methodList...)
	groupMiddlewareBeforeList = append(groupMiddlewareBeforeList, this.middlewareBeforeList...)
	groupMiddlewareAfterList = append(groupMiddlewareAfterList, this.middlewareAfterList...)
	groupPrefix += this.prefix
	r := &Router{
		Call:                 call,
		MethodList:           groupMethodList,
		MiddlewareBeforeList: groupMiddlewareBeforeList,
		MiddlewareAfterList:  groupMiddlewareAfterList,
		Parameter:            map[string]string{},
		IsRegexp:             false,
	}
	rule = groupPrefix + `/` + strings.Trim(rule, `/`)
	rulePartList := []string{}
	for _, rulePart := range strings.Split(rule, `/`) {
		if _is.Empty(rulePart) {
			continue
		}
		if `:` == rulePart[0:1] {
			matchedList := regexp.MustCompile(`:(\w+)(.*)`).FindStringSubmatch(rulePart)
			if len(matchedList) > 0 {
				r.IsRegexp = true
				r.Parameter[matchedList[1]] = ""
				if _is.Empty(matchedList[2]) {
					rulePartList = append(rulePartList, `([\w-]+)`)
				} else {
					rulePartList = append(rulePartList, matchedList[2])
				}
			}
			continue
		}
		rulePartList = append(rulePartList, rulePart)
	}
	r.Rule = `/` + _list.Implode(`/`, rulePartList)
	if r.IsRegexp {
		r.Rule = `^` + r.Rule + `$`
	}
	RouterList = append(RouterList, r)
}
func (this *router) MiddlewareBefore(middleware func(ctx *_context.Context)) *router {
	return this.MiddlewareBeforeList([]func(ctx *_context.Context){middleware})
}
func (this *router) MiddlewareBeforeList(middlewareList []func(ctx *_context.Context)) *router {
	this.middlewareBeforeList = append(this.middlewareBeforeList, middlewareList...)
	return this
}
func (this *router) MiddlewareAfter(middleware func(ctx *_context.Context)) *router {
	return this.MiddlewareAfterList([]func(ctx *_context.Context){middleware})
}
func (this *router) MiddlewareAfterList(middlewareList []func(ctx *_context.Context)) *router {
	this.middlewareAfterList = append(this.middlewareAfterList, middlewareList...)
	return this
}
func (this *router) Prefix(prefix string) *router {
	prefixTrimmed := strings.Trim(prefix, `/`)
	if len(prefixTrimmed) > 0 {
		this.prefix += prefix
	}
	return this
}
func (this *router) Business(business string) *router {
	prefixTrimmed := strings.Trim(business, `/`)
	if len(prefixTrimmed) > 0 {
		this.business = business
	}
	return this
}
func (this *router) Group(group func()) {
	groupList = append(groupList, this)
	group()
	if l := len(groupList); l > 0 {
		groupList = groupList[0 : l-1]
	}
}

var groupList []*router = []*router{}
var RouterList []*Router = []*Router{}

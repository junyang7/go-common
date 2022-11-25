package _router

import (
	"github.com/junyang7/go-common/src/_context"
	"github.com/junyang7/go-common/src/_server/_router"
	"strings"
)

type router struct {
	prefix         string
	path           string
	methodList     []string
	middlewareList []func(ctx *_context.Context)
	handler        func(ctx *_context.Context)
}

var groupList = []*router{}

func Group(group func()) {
	this := &router{}
	groupList = append(groupList, this)
	group()
	if l := len(groupList); l > 0 {
		groupList = groupList[0 : l-1]
	}
}
func Prefix(prefix string) *router {
	this := &router{}
	this.prefix = prefix
	return this
}
func Middleware(middleware func(ctx *_context.Context)) *router {
	this := &router{}
	this.middlewareList = append(this.middlewareList, middleware)
	return this
}
func MiddlewareList(middlewareList []func(ctx *_context.Context)) *router {
	this := &router{}
	this.middlewareList = append(this.middlewareList, middlewareList...)
	return this
}
func Post(path string, handler func(ctx *_context.Context)) {
	this := &router{}
	this.MethodList([]string{"POST"}, path, handler)
}
func Delete(path string, handler func(ctx *_context.Context)) {
	this := &router{}
	this.MethodList([]string{"DELETE"}, path, handler)
}
func Put(path string, handler func(ctx *_context.Context)) {
	this := &router{}
	this.MethodList([]string{"PUT"}, path, handler)
}
func Get(path string, handler func(ctx *_context.Context)) {
	this := &router{}
	this.MethodList([]string{"GET"}, path, handler)
}
func Any(path string, handler func(ctx *_context.Context)) {
	this := &router{}
	this.MethodList([]string{"ANY"}, path, handler)
}
func Method(method string, path string, handler func(ctx *_context.Context)) {
	this := &router{}
	this.MethodList([]string{strings.ToUpper(method)}, path, handler)
}
func MethodList(methodList []string, path string, handler func(ctx *_context.Context)) {
	this := &router{}
	this.MethodList(methodList, path, handler)
}

func (this *router) Group(group func()) {
	groupList = append(groupList, this)
	group()
	if l := len(groupList); l > 0 {
		groupList = groupList[0 : l-1]
	}
}
func (this *router) Prefix(prefix string) *router {
	this.prefix = prefix
	return this
}
func (this *router) Middleware(middleware func(ctx *_context.Context)) *router {
	this.middlewareList = append(this.middlewareList, middleware)
	return this
}
func (this *router) MiddlewareList(middlewareList []func(ctx *_context.Context)) *router {
	this.middlewareList = append(this.middlewareList, middlewareList...)
	return this
}
func (this *router) Post(path string, handler func(ctx *_context.Context)) {
	this.MethodList([]string{"POST"}, path, handler)
}
func (this *router) Delete(path string, handler func(ctx *_context.Context)) {
	this.MethodList([]string{"DELETE"}, path, handler)
}
func (this *router) Put(path string, handler func(ctx *_context.Context)) {
	this.MethodList([]string{"PUT"}, path, handler)
}
func (this *router) Get(path string, handler func(ctx *_context.Context)) {
	this.MethodList([]string{"GET"}, path, handler)
}
func (this *router) Any(path string, handler func(ctx *_context.Context)) {
	this.MethodList([]string{"ANY"}, path, handler)
}
func (this *router) Method(method string, path string, handler func(ctx *_context.Context)) {
	this.MethodList([]string{strings.ToUpper(method)}, path, handler)
}
func (this *router) MethodList(methodList []string, path string, handler func(ctx *_context.Context)) {
	var _path string
	var _methodList []string
	var _middlewareList []func(ctx *_context.Context)
	for _, group := range groupList {
		_path += group.prefix
		_middlewareList = append(_middlewareList, group.middlewareList...)
	}
	_path += this.prefix
	_path += path
	_methodList = append(_methodList, methodList...)
	_middlewareList = append(_middlewareList, this.middlewareList...)
	_router.RouterMap[_path] = &_router.Router{
		Path:           _path,
		MethodList:     _methodList,
		MiddlewareList: _middlewareList,
		Handler:        handler,
	}
}

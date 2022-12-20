package _router

import "github.com/junyang7/go-common/src/_context"

type Router struct {
	Path                 string
	MethodList           []string
	MiddlewareBeforeList []func(ctx *_context.Context)
	MiddlewareAfterList  []func(ctx *_context.Context)
	Handler              func(ctx *_context.Context)
}

var RouterMap = map[string]*Router{}

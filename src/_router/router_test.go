package _router

import (
	"fmt"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_assert"
	"github.com/junyang7/go-common/src/_context"
	"regexp"
	"testing"
)

func testFunc(ctx *_context.Context) {
	ctx.STORE = map[string]interface{}{"test": "ok"}
}
func TestAny(t *testing.T) {
	// 普通
	{
		rule := `/router/any`
		Any(rule, testFunc)
		r := RouterList[0]
		{
			var expect string = rule
			get := r.Rule
			_assert.Equal(t, expect, get)
		}
		{
			var expect string = "func(*_context.Context)"
			get := fmt.Sprintf("%T", r.Call)
			_assert.Equal(t, expect, get)
		}
		{
			var expect []string = []string{"ANY"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
		//{
		//	var expect string = "[]func(*_context.Context)"
		//	get := fmt.Sprintf("%T", r.MiddlewareList)
		//	_assert.Equal(t, expect, get)
		//}
		{
			var expect []string = []string{}
			get := r.ParameterList
			_assert.EqualByList(t, expect, get)
		}
		{
			var expect bool = false
			get := r.IsRegexp
			_assert.Equal(t, expect, get)
		}
		{
			ctx := &_context.Context{}
			r.Call(ctx)
			var expect string = "ok"
			get := _as.String(ctx.STORE["test"])
			_assert.Equal(t, expect, get)
		}
		RouterList = []*Router{}
	}
	// 正则
	{
		rule := `/:a/:b(\w+)/c/d`
		Any(rule, testFunc)
		r := RouterList[0]
		{
			var expect bool = true
			get := r.IsRegexp
			_assert.Equal(t, expect, get)
		}
		{
			var expect string = `^/([\w-]+)/(\w+)/c/d$`
			get := r.Rule
			_assert.Equal(t, expect, get)
		}
		{
			var expect []string = []string{"a", "b"}
			get := r.ParameterList
			_assert.Equal(t, expect, get)
		}
		{
			rule := `/a/b/c/d`
			var expect []string = []string{`/a/b/c/d`, `a`, `b`}
			get := regexp.MustCompile(r.Rule).FindStringSubmatch(rule)
			_assert.EqualByList(t, expect, get)
		}
		{
			rule := `/a/b/c/d/e`
			var expect []string = []string{}
			get := regexp.MustCompile(r.Rule).FindStringSubmatch(rule)
			_assert.EqualByList(t, expect, get)
		}
		{
			rule := `a/b/c/d`
			var expect []string = []string{}
			get := regexp.MustCompile(r.Rule).FindStringSubmatch(rule)
			_assert.EqualByList(t, expect, get)
		}
		{
			rule := `/e/a/b/c/d`
			var expect []string = []string{}
			get := regexp.MustCompile(r.Rule).FindStringSubmatch(rule)
			_assert.EqualByList(t, expect, get)
		}
		RouterList = []*Router{}
	}
}
func TestGet(t *testing.T) {
	rule := `/router/any`
	Get(rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{"GET"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}
func TestPost(t *testing.T) {
	rule := `/router/post`
	Post(rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{"POST"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}
func TestPut(t *testing.T) {
	rule := `/router/put`
	Put(rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{"PUT"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}
func TestDelete(t *testing.T) {
	rule := `/router/delete`
	Delete(rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{"DELETE"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}
func TestOptions(t *testing.T) {
	rule := `/router/options`
	Options(rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{"OPTIONS"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}
func TestHead(t *testing.T) {
	rule := `/router/head`
	Head(rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{"HEAD"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}
func TestPatch(t *testing.T) {
	rule := `/router/patch`
	Patch(rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{"PATCH"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}
func TestMethod(t *testing.T) {
	rule := `/router/method`
	method := `CONNECT`
	Method(method, rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{"CONNECT"}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}
func TestMethodList(t *testing.T) {
	rule := `/router/methodList`
	methodList := []string{`CONNECT`, `TRACE`}
	MethodList(methodList, rule, testFunc)
	r := RouterList[0]
	{
		{
			var expect []string = []string{`CONNECT`, `TRACE`}
			get := r.MethodList
			_assert.EqualByList(t, expect, get)
		}
	}
	RouterList = []*Router{}
}

//	func TestMiddleware(t *testing.T) {
//		rule := `/router/middleware`
//		Middleware(testFunc).Any(rule, testFunc)
//		r := RouterList[0]
//		{
//			{
//				var expect string = "[]func(*_context.Context)"
//				get := fmt.Sprintf("%T", r.MiddlewareList)
//				_assert.Equal(t, expect, get)
//			}
//			{
//				m := r.MiddlewareList[0]
//				ctx := &_context.Context{}
//				m(ctx)
//				var expect string = "ok"
//				get := _as.String(ctx.STORE["test"])
//				_assert.Equal(t, expect, get)
//			}
//		}
//		RouterList = []*Router{}
//	}
//
//	func TestMiddlewareList(t *testing.T) {
//		rule := `/router/middlewareList`
//		middlewareList := []func(ctx *_context.Context){testFunc}
//		MiddlewareList(middlewareList).Any(rule, testFunc)
//		r := RouterList[0]
//		{
//			{
//				var expect string = "[]func(*_context.Context)"
//				get := fmt.Sprintf("%T", r.MiddlewareList)
//				_assert.Equal(t, expect, get)
//			}
//			{
//				m := r.MiddlewareList[0]
//				ctx := &_context.Context{}
//				m(ctx)
//				var expect string = "ok"
//				get := _as.String(ctx.STORE["test"])
//				_assert.Equal(t, expect, get)
//			}
//		}
//		RouterList = []*Router{}
//	}
func TestPrefix(t *testing.T) {
	{
		rule := `/router/prefix`
		Prefix(`/prefix`).Any(rule, testFunc)
		r := RouterList[0]
		var expect string = `/prefix/router/prefix`
		get := r.Rule
		_assert.Equal(t, expect, get)
		RouterList = []*Router{}
	}
	{
		rule := `/router/prefix`
		Prefix(`/`).Any(rule, testFunc)
		r := RouterList[0]
		var expect string = `/router/prefix`
		get := r.Rule
		_assert.Equal(t, expect, get)
		RouterList = []*Router{}
	}
	{
		rule := `/router/prefix`
		Prefix(`/prefix/`).Any(rule, testFunc)
		r := RouterList[0]
		var expect string = `/prefix/router/prefix`
		get := r.Rule
		_assert.Equal(t, expect, get)
		RouterList = []*Router{}
	}
	{
		rule := `/router/prefix`
		Prefix(`///prefix///`).Any(rule, testFunc)
		r := RouterList[0]
		var expect string = `/prefix/router/prefix`
		get := r.Rule
		_assert.Equal(t, expect, get)
		RouterList = []*Router{}
	}
	{
		rule := `///router/prefix///`
		Prefix(`///prefix///`).Any(rule, testFunc)
		r := RouterList[0]
		var expect string = `/prefix/router/prefix`
		get := r.Rule
		_assert.Equal(t, expect, get)
		RouterList = []*Router{}
	}
	{
		rule := ``
		Prefix(``).Any(rule, testFunc)
		r := RouterList[0]
		var expect string = `/`
		get := r.Rule
		_assert.Equal(t, expect, get)
		RouterList = []*Router{}
	}
}
func TestGroup(t *testing.T) {
	RouterList = []*Router{}
	Group(func() {
		Any("/r1", testFunc)
		Any("/r2", testFunc)
	})
	Prefix("/p1").Group(func() {
		Any("/r3", testFunc)
		Any("/r4", testFunc)
		Prefix("/p2").Group(func() {
			Any("/r5", testFunc)
			Any("/r6", testFunc)
		})
	})
	Group(func() {
		Any("/r7", testFunc)
		Any("/r8", testFunc)
	})
	Prefix("/p3").Group(func() {
		Any("/r9", testFunc)
		Any("/r10", testFunc)
		Prefix("/p4").Group(func() {
			Any("/r11", testFunc)
			Any("/r12", testFunc)
		})
		Prefix("/p5").Group(func() {
			Any("/r13", testFunc)
			Any("/r14", testFunc)
		})
		Any("/r15", testFunc)
		Any("/r16", testFunc)
	})
	var expect []string = []string{
		`/r1`,
		`/r2`,
		`/p1/r3`,
		`/p1/r4`,
		`/p1/p2/r5`,
		`/p1/p2/r6`,
		`/r7`,
		`/r8`,
		`/p3/r9`,
		`/p3/r10`,
		`/p3/p4/r11`,
		`/p3/p4/r12`,
		`/p3/p5/r13`,
		`/p3/p5/r14`,
		`/p3/r15`,
		`/p3/r16`,
	}
	get := []string{}
	for _, r := range RouterList {
		get = append(get, r.Rule)
		fmt.Println(r.Rule)
	}
	_assert.EqualByList(t, expect, get)
}

//func TestRouter_Any(t *testing.T) {
//	RouterList = []*Router{}
//	middleware := testFunc
//	Middleware(middleware).Prefix(`/p1`).Group(func() {
//		Any(`/r1`, testFunc)
//		Any(`/r2`, testFunc)
//		Middleware(middleware).Prefix(`/p2`).Group(func() {
//			Any(`/r3`, testFunc)
//			Middleware(middleware).Any(`/r4`, testFunc)
//		})
//	})
//	var expect []string = []string{
//		`/p1/r1`,
//		`/p1/r2`,
//		`/p1/p2/r3`,
//		`/p1/p2/r4`,
//	}
//	get := []string{}
//	for _, r := range RouterList {
//		get = append(get, r.Rule)
//	}
//	_assert.EqualByList(t, expect, get)
//	_assert.Equal(t, 1, len(RouterList[0].MiddlewareList))
//	_assert.Equal(t, 1, len(RouterList[1].MiddlewareList))
//	_assert.Equal(t, 2, len(RouterList[2].MiddlewareList))
//	_assert.Equal(t, 3, len(RouterList[3].MiddlewareList))
//}

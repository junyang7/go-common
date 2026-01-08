package _router

import (
	"github.com/junyang7/go-common/_context"
	"github.com/junyang7/go-common/_is"
	"github.com/junyang7/go-common/_list"
	"regexp"
	"strings"
	"sync"
)

// Manager 路由管理器 - 实例级路由管理，线程安全
type Manager struct {
	routers   []*Router
	groupList []*router
	mu        sync.RWMutex
	frozen    bool // 启动后冻结，防止运行时修改
}

// NewManager 创建路由管理器
func NewManager() *Manager {
	return &Manager{
		routers:   make([]*Router, 0, 16),
		groupList: make([]*router, 0, 4),
	}
}

// Add 添加路由（内部方法）
func (m *Manager) add(r *Router) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if m.frozen {
		panic("router is frozen, cannot add new routes after server started")
	}
	
	m.routers = append(m.routers, r)
}

// Match 匹配路由（高性能，只读锁）
func (m *Manager) Match(path string) (*Router, map[string]string) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	params := make(map[string]string)
	
	// 精确匹配优先
	for _, r := range m.routers {
		if !r.IsRegexp && r.Rule == path {
			return r, params
		}
	}
	
	// 正则匹配
	for _, r := range m.routers {
		if r.IsRegexp {
			matchedList := regexp.MustCompile(r.Rule).FindStringSubmatch(path)
			if len(matchedList) > 0 {
				// 提取参数
				for index, parameter := range r.ParameterList {
					params[parameter] = matchedList[index+1]
				}
				return r, params
			}
		}
	}
	
	return nil, params
}

// Freeze 冻结路由表（服务器启动后调用）
func (m *Manager) Freeze() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.frozen = true
}

// Count 返回路由数量（用于测试）
func (m *Manager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.routers)
}

// List 返回所有路由（用于调试）
func (m *Manager) List() []*Router {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	result := make([]*Router, len(m.routers))
	copy(result, m.routers)
	return result
}

// Builder 路由构建器（链式调用）
type Builder struct {
	manager              *Manager
	prefix               string
	middlewareBeforeList []func(ctx *_context.Context)
	middlewareAfterList  []func(ctx *_context.Context)
	methodList           []string
}

// NewBuilder 创建路由构建器
func NewBuilder(manager *Manager) *Builder {
	return &Builder{
		manager:              manager,
		middlewareBeforeList: []func(ctx *_context.Context){},
		middlewareAfterList:  []func(ctx *_context.Context){},
		methodList:           []string{},
	}
}

// Any 注册 ANY 方法路由
func (b *Builder) Any(rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{"ANY"}, rule, call)
}

// Get 注册 GET 方法路由
func (b *Builder) Get(rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{"GET"}, rule, call)
}

// Post 注册 POST 方法路由
func (b *Builder) Post(rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{"POST"}, rule, call)
}

// Put 注册 PUT 方法路由
func (b *Builder) Put(rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{"PUT"}, rule, call)
}

// Delete 注册 DELETE 方法路由
func (b *Builder) Delete(rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{"DELETE"}, rule, call)
}

// Options 注册 OPTIONS 方法路由
func (b *Builder) Options(rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{"OPTIONS"}, rule, call)
}

// Head 注册 HEAD 方法路由
func (b *Builder) Head(rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{"HEAD"}, rule, call)
}

// Patch 注册 PATCH 方法路由
func (b *Builder) Patch(rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{"PATCH"}, rule, call)
}

// Method 注册指定方法路由
func (b *Builder) Method(method string, rule string, call func(ctx *_context.Context)) {
	b.MethodList([]string{method}, rule, call)
}

// MethodList 注册多个方法路由
func (b *Builder) MethodList(methodList []string, rule string, call func(ctx *_context.Context)) {
	// 合并 Group 的前缀和中间件
	var groupMethodList []string = methodList
	var groupMiddlewareBeforeList []func(ctx *_context.Context) = []func(ctx *_context.Context){}
	var groupMiddlewareAfterList []func(ctx *_context.Context) = []func(ctx *_context.Context){}
	var groupPrefix string = ``
	
	for _, g := range b.manager.groupList {
		groupMethodList = append(groupMethodList, g.methodList...)
		groupMiddlewareBeforeList = append(groupMiddlewareBeforeList, g.middlewareBeforeList...)
		groupMiddlewareAfterList = append(groupMiddlewareAfterList, g.middlewareAfterList...)
		groupPrefix += g.prefix
	}
	
	groupMethodList = append(groupMethodList, b.methodList...)
	groupMiddlewareBeforeList = append(groupMiddlewareBeforeList, b.middlewareBeforeList...)
	groupMiddlewareAfterList = append(groupMiddlewareAfterList, b.middlewareAfterList...)
	groupPrefix += b.prefix
	
	r := &Router{
		Call:                 call,
		MethodList:           groupMethodList,
		MiddlewareBeforeList: groupMiddlewareBeforeList,
		MiddlewareAfterList:  groupMiddlewareAfterList,
		ParameterList:        []string{},
		IsRegexp:             false,
	}
	
	// 规范化路由规则
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
				r.ParameterList = append(r.ParameterList, matchedList[1])
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
	
	b.manager.add(r)
}

// MiddlewareBefore 添加前置中间件
func (b *Builder) MiddlewareBefore(middleware func(ctx *_context.Context)) *Builder {
	return b.MiddlewareBeforeList([]func(ctx *_context.Context){middleware})
}

// MiddlewareBeforeList 添加多个前置中间件
func (b *Builder) MiddlewareBeforeList(middlewareList []func(ctx *_context.Context)) *Builder {
	b.middlewareBeforeList = append(b.middlewareBeforeList, middlewareList...)
	return b
}

// MiddlewareAfter 添加后置中间件
func (b *Builder) MiddlewareAfter(middleware func(ctx *_context.Context)) *Builder {
	return b.MiddlewareAfterList([]func(ctx *_context.Context){middleware})
}

// MiddlewareAfterList 添加多个后置中间件
func (b *Builder) MiddlewareAfterList(middlewareList []func(ctx *_context.Context)) *Builder {
	b.middlewareAfterList = append(b.middlewareAfterList, middlewareList...)
	return b
}

// Prefix 设置路由前缀
func (b *Builder) Prefix(prefix string) *Builder {
	prefixTrimmed := strings.Trim(prefix, `/`)
	if len(prefixTrimmed) > 0 {
		b.prefix += prefix
	}
	return b
}

// Group 路由组
func (b *Builder) Group(group func()) {
	// 将当前 builder 转换为 router 临时对象
	r := &router{
		prefix:               b.prefix,
		middlewareBeforeList: b.middlewareBeforeList,
		middlewareAfterList:  b.middlewareAfterList,
		methodList:           b.methodList,
	}
	
	b.manager.mu.Lock()
	b.manager.groupList = append(b.manager.groupList, r)
	b.manager.mu.Unlock()
	
	group()
	
	b.manager.mu.Lock()
	if l := len(b.manager.groupList); l > 0 {
		b.manager.groupList = b.manager.groupList[0 : l-1]
	}
	b.manager.mu.Unlock()
}


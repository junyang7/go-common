package _server

import (
	"context"
	"fmt"
	"github.com/junyang7/go-common/_codeMessage"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_context"
	"github.com/junyang7/go-common/_exception"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_list"
	"github.com/junyang7/go-common/_render"
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_router"
	"github.com/junyang7/go-common/_structure"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

// API 路径前缀
const APIPrefix = "/api/"

// CORS 相关常量
var (
	DefaultCORSHeaders = []string{"content-type", "authorization"}
	DefaultCORSMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"}
)

// apiEngine API 服务器引擎
type apiEngine struct {
	*BaseEngine
	prefix      string
	corsHeaders []string
	corsMethods []string
	handler     *http.Server
}

// Api 创建 API 引擎
func Api() *apiEngine {
	return &apiEngine{
		BaseEngine:  newBaseEngine(),
		prefix:      APIPrefix,
		corsHeaders: DefaultCORSHeaders,
		corsMethods: DefaultCORSMethods,
	}
}

// Load 从配置加载（链式调用）
func (a *apiEngine) Load(conf _conf.Conf, business string) *apiEngine {
	load(conf)
	raw := _conf.Get(business).Value()
	var serverApi _structure.ServerApi
	
	if err := decodeConfig(raw, &serverApi); err != nil {
		_interceptor.Insure(false).
			Message("API配置解析失败").
			Data(map[string]interface{}{"error": err.Error()}).
			Do()
	}
	
	a.setHost(serverApi.Host)
	a.setPort(serverApi.Port)
	a.setDebug(serverApi.Debug)
	a.setOrigin(serverApi.Origin)
	
	return a
}

// Debug 设置调试模式（链式调用）
func (a *apiEngine) Debug(debug bool) *apiEngine {
	a.setDebug(debug)
	return a
}

// Network 设置网络类型（链式调用）
func (a *apiEngine) Network(network string) *apiEngine {
	a.setNetwork(network)
	return a
}

// Host 设置主机地址（链式调用）
func (a *apiEngine) Host(host string) *apiEngine {
	a.setHost(host)
	return a
}

// Port 设置端口（链式调用）
func (a *apiEngine) Port(port string) *apiEngine {
	a.setPort(port)
	return a
}

// Origin 设置跨域白名单（链式调用）
func (a *apiEngine) Origin(origin []string) *apiEngine {
	a.setOrigin(origin)
	return a
}

// Prefix 设置 API 路径前缀（链式调用）
func (a *apiEngine) Prefix(prefix string) *apiEngine {
	a.prefix = prefix
	return a
}

// CORSHeaders 设置 CORS 允许的 Headers（链式调用）
func (a *apiEngine) CORSHeaders(headers []string) *apiEngine {
	a.corsHeaders = headers
	return a
}

// CORSMethods 设置 CORS 允许的 Methods（链式调用）
func (a *apiEngine) CORSMethods(methods []string) *apiEngine {
	a.corsMethods = methods
	return a
}

// Router 添加路由（链式调用）
func (a *apiEngine) Router(router *_router.Router) *apiEngine {
	a.addRouter(router)
	return a
}

// RouterManager 设置自定义路由管理器（链式调用）
// 注意: 仅用于测试场景，生产环境使用默认全局路由即可
func (a *apiEngine) RouterManager(manager *_router.Manager) *apiEngine {
	a.setRouterManager(manager)
	return a
}

// Run 启动 API 服务器（阻塞模式，适配旧 API）
func (a *apiEngine) Run() {
	ctx := context.Background()
	if err := a.RunWithContext(ctx); err != nil && err != http.ErrServerClosed {
		_interceptor.Insure(false).Message(err).Do()
	}
}

// RunWithContext 启动 API 服务器（支持 context 控制）
func (a *apiEngine) RunWithContext(ctx context.Context) error {
	// 1. 验证配置
	if err := a.validateConfig(); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	
	// 2. 执行启动前回调
	if err := a.executeBeforeStart(); err != nil {
		return fmt.Errorf("before start callback failed: %w", err)
	}
	
	// 3. 监听端口
	if err := a.listen(ctx); err != nil {
		return err
	}
	
	// 4. 冻结路由表
	a.routerManager.Freeze()
	
	// 5. 创建 HTTP 服务器
	mux := http.NewServeMux()
	mux.HandleFunc(a.prefix, a.ServeHTTP)
	
	a.handler = &http.Server{
		Handler: mux,
	}
	
	// 6. 执行启动后回调
	a.executeAfterStart()
	
	// 7. 监听 context 取消信号（优雅关闭）
	go func() {
		<-ctx.Done()
		a.executeBeforeStop()
		
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		
		if err := a.handler.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("⚠️  Server shutdown error: %v\n", err)
		}
		
		a.shutdown()
		a.executeAfterStop()
	}()
	
	// 8. 启动服务（阻塞）
	err := a.handler.Serve(a.listener)
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}
	
	return nil
}

// ServeHTTP 处理 HTTP 请求
func (a *apiEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := &apiProcessor{
		w:           w,
		r:           r,
		origin:      a.getOrigin(),
		corsHeaders: a.corsHeaders,
		corsMethods: a.corsMethods,
		debug:       a.GetDebug(),
		manager:     a.routerManager,
	}
	p.do()
}

// apiProcessor API 请求处理器（请求级）
type apiProcessor struct {
	w           http.ResponseWriter
	r           *http.Request
	ctx         *_context.Context
	origin      []string
	corsHeaders []string
	corsMethods []string
	router      *_router.Router
	routerParams map[string]string
	debug       bool
	manager     *_router.Manager
}

// do 处理请求（总入口）
func (p *apiProcessor) do() {
	defer func() {
		if err := recover(); err != nil {
			p.exception(err)
		}
	}()
	
	// 1. 创建上下文
	p.ctx = _context.New(p.w, p.r, p.debug)
	
	// 2. 设置基础 CORS headers
	p.w.Header().Set("access-control-allow-credentials", "true")
	
	// 3. 检查 CORS
	p.checkOrigin()
	
	// 4. OPTIONS 预检请求直接返回
	if p.ctx.ServerParameter("method").String().Value() == "OPTIONS" {
		p.w.WriteHeader(http.StatusOK)
		return
	}
	
	// 5. 检查路由
	p.checkRouter()
	
	// 6. 检查请求方法
	p.checkRouterMethod()
	
	// 7. 执行前置中间件
	p.middlewareBefore()
	
	// 8. 执行业务逻辑
	p.business()
	
	// 9. 执行后置中间件
	p.middlewareAfter()
}

// checkOrigin 检查跨域（修复安全漏洞）
func (p *apiProcessor) checkOrigin() {
	originHeader := p.ctx.Header("origin").String().Value()
	if originHeader == "" {
		// 非跨域请求
		return
	}
	
	// 严格解析 Origin
	parsedOrigin, err := url.Parse(originHeader)
	if err != nil {
		_interceptor.Insure(false).
			Message("无效的 Origin 格式").
			Data(map[string]interface{}{"origin": originHeader, "error": err.Error()}).
			Do()
		return
	}
	
	// 验证协议
	if parsedOrigin.Scheme != "http" && parsedOrigin.Scheme != "https" {
		_interceptor.Insure(false).
			Message("不支持的协议").
			Data(map[string]interface{}{"origin": originHeader, "scheme": parsedOrigin.Scheme}).
			Do()
		return
	}
	
	// 检查白名单
	allowed := false
	allowCredentials := false
	
	for _, allowedOrigin := range p.origin {
		// 通配符 * - 允许所有域（但不能带 credentials）
		if allowedOrigin == "*" {
			p.w.Header().Set("access-control-allow-origin", "*")
			allowed = true
			allowCredentials = false
			break
		}
		
		// 精确匹配
		if allowedOrigin == parsedOrigin.Host {
			p.w.Header().Set("access-control-allow-origin", originHeader)
			allowed = true
			allowCredentials = true
			break
		}
		
		// 子域名匹配（以 . 开头）
		if strings.HasPrefix(allowedOrigin, ".") {
			suffix := allowedOrigin[1:] // 移除前导点
			// 严格的后缀匹配：必须是完整的域名段
			if parsedOrigin.Host == suffix || strings.HasSuffix(parsedOrigin.Host, "."+suffix) {
				p.w.Header().Set("access-control-allow-origin", originHeader)
				allowed = true
				allowCredentials = true
				break
			}
		}
	}
	
	if !allowed {
		_interceptor.Insure(false).
			Message("不支持的跨域请求").
			Data(map[string]interface{}{"origin": originHeader}).
			Do()
		return
	}
	
	// 设置 CORS headers
	if allowCredentials {
		p.w.Header().Set("access-control-allow-credentials", "true")
	}
	
	p.w.Header().Set("access-control-allow-headers", strings.Join(p.corsHeaders, ", "))
	p.w.Header().Set("access-control-expose-headers", strings.Join(p.corsHeaders, ", "))
	p.w.Header().Set("access-control-allow-methods", strings.Join(p.corsMethods, ", "))
	p.w.Header().Set("access-control-max-age", "86400") // 24小时缓存
}

// checkRouter 检查路由（使用新的路由管理器）
func (p *apiProcessor) checkRouter() {
	path := p.ctx.ServerParameter("path").String().Value()
	
	// 使用路由管理器匹配（高性能）
	router, params := p.manager.Match(path)
	
	if router == nil {
		_interceptor.Insure(false).
			Message("不支持的路由地址").
			Data(map[string]interface{}{"path": path}).
			Do()
	}
	
	if router.Call == nil {
		_interceptor.Insure(false).
			Message("路由处理方法未定义").
			Data(map[string]interface{}{"path": path}).
			Do()
	}
	
	p.router = router
	p.routerParams = params
	
	// 将路由参数注入到 ctx 中（保持兼容性）
	for key, value := range params {
		p.ctx.GET[key] = value
		p.ctx.POST[key] = value
		p.ctx.REQUEST[key] = value
	}
}

// checkRouterMethod 检查请求方法
func (p *apiProcessor) checkRouterMethod() {
	method := p.ctx.ServerParameter("method").String().Value()
	
	_interceptor.Insure(
		_list.In("ANY", p.router.MethodList) || _list.In(method, p.router.MethodList),
	).
		Message("不支持的请求方法").
		Data(map[string]interface{}{
			"method":  method,
			"allowed": p.router.MethodList,
		}).
		Do()
}

// middlewareBefore 执行前置中间件
func (p *apiProcessor) middlewareBefore() {
	for _, middleware := range p.router.MiddlewareBeforeList {
		middleware(p.ctx)
	}
}

// business 执行业务逻辑
func (p *apiProcessor) business() {
	p.router.Call(p.ctx)
}

// middlewareAfter 执行后置中间件
func (p *apiProcessor) middlewareAfter() {
	for _, middleware := range p.router.MiddlewareAfterList {
		middleware(p.ctx)
	}
}

// exception 异常处理（完善错误信息）
func (p *apiProcessor) exception(err any) {
	res := _response.New()
	
	switch e := err.(type) {
	case *_exception.Exception:
		// 业务异常
		res.Code = e.Code
		res.Message = e.Message
		res.Data = e.Data
	default:
		// 未知异常
		res.Code = _codeMessage.ErrDefault.Code
		res.Message = fmt.Sprintf("%v", err)
	}
	
	// 调试模式下添加错误文件和行号
	if p.debug {
		// 获取调用栈信息（跳过 runtime 内部调用）
		pcs := make([]uintptr, 10)
		n := runtime.Callers(3, pcs) // 跳过 3 层：Callers, exception, defer
		if n > 0 {
			frames := runtime.CallersFrames(pcs[:n])
			frame, _ := frames.Next()
			res.File = frame.File
			res.Line = frame.Line
		}
	}
	
	_render.JSON(p.w, res)
}


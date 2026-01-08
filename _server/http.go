package _server

import (
	"context"
	"fmt"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_router"
	"github.com/junyang7/go-common/_structure"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// httpEngine HTTP 混合服务器引擎（API + 静态文件 + SPA）
type httpEngine struct {
	*BaseEngine
	root        string
	prefix      string
	corsHeaders []string
	corsMethods []string
	handler     *http.Server
}

// Http 创建 HTTP 引擎
func Http() *httpEngine {
	return &httpEngine{
		BaseEngine:  newBaseEngine(),
		prefix:      APIPrefix,
		corsHeaders: DefaultCORSHeaders,
		corsMethods: DefaultCORSMethods,
	}
}

// Load 从配置加载（链式调用）
func (h *httpEngine) Load(conf _conf.Conf, business string) *httpEngine {
	load(conf)
	raw := _conf.Get(business).Value()
	var serverHttp _structure.ServerHttp
	
	if err := decodeConfig(raw, &serverHttp); err != nil {
		_interceptor.Insure(false).
			Message("HTTP配置解析失败").
			Data(map[string]interface{}{"error": err.Error()}).
			Do()
	}
	
	h.setHost(serverHttp.Host)
	h.setPort(serverHttp.Port)
	h.setDebug(serverHttp.Debug)
	h.setOrigin(serverHttp.Origin)
	h.root = serverHttp.Root
	
	return h
}

// Debug 设置调试模式（链式调用）
func (h *httpEngine) Debug(debug bool) *httpEngine {
	h.setDebug(debug)
	return h
}

// Network 设置网络类型（链式调用）
func (h *httpEngine) Network(network string) *httpEngine {
	h.setNetwork(network)
	return h
}

// Host 设置主机地址（链式调用）
func (h *httpEngine) Host(host string) *httpEngine {
	h.setHost(host)
	return h
}

// Port 设置端口（链式调用）
func (h *httpEngine) Port(port string) *httpEngine {
	h.setPort(port)
	return h
}

// Origin 设置跨域白名单（链式调用）
func (h *httpEngine) Origin(origin []string) *httpEngine {
	h.setOrigin(origin)
	return h
}

// Root 设置静态文件根目录（链式调用）
func (h *httpEngine) Root(root string) *httpEngine {
	h.root = root
	return h
}

// Prefix 设置 API 路径前缀（链式调用）
func (h *httpEngine) Prefix(prefix string) *httpEngine {
	h.prefix = prefix
	return h
}

// Router 添加路由（链式调用）
func (h *httpEngine) Router(router *_router.Router) *httpEngine {
	h.addRouter(router)
	return h
}

// RouterManager 设置自定义路由管理器（链式调用）
// 注意: 仅用于测试场景，生产环境使用默认全局路由即可
func (h *httpEngine) RouterManager(manager *_router.Manager) *httpEngine {
	h.setRouterManager(manager)
	return h
}

// getRoot 获取静态文件根目录
func (h *httpEngine) getRoot() string {
	return h.root
}

// validateConfig 验证 HTTP 配置
func (h *httpEngine) validateConfig() error {
	if err := h.BaseEngine.validateConfig(); err != nil {
		return err
	}
	
	// 验证 root 路径
	if h.root == "" {
		return fmt.Errorf("root directory cannot be empty")
	}
	
	// 检查目录是否存在
	if info, err := os.Stat(h.root); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("root directory does not exist: %s", h.root)
		}
		return fmt.Errorf("cannot access root directory: %w", err)
	} else if !info.IsDir() {
		return fmt.Errorf("root path is not a directory: %s", h.root)
	}
	
	// 转换为绝对路径
	absRoot, err := filepath.Abs(h.root)
	if err != nil {
		return fmt.Errorf("cannot resolve absolute path: %w", err)
	}
	h.root = absRoot
	
	return nil
}

// Run 启动 HTTP 服务器（阻塞模式，适配旧 API）
func (h *httpEngine) Run() {
	ctx := context.Background()
	if err := h.RunWithContext(ctx); err != nil && err != http.ErrServerClosed {
		_interceptor.Insure(false).Message(err).Do()
	}
}

// RunWithContext 启动 HTTP 服务器（支持 context 控制）
func (h *httpEngine) RunWithContext(ctx context.Context) error {
	// 1. 验证配置
	if err := h.validateConfig(); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	
	// 2. 执行启动前回调
	if err := h.executeBeforeStart(); err != nil {
		return fmt.Errorf("before start callback failed: %w", err)
	}
	
	// 3. 监听端口
	if err := h.listen(ctx); err != nil {
		return err
	}
	
	// 4. 冻结路由表
	h.routerManager.Freeze()
	
	// 5. 创建 HTTP 服务器
	mux := http.NewServeMux()
	
	// API 路由（使用完整的 API 处理器）
	apiEngine := Api().
		Debug(h.GetDebug()).
		Host(h.getHost()).
		Port(h.getPort()).
		Origin(h.getOrigin()).
		Prefix(h.prefix)
	
	// 复制路由到 API 引擎
	apiEngine.routerManager = h.routerManager
	
	mux.HandleFunc(h.prefix, apiEngine.ServeHTTP)
	
	// 静态文件路由（带安全检查）
	mux.HandleFunc("/", h.serveStaticFiles)
	
	h.handler = &http.Server{
		Handler: mux,
	}
	
	// 6. 执行启动后回调
	h.executeAfterStart()
	
	// 7. 监听 context 取消信号（优雅关闭）
	go func() {
		<-ctx.Done()
		h.executeBeforeStop()
		
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		
		if err := h.handler.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("⚠️  Server shutdown error: %v\n", err)
		}
		
		h.shutdown()
		h.executeAfterStop()
	}()
	
	// 8. 启动服务（阻塞）
	err := h.handler.Serve(h.listener)
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}
	
	return nil
}

// serveStaticFiles 提供静态文件服务（带完整安全检查）
func (h *httpEngine) serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	root := h.getRoot()
	
	// 1. 清理和规范化请求路径
	requestPath := filepath.Clean("/" + r.URL.Path)
	
	// 2. 拼接完整路径
	fullPath := filepath.Join(root, requestPath)
	
	// 3. 防止路径穿越攻击（第一层检查）
	rel, err := filepath.Rel(root, fullPath)
	if err != nil || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		http.Error(w, "403 Forbidden - Path traversal detected", http.StatusForbidden)
		if h.GetDebug() {
			fmt.Printf("⚠️  Path traversal attempt: %s -> %s\n", requestPath, fullPath)
		}
		return
	}
	
	// 4. 检查符号链接（防止绕过）
	realPath, err := filepath.EvalSymlinks(fullPath)
	if err == nil {
		// 符号链接存在，验证真实路径是否在 root 内
		rel, err := filepath.Rel(root, realPath)
		if err != nil || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			http.Error(w, "403 Forbidden - Symlink escape detected", http.StatusForbidden)
			if h.GetDebug() {
				fmt.Printf("⚠️  Symlink escape attempt: %s -> %s -> %s\n", requestPath, fullPath, realPath)
			}
			return
		}
		fullPath = realPath
	}
	
	// 5. 检查文件状态
	info, err := os.Stat(fullPath)
	
	// 5.1 文件存在且不是目录
	if err == nil && !info.IsDir() {
		http.ServeFile(w, r, fullPath)
		return
	}
	
	// 5.2 目录：拒绝访问（安全策略）
	if err == nil && info.IsDir() {
		// 不自动提供目录索引，防止信息泄露
		http.Error(w, "403 Forbidden - Directory listing denied", http.StatusForbidden)
		return
	}
	
	// 6. 文件不存在：SPA fallback（仅对非静态资源路径）
	if err != nil && os.IsNotExist(err) {
		// 静态资源路径（常见前缀）直接返回 404
		staticPrefixes := []string{"/static/", "/assets/", "/js/", "/css/", "/img/", "/images/", "/fonts/"}
		for _, prefix := range staticPrefixes {
			if strings.HasPrefix(requestPath, prefix) {
				http.NotFound(w, r)
				return
			}
		}
		
		// 静态文件扩展名直接返回 404
		staticExtensions := []string{".js", ".css", ".png", ".jpg", ".jpeg", ".gif", ".svg", ".ico", ".woff", ".woff2", ".ttf", ".eot"}
		for _, ext := range staticExtensions {
			if strings.HasSuffix(requestPath, ext) {
				http.NotFound(w, r)
				return
			}
		}
		
		// 其他路径：尝试返回 index.html（支持 SPA History 模式）
		indexPath := filepath.Join(root, "index.html")
		if _, err := os.Stat(indexPath); err == nil {
			http.ServeFile(w, r, indexPath)
			return
		}
		
		// index.html 也不存在，返回 404
		http.NotFound(w, r)
		return
	}
	
	// 7. 其他错误
	http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	if h.GetDebug() {
		fmt.Printf("⚠️  File access error: %v\n", err)
	}
}


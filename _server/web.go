package _server

import (
	"context"
	"fmt"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_structure"
	"net/http"
	"os"
	"path/filepath"
)

// webEngine Web 静态文件服务器引擎
type webEngine struct {
	*BaseEngine
	root string
}

// Web 创建 Web 引擎
func Web() *webEngine {
	return &webEngine{
		BaseEngine: newBaseEngine(),
	}
}

// Load 从配置加载（链式调用）
func (w *webEngine) Load(conf _conf.Conf, business string) *webEngine {
	load(conf)
	raw := _conf.Get(business).Value()
	var serverWeb _structure.ServerWeb
	
	// 使用更高效的类型断言或 mapstructure
	if err := decodeConfig(raw, &serverWeb); err != nil {
		_interceptor.Insure(false).
			Message("Web配置解析失败").
			Data(map[string]interface{}{"error": err.Error()}).
			Do()
	}
	
	w.setHost(serverWeb.Host)
	w.setPort(serverWeb.Port)
	w.setDebug(serverWeb.Debug)
	w.setOrigin(serverWeb.Origin)
	w.root = serverWeb.Root
	
	return w
}

// Debug 设置调试模式（链式调用）
func (w *webEngine) Debug(debug bool) *webEngine {
	w.setDebug(debug)
	return w
}

// Network 设置网络类型（链式调用）
func (w *webEngine) Network(network string) *webEngine {
	w.setNetwork(network)
	return w
}

// Host 设置主机地址（链式调用）
func (w *webEngine) Host(host string) *webEngine {
	w.setHost(host)
	return w
}

// Port 设置端口（链式调用）
func (w *webEngine) Port(port string) *webEngine {
	w.setPort(port)
	return w
}

// Origin 设置跨域白名单（链式调用）
func (w *webEngine) Origin(origin []string) *webEngine {
	w.setOrigin(origin)
	return w
}

// Root 设置静态文件根目录（链式调用）
func (w *webEngine) Root(root string) *webEngine {
	w.root = root
	return w
}

// RouterManager 设置自定义路由管理器（链式调用）
// 注意: 仅用于测试场景，生产环境使用默认全局路由即可
func (w *webEngine) RouterManager(manager *_router.Manager) *webEngine {
	w.setRouterManager(manager)
	return w
}

// getRoot 获取静态文件根目录
func (w *webEngine) getRoot() string {
	return w.root
}

// validateConfig 验证 Web 配置
func (w *webEngine) validateConfig() error {
	if err := w.BaseEngine.validateConfig(); err != nil {
		return err
	}
	
	// 验证 root 路径
	if w.root == "" {
		return fmt.Errorf("root directory cannot be empty")
	}
	
	// 检查目录是否存在
	if info, err := os.Stat(w.root); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("root directory does not exist: %s", w.root)
		}
		return fmt.Errorf("cannot access root directory: %w", err)
	} else if !info.IsDir() {
		return fmt.Errorf("root path is not a directory: %s", w.root)
	}
	
	// 转换为绝对路径
	absRoot, err := filepath.Abs(w.root)
	if err != nil {
		return fmt.Errorf("cannot resolve absolute path: %w", err)
	}
	w.root = absRoot
	
	return nil
}

// Run 启动 Web 服务器（阻塞模式，适配旧 API）
func (w *webEngine) Run() {
	ctx := context.Background()
	if err := w.RunWithContext(ctx); err != nil && err != http.ErrServerClosed {
		_interceptor.Insure(false).Message(err).Do()
	}
}

// RunWithContext 启动 Web 服务器（支持 context 控制）
func (w *webEngine) RunWithContext(ctx context.Context) error {
	// 1. 验证配置
	if err := w.validateConfig(); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	
	// 2. 执行启动前回调
	if err := w.executeBeforeStart(); err != nil {
		return fmt.Errorf("before start callback failed: %w", err)
	}
	
	// 3. 监听端口
	if err := w.listen(ctx); err != nil {
		return err
	}
	
	// 4. 冻结路由表
	w.routerManager.Freeze()
	
	// 5. 创建 HTTP 服务器
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(w.getRoot())))
	
	server := &http.Server{
		Handler: mux,
	}
	
	// 6. 执行启动后回调
	w.executeAfterStart()
	
	// 7. 监听 context 取消信号（优雅关闭）
	go func() {
		<-ctx.Done()
		w.executeBeforeStop()
		
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		
		if err := server.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("⚠️  Server shutdown error: %v\n", err)
		}
		
		w.shutdown()
		w.executeAfterStop()
	}()
	
	// 8. 启动服务（阻塞）
	err := server.Serve(w.listener)
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}
	
	return nil
}

// decodeConfig 配置解码辅助函数
func decodeConfig(raw interface{}, target interface{}) error {
	// 使用 JSON 编解码实现类型转换（临时方案，后续可优化为 mapstructure）
	encoded := _json.Encode(raw)
	return _json.Decode(encoded, target)
}


package _server

import (
	"context"
	"fmt"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_router"
	"net"
	"sync"
)

// 常量定义
const (
	DefaultNetwork = "tcp"
	DefaultHost    = "0.0.0.0"
	DefaultPort    = "0"
)

// BaseEngine 服务器基础引擎（所有引擎的公共基类）
type BaseEngine struct {
	debug         bool
	network       string
	host          string
	port          string
	origin        []string
	routerManager *_router.Manager
	listener      net.Listener
	mu            sync.RWMutex
	started       bool
	
	// 回调函数
	onBeforeStart func() error
	onAfterStart  func()
	onBeforeStop  func()
	onAfterStop   func()
}

// newBaseEngine 创建基础引擎（默认使用全局路由管理器）
func newBaseEngine() *BaseEngine {
	return &BaseEngine{
		debug:         false,
		network:       DefaultNetwork,
		host:          DefaultHost,
		port:          DefaultPort,
		origin:        []string{},
		routerManager: _router.GetDefaultManager(), // 默认使用全局管理器，兼容 init() 路由注册
	}
}

// setRouterManager 设置自定义路由管理器
// 注意: 使用自定义管理器后，init() 中注册的全局路由将不可用
// 主要用于测试场景的路由隔离，生产环境建议使用默认的全局路由
func (b *BaseEngine) setRouterManager(manager *_router.Manager) {
	if manager != nil {
		b.routerManager = manager
	}
}

// Debug 设置调试模式
func (b *BaseEngine) setDebug(debug bool) {
	b.debug = debug
}

// GetDebug 获取调试模式
func (b *BaseEngine) GetDebug() bool {
	return b.debug
}

// setNetwork 设置网络类型
func (b *BaseEngine) setNetwork(network string) {
	b.network = network
}

// getNetwork 获取网络类型（带默认值）
func (b *BaseEngine) getNetwork() string {
	if b.network != "" {
		return b.network
	}
	return DefaultNetwork
}

// setHost 设置主机地址
func (b *BaseEngine) setHost(host string) {
	b.host = host
}

// getHost 获取主机地址（带默认值）
func (b *BaseEngine) getHost() string {
	if b.host != "" {
		return b.host
	}
	return DefaultHost
}

// setPort 设置端口
func (b *BaseEngine) setPort(port string) {
	b.port = port
}

// getPort 获取端口（带默认值）
func (b *BaseEngine) getPort() string {
	if b.port != "" {
		return b.port
	}
	return DefaultPort
}

// setOrigin 设置跨域白名单
func (b *BaseEngine) setOrigin(origin []string) {
	b.origin = origin
}

// getOrigin 获取跨域白名单
func (b *BaseEngine) getOrigin() []string {
	return b.origin
}

// getAddr 获取完整地址
func (b *BaseEngine) getAddr() string {
	return fmt.Sprintf("%s:%s", b.getHost(), b.getPort())
}

// GetRouterManager 获取路由管理器
func (b *BaseEngine) GetRouterManager() *_router.Manager {
	return b.routerManager
}

// AddRouter 添加路由（实例级）
func (b *BaseEngine) addRouter(router *_router.Router) {
	if router == nil {
		_interceptor.Insure(false).Message("router cannot be nil").Do()
	}
	b.routerManager.add(router)
}

// listen 监听端口（统一实现）
func (b *BaseEngine) listen(ctx context.Context) error {
	listener, err := net.Listen(b.getNetwork(), b.getAddr())
	if err != nil {
		return fmt.Errorf("listen failed on %s: %w", b.getAddr(), err)
	}
	
	b.mu.Lock()
	b.listener = listener
	b.started = true
	b.mu.Unlock()
	
	fmt.Printf("🚀 Server is running on: %s\n", listener.Addr().String())
	
	return nil
}

// GetAddr 获取实际监听地址（启动后可用）
func (b *BaseEngine) GetAddr() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	if b.listener != nil {
		return b.listener.Addr().String()
	}
	return b.getAddr()
}

// IsStarted 是否已启动
func (b *BaseEngine) IsStarted() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.started
}

// shutdown 优雅关闭（统一实现）
func (b *BaseEngine) shutdown() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if !b.started {
		return nil
	}
	
	if b.listener != nil {
		if err := b.listener.Close(); err != nil {
			return fmt.Errorf("close listener failed: %w", err)
		}
		b.listener = nil
	}
	
	b.started = false
	fmt.Println("✅ Server stopped gracefully.")
	
	return nil
}

// validateConfig 验证配置（子类可重写）
func (b *BaseEngine) validateConfig() error {
	// 基础验证
	if b.network != "tcp" && b.network != "tcp4" && b.network != "tcp6" && b.network != "unix" {
		return fmt.Errorf("invalid network type: %s", b.network)
	}
	
	// 子类可以添加更多验证
	return nil
}

// SetBeforeStartCallback 设置启动前回调
func (b *BaseEngine) SetBeforeStartCallback(callback func() error) {
	b.onBeforeStart = callback
}

// SetAfterStartCallback 设置启动后回调
func (b *BaseEngine) SetAfterStartCallback(callback func()) {
	b.onAfterStart = callback
}

// SetBeforeStopCallback 设置停止前回调
func (b *BaseEngine) SetBeforeStopCallback(callback func()) {
	b.onBeforeStop = callback
}

// SetAfterStopCallback 设置停止后回调
func (b *BaseEngine) SetAfterStopCallback(callback func()) {
	b.onAfterStop = callback
}

// executeCallbacks 执行回调
func (b *BaseEngine) executeBeforeStart() error {
	if b.onBeforeStart != nil {
		return b.onBeforeStart()
	}
	return nil
}

func (b *BaseEngine) executeAfterStart() {
	if b.onAfterStart != nil {
		b.onAfterStart()
	}
}

func (b *BaseEngine) executeBeforeStop() {
	if b.onBeforeStop != nil {
		b.onBeforeStop()
	}
}

func (b *BaseEngine) executeAfterStop() {
	if b.onAfterStop != nil {
		b.onAfterStop()
	}
}


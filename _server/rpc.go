package _server

import (
	"context"
	"fmt"
	"github.com/junyang7/go-common/_codeMessage"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_exception"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_pb"
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_router"
	"google.golang.org/grpc"
	"runtime"
)

// rpcEngine RPC 服务器引擎（gRPC）
type rpcEngine struct {
	*BaseEngine
	conf   _conf.Conf
	server *grpc.Server
}

// Rpc 创建 RPC 引擎
func Rpc() *rpcEngine {
	return &rpcEngine{
		BaseEngine: newBaseEngine(),
	}
}

// Conf 设置配置（链式调用）
func (rpc *rpcEngine) Conf(conf _conf.Conf) *rpcEngine {
	rpc.conf = conf
	return rpc
}

// Debug 设置调试模式（链式调用）
func (rpc *rpcEngine) Debug(debug bool) *rpcEngine {
	rpc.setDebug(debug)
	return rpc
}

// Network 设置网络类型（链式调用）
func (rpc *rpcEngine) Network(network string) *rpcEngine {
	rpc.setNetwork(network)
	return rpc
}

// Addr 设置地址（链式调用）
func (rpc *rpcEngine) Addr(addr string) *rpcEngine {
	// 解析 addr 为 host:port
	// 简化实现：假设 addr 格式为 "host:port"
	// TODO: 可以改进为更严格的解析
	rpc.setHost(addr)
	return rpc
}

// Router 添加路由（链式调用）
func (rpc *rpcEngine) Router(router *_router.Router) *rpcEngine {
	rpc.addRouter(router)
	return rpc
}

// RouterManager 设置自定义路由管理器（链式调用）
// 注意: 仅用于测试场景，生产环境使用默认全局路由即可
func (rpc *rpcEngine) RouterManager(manager *_router.Manager) *rpcEngine {
	rpc.setRouterManager(manager)
	return rpc
}

// Run 启动 RPC 服务器（阻塞模式，适配旧 API）
func (rpc *rpcEngine) Run() {
	ctx := context.Background()
	if err := rpc.RunWithContext(ctx); err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
}

// RunWithContext 启动 RPC 服务器（支持 context 控制）
func (rpc *rpcEngine) RunWithContext(ctx context.Context) error {
	// 1. 验证配置
	if err := rpc.validateConfig(); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	
	// 2. 执行启动前回调
	if err := rpc.executeBeforeStart(); err != nil {
		return fmt.Errorf("before start callback failed: %w", err)
	}
	
	// 3. 监听端口
	if err := rpc.listen(ctx); err != nil {
		return err
	}
	
	// 4. 冻结路由表
	rpc.routerManager.Freeze()
	
	// 5. 创建 gRPC 服务器
	rpc.server = grpc.NewServer()
	
	// 注册 RPC 服务
	_pb.RegisterServiceServer(rpc.server, &rpcCallHandler{
		manager: rpc.routerManager,
		debug:   rpc.GetDebug(),
	})
	
	// 6. 执行启动后回调
	rpc.executeAfterStart()
	
	// 7. 监听 context 取消信号（优雅关闭）
	go func() {
		<-ctx.Done()
		rpc.executeBeforeStop()
		
		// gRPC 优雅关闭
		rpc.server.GracefulStop()
		
		rpc.shutdown()
		rpc.executeAfterStop()
	}()
	
	// 8. 启动服务（阻塞）
	if err := rpc.server.Serve(rpc.listener); err != nil {
		return fmt.Errorf("server error: %w", err)
	}
	
	return nil
}

// rpcCallHandler RPC 调用处理器
type rpcCallHandler struct {
	_pb.UnimplementedServiceServer
	manager *_router.Manager
	debug   bool
}

// Call 处理 RPC 调用
func (h *rpcCallHandler) Call(ctx context.Context, req *_pb.Request) (oRes *_pb.Response, oErr error) {
	res := _response.New()
	
	defer func() {
		if err := recover(); err != nil {
			// 异常处理
			switch e := err.(type) {
			case *_exception.Exception:
				res.Code = e.Code
				res.Message = e.Message
				res.Data = e.Data
			default:
				res.Code = _codeMessage.ErrDefault.Code
				res.Message = fmt.Sprintf("%v", err)
			}
			
			// 调试模式下添加错误信息
			if h.debug {
				pcs := make([]uintptr, 10)
				n := runtime.Callers(3, pcs)
				if n > 0 {
					frames := runtime.CallersFrames(pcs[:n])
					frame, _ := frames.Next()
					res.File = frame.File
					res.Line = frame.Line
				}
			}
			
			oRes = &_pb.Response{Response: _json.Encode(res)}
		}
	}()
	
	// TODO: 实现完整的 RPC 路由匹配和调用逻辑
	// 这里需要：
	// 1. 从 req 中解析路由路径
	// 2. 使用 manager.Match() 匹配路由
	// 3. 创建 Context 对象
	// 4. 执行中间件和业务逻辑
	// 5. 返回结果
	
	// 当前为占位实现
	res.Code = _codeMessage.ErrNone.Code
	res.Message = _codeMessage.ErrNone.Message
	res.Data = map[string]string{"status": "RPC handler not fully implemented"}
	
	oRes = &_pb.Response{Response: _json.Encode(res)}
	return oRes, nil
}


// Package _server 提供多种服务器引擎实现
//
// 本包已重构为模块化架构，提供以下服务器类型：
//   - Web(): 静态文件服务器
//   - Api(): RESTful API 服务器
//   - Http(): 混合服务器（API + 静态文件 + SPA）
//   - Rpc(): gRPC 服务器
//   - Cli(): 命令行工具（待实现）
//   - Job(): 定时任务引擎（待实现）
//   - File(): 文件处理引擎（待实现）
//   - Websocket(): WebSocket 服务器（待实现）
//
// 重构亮点：
//   ✅ 消除全局变量污染 - 实例级路由管理
//   ✅ 修复 CORS 安全漏洞 - 严格的 Origin 验证
//   ✅ 修复路径遍历漏洞 - 符号链接检查
//   ✅ 优雅关闭支持 - context.Context 控制
//   ✅ 平滑启动机制 - 回调钩子
//   ✅ 配置验证 - 启动前检查
//   ✅ 高性能路由 - 精确匹配优先
//   ✅ 线程安全 - 路由表冻结
//
// 使用示例：
//
//	// 方式1: 阻塞模式（兼容旧 API）
//	_server.Api().
//	    Host("0.0.0.0").
//	    Port("8080").
//	    Origin([]string{"localhost", ".example.com"}).
//	    Router(myRouter).
//	    Run()
//
//	// 方式2: Context 控制（推荐）
//	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
//	defer cancel()
//
//	server := _server.Api().
//	    Host("0.0.0.0").
//	    Port("8080").
//	    Origin([]string{"localhost"}).
//	    Router(myRouter)
//
//	if err := server.RunWithContext(ctx); err != nil {
//	    log.Fatal(err)
//	}
//
// 平滑启动示例：
//
//	server := _server.Api().Host("0.0.0.0").Port("8080")
//
//	// 设置启动回调
//	server.SetBeforeStartCallback(func() error {
//	    fmt.Println("准备启动服务器...")
//	    return nil
//	})
//
//	server.SetAfterStartCallback(func() {
//	    fmt.Println("服务器已启动！")
//	})
//
//	server.SetBeforeStopCallback(func() {
//	    fmt.Println("正在关闭服务器...")
//	})
//
//	server.SetAfterStopCallback(func() {
//	    fmt.Println("服务器已关闭！")
//	})
//
//	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
//	defer cancel()
//
//	server.RunWithContext(ctx)
//
package _server


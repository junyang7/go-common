package _cmd

import (
	"bytes"
	"context"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_interceptor"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Result 命令执行结果
type Result struct {
	Stdout   []byte // 标准输出
	Stderr   []byte // 标准错误
	ExitCode int    // 退出码
	Error    error  // 错误信息
}

// ============================================================
// 简单方式：直接调用函数（适合 90% 的场景）
// 默认：失败就抛出（panic）
// ============================================================

// Execute 执行命令并返回标准输出（失败抛出）
func Execute(name string, arg ...string) []byte {
	cmd := exec.Command(name, arg...)
	b, err := cmd.Output()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return b
}

// ExecuteAsInt64 执行命令并将输出转换为 int64（失败抛出）
func ExecuteAsInt64(name string, arg ...string) int64 {
	return _as.Int64(Execute(name, arg...))
}

// ExecuteAsString 执行命令并将输出转换为 string（失败抛出）
func ExecuteAsString(name string, arg ...string) string {
	return _as.String(Execute(name, arg...))
}

// ExecuteInteractive 执行命令并将标准输出/错误重定向到当前进程（失败抛出）
// 适用于需要实时查看输出的场景
func ExecuteInteractive(cmd string) {
	handler := exec.Command("/bin/bash", "-c", cmd)
	handler.Stdout = os.Stdout
	handler.Stderr = os.Stderr
	if err := handler.Run(); err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
}

// ExecuteByStd 已废弃，使用 ExecuteInteractive 代替
// Deprecated: Use ExecuteInteractive instead
func ExecuteByStd(cmd string) {
	ExecuteInteractive(cmd)
}

// ExecuteSafe 安全执行命令，返回详细结果（不会抛出）
// 适用于：需要手动处理错误的少数场景
func ExecuteSafe(name string, arg ...string) *Result {
	cmd := exec.Command(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			exitCode = -1
		}
	}

	return &Result{
		Stdout:   stdout.Bytes(),
		Stderr:   stderr.Bytes(),
		ExitCode: exitCode,
		Error:    err,
	}
}

// ============================================================
// 复杂方式：构建器模式（适合需要精细控制的场景）
// 默认：失败就抛出（panic）
// ============================================================

// Command 命令构建器，支持链式配置
type Command struct {
	name    string
	args    []string
	dir     string
	env     []string
	stdin   io.Reader
	stdout  io.Writer
	stderr  io.Writer
	timeout time.Duration
	ctx     context.Context
}

// New 创建一个新的命令构建器（复杂场景使用）
func New(name string, arg ...string) *Command {
	return &Command{
		name: name,
		args: arg,
		ctx:  context.Background(),
	}
}

// Dir 设置工作目录
func (c *Command) Dir(dir string) *Command {
	c.dir = dir
	return c
}

// Env 设置环境变量（完全替换）
func (c *Command) Env(env []string) *Command {
	c.env = env
	return c
}

// AddEnv 添加环境变量（在现有基础上添加）
func (c *Command) AddEnv(key, value string) *Command {
	if c.env == nil {
		c.env = os.Environ()
	}
	c.env = append(c.env, key+"="+value)
	return c
}

// Stdin 设置标准输入
func (c *Command) Stdin(r io.Reader) *Command {
	c.stdin = r
	return c
}

// StdinString 设置标准输入（字符串）
func (c *Command) StdinString(s string) *Command {
	c.stdin = strings.NewReader(s)
	return c
}

// Stdout 设置标准输出
func (c *Command) Stdout(w io.Writer) *Command {
	c.stdout = w
	return c
}

// Stderr 设置标准错误输出
func (c *Command) Stderr(w io.Writer) *Command {
	c.stderr = w
	return c
}

// Timeout 设置超时时间
func (c *Command) Timeout(d time.Duration) *Command {
	c.timeout = d
	return c
}

// Context 设置 context
func (c *Command) Context(ctx context.Context) *Command {
	c.ctx = ctx
	return c
}

// Run 执行命令（不捕获输出，失败抛出）⭐ 默认行为
func (c *Command) Run() {
	result := c.runInternal()
	if !result.Success() {
		_interceptor.Insure(false).Message(result.Error).Do()
	}
}

// Output 执行命令并捕获输出（失败抛出）⭐ 默认行为
func (c *Command) Output() []byte {
	result := c.outputInternal()
	if !result.Success() {
		_interceptor.Insure(false).Message(result.Error).Do()
	}
	return result.Stdout
}

// OutputAsString 执行命令并返回字符串输出（失败抛出）
func (c *Command) OutputAsString() string {
	return string(c.Output())
}

// OutputAsInt64 执行命令并返回 int64 输出（失败抛出）
func (c *Command) OutputAsInt64() int64 {
	return _as.Int64(c.Output())
}

// RunSafe 安全执行命令（不捕获输出，不抛出）
// 适用于：需要手动处理错误的少数场景
func (c *Command) RunSafe() *Result {
	return c.runInternal()
}

// OutputSafe 安全执行命令并捕获输出（不抛出）
// 适用于：需要手动处理错误的少数场景
func (c *Command) OutputSafe() *Result {
	return c.outputInternal()
}

// runInternal 内部执行方法（不捕获输出）
func (c *Command) runInternal() *Result {
	cmd := c.build()

	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			exitCode = -1
		}
	}

	return &Result{
		Stdout:   nil,
		Stderr:   nil,
		ExitCode: exitCode,
		Error:    err,
	}
}

// outputInternal 内部执行方法（捕获输出）
func (c *Command) outputInternal() *Result {
	cmd := c.build()

	var stdout, stderr bytes.Buffer
	if c.stdout == nil {
		cmd.Stdout = &stdout
	}
	if c.stderr == nil {
		cmd.Stderr = &stderr
	}

	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			exitCode = -1
		}
	}

	return &Result{
		Stdout:   stdout.Bytes(),
		Stderr:   stderr.Bytes(),
		ExitCode: exitCode,
		Error:    err,
	}
}

// build 构建 exec.Cmd
func (c *Command) build() *exec.Cmd {
	var cmd *exec.Cmd

	// 应用超时或 context
	if c.timeout > 0 {
		ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
		_ = cancel // 避免 unused 警告
		cmd = exec.CommandContext(ctx, c.name, c.args...)
	} else if c.ctx != nil {
		cmd = exec.CommandContext(c.ctx, c.name, c.args...)
	} else {
		cmd = exec.Command(c.name, c.args...)
	}

	// 应用配置
	if c.dir != "" {
		cmd.Dir = c.dir
	}
	if c.env != nil {
		cmd.Env = c.env
	}
	if c.stdin != nil {
		cmd.Stdin = c.stdin
	}
	if c.stdout != nil {
		cmd.Stdout = c.stdout
	}
	if c.stderr != nil {
		cmd.Stderr = c.stderr
	}

	return cmd
}

// ============================================================
// 便捷函数（快捷方式）- 默认失败抛出
// ============================================================

// ExecuteWithTimeout 带超时的命令执行（失败抛出）
func ExecuteWithTimeout(timeout time.Duration, name string, arg ...string) []byte {
	return New(name, arg...).Timeout(timeout).Output()
}

// ExecuteWithContext 使用 context 执行命令（失败抛出）
func ExecuteWithContext(ctx context.Context, name string, arg ...string) []byte {
	return New(name, arg...).Context(ctx).Output()
}

// ExecuteWithDir 在指定目录执行命令（失败抛出）
func ExecuteWithDir(dir string, name string, arg ...string) []byte {
	return New(name, arg...).Dir(dir).Output()
}

// ExecuteWithEnv 使用自定义环境变量执行命令（失败抛出）
func ExecuteWithEnv(env []string, name string, arg ...string) []byte {
	return New(name, arg...).Env(env).Output()
}

// Pipe 创建管道命令（将第一个命令的输出作为第二个命令的输入，失败抛出）
func Pipe(cmd1, cmd2 *Command) []byte {
	// 执行第一个命令
	output1 := cmd1.Output()
	
	// 将第一个命令的输出作为第二个命令的输入
	cmd2.Stdin(bytes.NewReader(output1))
	return cmd2.Output()
}

// PipeSafe 安全管道命令（不抛出）
func PipeSafe(cmd1, cmd2 *Command) *Result {
	// 执行第一个命令
	result1 := cmd1.OutputSafe()
	if !result1.Success() {
		return result1
	}
	
	// 将第一个命令的输出作为第二个命令的输入
	cmd2.Stdin(bytes.NewReader(result1.Stdout))
	return cmd2.OutputSafe()
}

// ============================================================
// Result 辅助方法
// ============================================================

// Success 检查命令是否执行成功
func (r *Result) Success() bool {
	return r.ExitCode == 0
}

// StdoutString 返回标准输出的字符串形式
func (r *Result) StdoutString() string {
	return string(r.Stdout)
}

// StderrString 返回标准错误的字符串形式
func (r *Result) StderrString() string {
	return string(r.Stderr)
}

// StdoutInt64 返回标准输出的 int64 形式
func (r *Result) StdoutInt64() int64 {
	return _as.Int64(r.Stdout)
}

// StdoutLines 返回标准输出按行分割的切片
func (r *Result) StdoutLines() []string {
	if len(r.Stdout) == 0 {
		return []string{}
	}
	lines := strings.Split(strings.TrimRight(string(r.Stdout), "\n"), "\n")
	return lines
}

// StderrLines 返回标准错误按行分割的切片
func (r *Result) StderrLines() []string {
	if len(r.Stderr) == 0 {
		return []string{}
	}
	lines := strings.Split(strings.TrimRight(string(r.Stderr), "\n"), "\n")
	return lines
}

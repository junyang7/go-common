package _cmd

import (
	"bytes"
	"context"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_assert"
	"os"
	"strings"
	"testing"
	"time"
)

var name string = "echo"
var args []string = []string{"-n", "1"}

// ============================================================
// 简单方式测试（默认 panic）
// ============================================================

func TestExecute(t *testing.T) {
	{
		var expect []byte = []byte("1")
		get := Execute(name, args...)
		_assert.Equal(t, _as.String(expect), _as.String(get))
	}
}

func TestExecuteAsInt64(t *testing.T) {
	{
		var expect int64 = 1
		get := ExecuteAsInt64(name, args...)
		_assert.Equal(t, expect, get)
	}
}

func TestExecuteAsString(t *testing.T) {
	{
		var expect string = "1"
		get := ExecuteAsString(name, args...)
		_assert.Equal(t, expect, get)
	}
}

func TestExecuteSafe(t *testing.T) {
	// 成功的命令
	{
		result := ExecuteSafe("echo", "-n", "hello")
		_assert.True(t, result.Success())
		_assert.Equal(t, 0, result.ExitCode)
		_assert.Equal(t, "hello", result.StdoutString())
		_assert.Nil(t, result.Error)
	}

	// 失败的命令（命令不存在）
	{
		result := ExecuteSafe("nonexistentcommand12345")
		_assert.False(t, result.Success())
		_assert.NotNil(t, result.Error)
	}

	// 退出码非0的命令
	{
		result := ExecuteSafe("sh", "-c", "exit 42")
		_assert.False(t, result.Success())
		_assert.Equal(t, 42, result.ExitCode)
	}

	// 有 stderr 输出的命令
	{
		result := ExecuteSafe("sh", "-c", "echo 'error message' >&2")
		_assert.True(t, result.Success())
		_assert.Contains(t, result.StderrString(), "error message")
	}
}

func TestExecuteInteractive(t *testing.T) {
	// 这个方法会输出到 stdout，很难测试
	// 这里只测试不会 panic
	ExecuteInteractive("echo 'test interactive'")
	_assert.True(t, true) // 如果执行到这里说明没有 panic
}

// ============================================================
// 构建器模式测试（默认 panic）
// ============================================================

func TestCommand_Basic(t *testing.T) {
	// 基础使用（成功）
	{
		output := New("echo", "-n", "hello").Output()
		_assert.Equal(t, "hello", string(output))
	}

	// 类型转换
	{
		str := New("echo", "-n", "hello").OutputAsString()
		_assert.Equal(t, "hello", str)
	}

	{
		num := New("echo", "-n", "123").OutputAsInt64()
		_assert.Equal(t, int64(123), num)
	}
}

func TestCommand_Safe(t *testing.T) {
	// 成功的命令
	{
		result := New("echo", "-n", "hello").OutputSafe()
		_assert.True(t, result.Success())
		_assert.Equal(t, "hello", result.StdoutString())
	}

	// 失败的命令（不会 panic）
	{
		result := New("sh", "-c", "exit 1").OutputSafe()
		_assert.False(t, result.Success())
		_assert.Equal(t, 1, result.ExitCode)
	}

	// RunSafe
	{
		result := New("echo", "test").RunSafe()
		_assert.True(t, result.Success())
	}
}

func TestCommand_Dir(t *testing.T) {
	// 设置工作目录
	{
		output := New("pwd").Dir("/tmp").Output()
		_assert.Contains(t, string(output), "/tmp")
	}
}

func TestCommand_Env(t *testing.T) {
	// 设置环境变量（完全替换）
	{
		output := New("sh", "-c", "printf '%s' \"$MY_VAR\"").
			Env([]string{"MY_VAR=hello"}).
			Output()
		_assert.Equal(t, "hello", string(output))
	}

	// 添加环境变量
	{
		output := New("sh", "-c", "printf '%s' \"$NEW_VAR\"").
			AddEnv("NEW_VAR", "world").
			Output()
		_assert.Equal(t, "world", string(output))
	}
}

func TestCommand_Stdin(t *testing.T) {
	// 使用 Reader
	{
		input := bytes.NewReader([]byte("hello from stdin"))
		output := New("cat").Stdin(input).Output()
		_assert.Equal(t, "hello from stdin", string(output))
	}

	// 使用字符串
	{
		output := New("cat").StdinString("hello world").Output()
		_assert.Equal(t, "hello world", string(output))
	}
}

func TestCommand_Stdout(t *testing.T) {
	// 重定向到 buffer
	{
		var buf bytes.Buffer
		New("echo", "test").Stdout(&buf).Run()
		_assert.Contains(t, buf.String(), "test")
	}
}

func TestCommand_Stderr(t *testing.T) {
	// 捕获 stderr
	{
		var stderr bytes.Buffer
		New("sh", "-c", "echo 'error' >&2").
			Stderr(&stderr).
			Run()
		_assert.Contains(t, stderr.String(), "error")
	}
}

func TestCommand_Timeout(t *testing.T) {
	// 正常完成
	{
		output := New("echo", "hello").
			Timeout(time.Second * 5).
			Output()
		_assert.Equal(t, "hello\n", string(output))
	}

	// 超时（Safe 模式测试）
	{
		result := New("sleep", "10").
			Timeout(time.Millisecond * 100).
			OutputSafe()
		_assert.False(t, result.Success())
		_assert.NotNil(t, result.Error)
	}
}

func TestCommand_Context(t *testing.T) {
	// 正常 context
	{
		ctx := context.Background()
		output := New("echo", "hello").Context(ctx).Output()
		_assert.Contains(t, string(output), "hello")
	}

	// 取消的 context（Safe 模式测试）
	{
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // 立即取消

		result := New("sleep", "1").Context(ctx).OutputSafe()
		_assert.False(t, result.Success())
	}
}

func TestCommand_ChainedCalls(t *testing.T) {
	// 链式调用多个配置
	{
		output := New("sh", "-c", "echo $VAR").
			Dir("/tmp").
			AddEnv("VAR", "test123").
			Timeout(time.Second * 5).
			Output()

		_assert.Contains(t, string(output), "test123")
	}
}

// ============================================================
// 管道测试
// ============================================================

func TestPipe(t *testing.T) {
	// 简单管道：echo "hello\nworld" | grep "world"
	{
		cmd1 := New("echo", "hello\nworld")
		cmd2 := New("grep", "world")
		output := Pipe(cmd1, cmd2)

		_assert.Contains(t, string(output), "world")
	}
}

func TestPipeSafe(t *testing.T) {
	// 成功的管道
	{
		cmd1 := New("echo", "hello\nworld")
		cmd2 := New("grep", "world")
		result := PipeSafe(cmd1, cmd2)

		_assert.True(t, result.Success())
		_assert.Contains(t, result.StdoutString(), "world")
	}

	// 第一个命令失败
	{
		cmd1 := New("sh", "-c", "exit 1")
		cmd2 := New("cat")
		result := PipeSafe(cmd1, cmd2)

		_assert.False(t, result.Success())
	}
}

// ============================================================
// Result 辅助方法测试
// ============================================================

func TestResult_Methods(t *testing.T) {
	// 基础方法
	{
		result := ExecuteSafe("echo", "-n", "123")
		_assert.True(t, result.Success())
		_assert.Equal(t, "123", result.StdoutString())
		_assert.Equal(t, int64(123), result.StdoutInt64())
		_assert.Equal(t, "", result.StderrString())
	}
}

func TestResult_Lines(t *testing.T) {
	// 单行
	{
		result := ExecuteSafe("echo", "-n", "single line")
		lines := result.StdoutLines()
		_assert.Len(t, lines, 1)
		_assert.Equal(t, "single line", lines[0])
	}

	// 多行
	{
		result := ExecuteSafe("printf", "line1\nline2\nline3")
		lines := result.StdoutLines()
		_assert.Len(t, lines, 3)
		_assert.Equal(t, "line1", lines[0])
		_assert.Equal(t, "line2", lines[1])
		_assert.Equal(t, "line3", lines[2])
	}

	// 空输出
	{
		result := ExecuteSafe("echo", "-n", "")
		lines := result.StdoutLines()
		_assert.Len(t, lines, 0)
	}

	// stderr 行
	{
		result := ExecuteSafe("sh", "-c", "printf 'err1\nerr2' >&2")
		lines := result.StderrLines()
		_assert.Len(t, lines, 2)
		_assert.Equal(t, "err1", lines[0])
		_assert.Equal(t, "err2", lines[1])
	}
}

// ============================================================
// 快捷函数测试
// ============================================================

func TestExecuteWithTimeout(t *testing.T) {
	// 正常完成的命令
	{
		result := ExecuteWithTimeout(time.Second*5, "echo", "-n", "hello")
		_assert.Equal(t, "hello", string(result))
	}
}

func TestExecuteWithContext(t *testing.T) {
	// 正常完成的命令
	{
		ctx := context.Background()
		result := ExecuteWithContext(ctx, "echo", "-n", "hello")
		_assert.Equal(t, "hello", string(result))
	}
}

func TestExecuteWithDir(t *testing.T) {
	// 在临时目录执行命令
	{
		result := ExecuteWithDir("/tmp", "pwd")
		_assert.Contains(t, string(result), "/tmp")
	}
}

func TestExecuteWithEnv(t *testing.T) {
	// 使用自定义环境变量
	{
		env := []string{"MY_VAR=hello"}
		result := ExecuteWithEnv(env, "sh", "-c", "printf '%s' \"$MY_VAR\"")
		_assert.Equal(t, "hello", string(result))
	}
}

// ============================================================
// 实际场景测试
// ============================================================

func TestRealWorldScenarios(t *testing.T) {
	// 场景1：Git 操作
	{
		// 检查是否在 git 仓库中
		result := New("git", "rev-parse", "--git-dir").OutputSafe()
		if result.Success() {
			// 获取当前分支
			branch := New("git", "branch", "--show-current").OutputSafe()
			_assert.True(t, branch.Success())
		}
	}

	// 场景2：文件处理
	{
		// 创建临时文件
		tmpfile := "/tmp/test_cmd_" + _as.String(time.Now().Unix()) + ".txt"
		New("sh", "-c", "echo 'test content' > "+tmpfile).Run()

		// 读取文件
		output := New("cat", tmpfile).Output()
		_assert.Contains(t, string(output), "test content")

		// 删除文件
		os.Remove(tmpfile)
	}

	// 场景3：使用 stdin 传递数据
	{
		output := New("wc", "-l").
			StdinString("line1\nline2\nline3\n").
			Output()
		_assert.Contains(t, strings.TrimSpace(string(output)), "3")
	}
}

// ============================================================
// 并发测试
// ============================================================

func TestCommand_Concurrent(t *testing.T) {
	done := make(chan bool, 10)

	// 10个并发执行
	for i := 0; i < 10; i++ {
		go func(index int) {
			output := New("echo", "-n", "test").Output()
			_assert.Equal(t, "test", string(output))
			done <- true
		}(i)
	}

	// 等待所有完成
	for i := 0; i < 10; i++ {
		<-done
	}

	_assert.True(t, true)
}

// ============================================================
// 性能测试
// ============================================================

func BenchmarkExecute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Execute("echo", "-n", "hello")
	}
}

func BenchmarkExecuteSafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecuteSafe("echo", "-n", "hello")
	}
}

func BenchmarkCommandBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("echo", "-n", "hello").Output()
	}
}

func BenchmarkCommandBuilderSafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("echo", "-n", "hello").OutputSafe()
	}
}

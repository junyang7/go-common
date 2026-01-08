package _conf

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_parameter"
	"testing"
)

// Mock 配置实现
type mockConf struct {
	data map[string]interface{}
}

func newMockConf() *mockConf {
	return &mockConf{
		data: make(map[string]interface{}),
	}
}

func (m *mockConf) Byte(byte []byte) Conf {
	return m
}

func (m *mockConf) Text(text string) Conf {
	return m
}

func (m *mockConf) File(path string) Conf {
	return m
}

func (m *mockConf) Get(path string) *_parameter.Parameter {
	value := m.data[path]
	return _parameter.New(path, value)
}

func (m *mockConf) Set(path string, value interface{}) {
	m.data[path] = value
}

// 测试前重置
func setup() {
	Reset()
}

func TestLoad(t *testing.T) {
	setup()

	// 加载前应该为空
	_assert.False(t, IsLoaded())

	// 加载配置
	conf := newMockConf()
	conf.Set("host", "localhost")
	Load(conf)

	// 加载后应该不为空
	_assert.True(t, IsLoaded())
}

func TestGet(t *testing.T) {
	setup()

	// 准备配置
	conf := newMockConf()
	conf.Set("host", "localhost")
	conf.Set("port", int64(3306))
	conf.Set("debug", true)
	Load(conf)

	// 测试获取字符串
	{
		value := Get("host").String().Value()
		_assert.Equal(t, "localhost", value)
	}

	// 测试获取数字
	{
		value := Get("port").Int64().Value()
		_assert.Equal(t, int64(3306), value)
	}

	// 测试获取布尔
	{
		value := Get("debug").Bool().Value()
		_assert.Equal(t, true, value)
	}
}

func TestGetWithDefault(t *testing.T) {
	setup()

	// 准备配置（不设置某些值）
	conf := newMockConf()
	conf.Set("host", "localhost")
	Load(conf)

	// 测试默认值
	{
		value := Get("timeout").Default(30).Int64().Value()
		_assert.Equal(t, int64(30), value)
	}
}

func TestReset(t *testing.T) {
	setup()

	// 加载配置
	conf := newMockConf()
	Load(conf)
	_assert.True(t, IsLoaded())

	// 重置
	Reset()
	_assert.False(t, IsLoaded())
}

func TestIsLoaded(t *testing.T) {
	setup()

	// 未加载时
	_assert.False(t, IsLoaded())

	// 加载后
	conf := newMockConf()
	Load(conf)
	_assert.True(t, IsLoaded())

	// 重置后
	Reset()
	_assert.False(t, IsLoaded())
}

func TestConcurrent(t *testing.T) {
	setup()

	conf := newMockConf()
	conf.Set("value", "test")
	Load(conf)

	// 并发访问测试
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func() {
			value := Get("value").String().Value()
			_assert.Equal(t, "test", value)
			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}

	_assert.True(t, true)
}

func TestChainCalls(t *testing.T) {
	setup()

	// 测试链式调用
	conf := newMockConf()
	result := conf.Text("test").File("test.json").Byte([]byte("test"))

	_assert.NotNil(t, result)
}


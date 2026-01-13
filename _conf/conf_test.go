package _conf

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_parameter"
	"testing"
)

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
func setup() {
	Reset()
}
func TestLoad(t *testing.T) {
	setup()
	_assert.False(t, IsLoaded())
	conf := newMockConf()
	conf.Set("host", "localhost")
	Load(conf)
	_assert.True(t, IsLoaded())
}
func TestGet(t *testing.T) {
	setup()
	conf := newMockConf()
	conf.Set("host", "localhost")
	conf.Set("port", int64(3306))
	conf.Set("debug", true)
	Load(conf)
	{
		value := Get("host").String().Value()
		_assert.Equal(t, "localhost", value)
	}
	{
		value := Get("port").Int64().Value()
		_assert.Equal(t, int64(3306), value)
	}
	{
		value := Get("debug").Bool().Value()
		_assert.Equal(t, true, value)
	}
}
func TestGetWithDefault(t *testing.T) {
	setup()
	conf := newMockConf()
	conf.Set("host", "localhost")
	Load(conf)
	{
		value := Get("timeout").Default(30).Int64().Value()
		_assert.Equal(t, int64(30), value)
	}
}
func TestReset(t *testing.T) {
	setup()
	conf := newMockConf()
	Load(conf)
	_assert.True(t, IsLoaded())
	Reset()
	_assert.False(t, IsLoaded())
}
func TestIsLoaded(t *testing.T) {
	setup()
	_assert.False(t, IsLoaded())
	conf := newMockConf()
	Load(conf)
	_assert.True(t, IsLoaded())
	Reset()
	_assert.False(t, IsLoaded())
}
func TestConcurrent(t *testing.T) {
	setup()
	conf := newMockConf()
	conf.Set("value", "test")
	Load(conf)
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			value := Get("value").String().Value()
			_assert.Equal(t, "test", value)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	_assert.True(t, true)
}
func TestChainCalls(t *testing.T) {
	setup()
	conf := newMockConf()
	result := conf.Text("test").File("test.json").Byte([]byte("test"))
	_assert.NotNil(t, result)
}

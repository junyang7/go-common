package _port

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestIsUsing(t *testing.T) {
	{
		port := 6379
		_isUsing := IsUsing(port)
		_assert.True(t, _isUsing)
	}
}
func TestGetList(t *testing.T) {
	{
		count := 5
		availablePorts := GetList(count)
		_assert.Equal(t, count, len(availablePorts))
		for _, port := range availablePorts {
			_assert.False(t, IsUsing(port))
		}
	}
}
func TestGet(t *testing.T) {
	{
		port := Get()
		_assert.False(t, IsUsing(port))
	}
}
func TestPortNew(t *testing.T) {
	{
		p := New()
		_assert.NotNil(t, p)
		_assert.Equal(t, 20000, p.min)
		_assert.Equal(t, 65535, p.max)
		_assert.Equal(t, 10, p.count)
		_assert.Empty(t, p.filter)
	}
}
func TestPortMin(t *testing.T) {
	{
		p := New().Min(10000)
		_assert.Equal(t, 10000, p.min)
	}
}
func TestPortMax(t *testing.T) {
	{
		p := New().Max(50000)
		_assert.Equal(t, 50000, p.max)
	}
}
func TestPortCount(t *testing.T) {
	{
		p := New().Count(5)
		_assert.Equal(t, 5, p.count)
	}
}
func TestPortFilter(t *testing.T) {
	{
		filterPorts := []int{8080, 9090}
		p := New().Filter(filterPorts)
		_assert.Equal(t, filterPorts, p.filter)
	}
}
func TestPortGetList(t *testing.T) {
	{
		filterPorts := []int{10001, 10003, 10005, 10007, 10009}
		p := New().Min(10000).Max(11000).Count(10).Filter(filterPorts)
		portList := p.GetList()
		_assert.Equal(t, 10, len(portList))
		for _, port := range portList {
			_assert.GreaterOrEqual(t, port, 10000)
			_assert.LessOrEqual(t, port, 11000)
			_assert.NotIn(t, port, filterPorts)
			_assert.False(t, IsUsing(port))
		}
	}
}

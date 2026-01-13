package _cmd

import (
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_assert"
	"testing"
)

var name string = "echo"
var args []string = []string{"-n", "1"}

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
func TestExecuteByStd(t *testing.T) {

	ExecuteByStd("echo 'TestExecuteByStd'")
	_assert.True(t, true)
}

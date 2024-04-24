package _cmd

import (
	"git.ziji.fun/junyang/go-common/_as"
	"git.ziji.fun/junyang/go-common/_assert"
	"testing"
)

var name string = "echo"
var args []string = []string{"-n", "1"}

func TestEncode(t *testing.T) {
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

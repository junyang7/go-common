package _password

import (
	"git.ziji.fun/junyang/go-common/_assert"
	"testing"
)

func TestHash(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestVerify(t *testing.T) {
	{
		var password string = "password"
		hash := Hash(password)
		var expect bool = true
		get := Verify(hash, password)
		_assert.Equal(t, expect, get)
	}
}

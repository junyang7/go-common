package _uuid

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestV4(t *testing.T) {
	{
		uuid := V4()
		_assert.Equal(t, len(uuid), 32)
		_assert.NotContains(t, uuid, "-")
		_assert.Regexp(t, uuid, "^[a-zA-Z0-9]{32}$")
	}
}

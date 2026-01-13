package _unixMicro

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	{
		currentUnixMicro := Get()
		_assert.True(t, currentUnixMicro > 0)
		convertedTime := time.Unix(0, currentUnixMicro*int64(time.Microsecond))
		_assert.True(t, convertedTime.Before(time.Now()))
		_assert.True(t, convertedTime.After(time.Now().Add(-24*time.Hour))) // 一天前的时间戳
	}
}

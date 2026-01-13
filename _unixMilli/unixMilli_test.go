package _unixMilli

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_time"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	{
		currentUnixMilli := Get()
		_assert.True(t, currentUnixMilli > 0)
		convertedTime := time.Unix(0, currentUnixMilli*int64(time.Microsecond))
		_assert.True(t, convertedTime.Before(time.Now()))
		_assert.True(t, convertedTime.After(time.Now().Add(-24*time.Hour)))
	}
}
func TestGetByTime(t *testing.T) {
	{
		specificTime := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
		expectedUnixMilli := specificTime.UnixMicro()
		getUnixMilli := GetByTime(specificTime)
		_assert.Equal(t, expectedUnixMilli, getUnixMilli)
	}
}
func TestGetByDatetime(t *testing.T) {
	{
		datetime := "2023-10-01 00:00:00"
		expectedUnixMilli := _time.GetByDatetime(datetime).UnixMicro()
		getUnixMilli := GetByDatetime(datetime)
		_assert.Equal(t, expectedUnixMilli, getUnixMilli)
	}
	{
		defer func() {
			err := recover()
			_assert.NotNil(t, err)
		}()
		invalidDatetime := "invalid-datetime"
		GetByDatetime(invalidDatetime)
	}
}

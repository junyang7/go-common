package _unix

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_time"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	{
		currentUnix := Get()
		_assert.True(t, currentUnix > 0)
		convertedTime := time.Unix(currentUnix, 0)
		_assert.True(t, convertedTime.Before(time.Now()))
		_assert.True(t, convertedTime.After(time.Now().Add(-24*time.Hour)))
	}
}
func TestGetByTime(t *testing.T) {
	{
		timestamp := time.Date(2026, time.January, 9, 15, 0, 0, 0, time.UTC)
		unixTimestamp := GetByTime(timestamp)
		expectedUnix := timestamp.Unix()
		_assert.Equal(t, expectedUnix, unixTimestamp)
	}
}
func TestGetByDatetime(t *testing.T) {
	{
		datetime := "2026-01-09 15:00:00"
		unixTimestamp := GetByDatetime(datetime)
		expectedTime := _time.GetByDatetime(datetime)
		expectedUnix := expectedTime.Unix()
		_assert.Equal(t, expectedUnix, unixTimestamp)
	}
}
func TestGetByInvalidDatetime(t *testing.T) {
	{
		defer func() {
			err := recover()
			_assert.NotNil(t, err)
		}()
		invalidDatetime := "2026-99-99 99:99:99"
		GetByDatetime(invalidDatetime)
	}
}

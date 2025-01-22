package _time

import (
	"github.com/junyang7/go-common/src/_assert"
	"testing"
	"time"
)

const unix int64 = 1712126604
const unixMilli int64 = 1712126604123
const unixMicro int64 = 1712126604123456
const date string = "2024-04-03"
const datetime string = "2024-04-03 14:43:24"
const datetimeMilli string = "2024-04-03 14:43:24.123"

func TestGet(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestGetByUnix(t *testing.T) {
	{
		var expect time.Time = time.Unix(unix, 0).In(loc)
		get := GetByUnix(unix)
		_assert.EqualByTime(t, expect, get)
	}
}
func TestGetByUnixMilli(t *testing.T) {
	{
		var expect time.Time = time.UnixMilli(unixMilli).In(loc)
		get := GetByUnixMilli(unixMilli)
		_assert.EqualByTime(t, expect, get)
	}
}
func TestGetByUnixMicro(t *testing.T) {
	{
		var expect time.Time = time.UnixMicro(unixMicro).In(loc)
		get := GetByUnixMicro(unixMicro)
		_assert.EqualByTime(t, expect, get)
	}
}
func TestGetByFormatAndString(t *testing.T) {
	{
		var expect time.Time = time.UnixMilli(unixMilli).In(loc)
		get := GetByFormatAndString(FormatDatetimeMilli, datetimeMilli)
		_assert.EqualByTime(t, expect, get)
	}
}
func TestGetByDate(t *testing.T) {
	{
		var expect time.Time = GetByFormatAndString(FormatDate, date)
		get := GetByDate(date)
		_assert.EqualByTime(t, expect, get)
	}
}
func TestGetByDatetime(t *testing.T) {
	{
		var expect time.Time = GetByFormatAndString(FormatDatetime, datetime)
		get := GetByDatetime(datetime)
		_assert.EqualByTime(t, expect, get)
	}
}
func TestGetByDatetimeMilli(t *testing.T) {
	{
		var expect time.Time = GetByFormatAndString(FormatDatetimeMilli, datetimeMilli)
		get := GetByDatetimeMilli(datetimeMilli)
		_assert.EqualByTime(t, expect, get)
	}
}

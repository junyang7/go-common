package _date

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_time"
	"testing"
)

const unix int64 = 1712126604
const unixMilli int64 = 1712126604123
const unixMicro int64 = 1712126604123456
const date string = "2024-04-03"
const datetime string = "2024-04-03 14:43:24"
const datetimeMilli string = "2024-04-03 14:43:24.123"
const datetimeMicro string = "2024-04-03 14:43:24.123456"

func TestGetByTime(t *testing.T) {
	{
		var expect string = date
		get := GetByTime(_time.GetByDate(date))
		_assert.Equal(t, expect, get)
	}
}
func TestGet(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestGetByUnix(t *testing.T) {
	{
		var expect string = date
		get := GetByTime(_time.GetByUnix(unix))
		_assert.Equal(t, expect, get)
	}
}
func TestGetByUnixMilli(t *testing.T) {
	{
		var expect string = date
		get := GetByTime(_time.GetByUnixMilli(unixMilli))
		_assert.Equal(t, expect, get)
	}
}
func TestGetByUnixMicro(t *testing.T) {
	{
		var expect string = date
		get := GetByTime(_time.GetByUnixMicro(unixMicro))
		_assert.Equal(t, expect, get)
	}
}
func TestGetListByDateSAndDateE(t *testing.T) {
	{
		start := "2024-04-01"
		end := "2024-04-03"
		expected := []string{"2024-04-01", "2024-04-02", "2024-04-03"}
		got := GetListByDateSAndDateE(start, end)
		_assert.Equal(t, expected, got)
	}
	{
		start := "2024-04-01"
		end := "2024-04-01"
		expected := []string{"2024-04-01"}
		got := GetListByDateSAndDateE(start, end)
		_assert.Equal(t, expected, got)
	}
	{
		defer func() {
			r := recover()
			_assert.NotEqual(t, nil, r)
		}()
		start := "2024-04-05"
		end := "2024-04-01"
		_ = GetListByDateSAndDateE(start, end)
	}
}

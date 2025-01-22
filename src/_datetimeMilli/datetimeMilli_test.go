package _datetimeMilli

import (
	"github.com/junyang7/go-common/src/_assert"
	"github.com/junyang7/go-common/src/_time"
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
		var expect string = datetimeMilli
		get := GetByTime(_time.GetByDatetimeMilli(datetimeMilli))
		_assert.Equal(t, expect, get)
	}
}
func TestGet(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestGetByUnix(t *testing.T) {
	{
		var expect string = datetimeMilli[0:len(datetimeMilli)-3] + "000"
		get := GetByTime(_time.GetByUnix(unix))
		_assert.Equal(t, expect, get)
	}
}
func TestGetByUnixMilli(t *testing.T) {
	{
		var expect string = datetimeMilli
		get := GetByTime(_time.GetByUnixMilli(unixMilli))
		_assert.Equal(t, expect, get)
	}
}
func TestGetByUnixMicro(t *testing.T) {
	{
		var expect string = datetimeMilli
		get := GetByTime(_time.GetByUnixMicro(unixMicro))
		_assert.Equal(t, expect, get)
	}
}

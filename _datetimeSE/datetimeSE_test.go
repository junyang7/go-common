package _datetimeSE

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestBuildByY(t *testing.T) {
	{
		datetimeS := "2020-09-01 00:59:00"
		datetimeE := "2024-01-01 12:00:00"
		list := BuildByY(datetimeS, datetimeE)
		_assert.Equal(t, 5, len(list))
		_assert.Equal(t, datetimeS, list[0].DatetimeS)
		_assert.Equal(t, "2020-12-31 23:59:59", list[0].DatetimeE)
		_assert.Equal(t, "2024-01-01 00:00:00", list[len(list)-1].DatetimeS)
		_assert.Equal(t, datetimeE, list[len(list)-1].DatetimeE)
	}
	{
		datetimeS := "2023-02-01 00:00:00"
		datetimeE := "2023-11-30 12:00:00"
		list := BuildByY(datetimeS, datetimeE)
		_assert.Equal(t, 1, len(list))
		_assert.Equal(t, datetimeS, list[0].DatetimeS)
		_assert.Equal(t, datetimeE, list[0].DatetimeE)
	}
}
func TestBuildByYm(t *testing.T) {
	{
		datetimeS := "2020-09-01 00:59:00"
		datetimeE := "2021-02-15 12:00:00"
		list := BuildByYm(datetimeS, datetimeE)
		_assert.Equal(t, 6, len(list))
		_assert.Equal(t, datetimeS, list[0].DatetimeS)
		_assert.Equal(t, "2020-09-30 23:59:59", list[0].DatetimeE)
		_assert.Equal(t, "2021-02-01 00:00:00", list[len(list)-1].DatetimeS)
		_assert.Equal(t, datetimeE, list[len(list)-1].DatetimeE)
	}
	{
		datetimeS := "2023-03-05 08:00:00"
		datetimeE := "2023-03-20 18:00:00"
		list := BuildByYm(datetimeS, datetimeE)
		_assert.Equal(t, 1, len(list))
		_assert.Equal(t, datetimeS, list[0].DatetimeS)
		_assert.Equal(t, "2023-03-31 23:59:59", list[0].DatetimeE)
	}
	{
		datetimeS := "2024-11-15 00:00:00"
		datetimeE := "2025-02-05 12:00:00"
		list := BuildByYm(datetimeS, datetimeE)
		_assert.Equal(t, 4, len(list))
		_assert.Equal(t, datetimeS, list[0].DatetimeS)
		_assert.Equal(t, "2024-11-30 23:59:59", list[0].DatetimeE)
		_assert.Equal(t, "2025-02-01 00:00:00", list[3].DatetimeS)
		_assert.Equal(t, datetimeE, list[3].DatetimeE)
	}
	{
		datetimeS := "2025-05-10 09:00:00"
		datetimeE := "2025-05-10 09:00:00"
		list := BuildByYm(datetimeS, datetimeE)
		_assert.Equal(t, 1, len(list))
		_assert.Equal(t, datetimeS, list[0].DatetimeS)
		_assert.Equal(t, "2025-05-31 23:59:59", list[0].DatetimeE)
	}
}

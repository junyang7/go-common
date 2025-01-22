package _unixMilli

import (
	"github.com/junyang7/go-common/src/_time"
	"time"
)

func Get() int64 {
	return GetByTime(time.Now())
}
func GetByTime(t time.Time) int64 {
	return t.UnixMicro()
}
func GetByDatetime(datetime string) int64 {
	return _time.GetByDatetime(datetime).UnixMicro()
}

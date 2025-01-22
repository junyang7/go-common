package _datetimeMilli

import (
	"github.com/junyang7/go-common/src/_time"
	"time"
)

func GetByTime(t time.Time) string {
	return t.Format(_time.FormatDatetimeMilli)
}
func Get() string {
	return GetByTime(time.Now())
}
func GetByUnix(unix int64) string {
	return GetByTime(_time.GetByUnix(unix))
}
func GetByUnixMilli(unixMilli int64) string {
	return GetByTime(_time.GetByUnixMilli(unixMilli))
}
func GetByUnixMicro(unixMicro int64) string {
	return GetByTime(_time.GetByUnixMicro(unixMicro))
}

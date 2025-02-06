package _unix

import (
	"github.com/junyang7/go-common/_time"
	"time"
)

func Get() int64 {
	return time.Now().Unix()
}
func GetByTime(t time.Time) int64 {
	return t.Unix()
}
func GetByDatetime(datetime string) int64 {
	return _time.GetByDatetime(datetime).Unix()
}

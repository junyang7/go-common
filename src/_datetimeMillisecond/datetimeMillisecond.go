package _datetimeMillisecond

import (
	"fmt"
	"github.com/junyang7/go-common/src/_time"
	"time"
)

func GetByTime(t time.Time) string {
	return t.Format(fmt.Sprintf("%s-%s-%s %s:%s:%s.%s", _time.YEAR, _time.MONTH, _time.DAY, _time.HOUR, _time.MINUTE, _time.SECOND, _time.MILLISECOND))
}
func Get() string {
	return GetByTime(time.Now())
}

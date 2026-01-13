package _date

import (
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_time"
	"time"
)

func GetByTime(t time.Time) string {
	return t.Format(_time.FormatDate)
}
func Get() string {
	return GetByTime(time.Now())
}
func GetByYmd() string {
	return _time.Get().Format(_time.FormatYmd)
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
func GetListByDateSAndDateE(dateS string, dateE string) []string {
	timeS := _time.GetByDate(dateS)
	timeE := _time.GetByDate(dateE)
	if timeS.After(timeE) {
		_interceptor.Insure(false).Message("开始日期<结束日期").Do()
	}
	dateList := []string{}
	for t := timeS; !t.After(timeE); t = t.AddDate(0, 0, 1) {
		dateList = append(dateList, GetByTime(t))
	}
	return dateList
}

package _time

import (
	"time"
)

const (
	Year                = "2006"
	Month               = "01"
	Day                 = "02"
	Hour                = "15"
	Minute              = "04"
	Second              = "05"
	Millisecond         = "000"
	FormatDate          = "2006-01-02"
	FormatDatetime      = "2006-01-02 15:04:05"
	FormatDatetimeMilli = "2006-01-02 15:04:05.000"
)

var loc *time.Location
var err error

func Get() time.Time {
	return time.Now().In(loc)
}
func GetByUnix(unix int64) time.Time {
	return time.Unix(unix, 0).In(loc)
}
func GetByUnixMilli(unixMilli int64) time.Time {
	return time.UnixMilli(unixMilli).In(loc)
}
func GetByUnixMicro(unixMicro int64) time.Time {
	return time.UnixMicro(unixMicro).In(loc)
}
func GetByFormatAndString(f string, s string) time.Time {
	t, err := time.ParseInLocation(f, s, loc)
	if nil != err {
		panic(err)
	}
	return t
}
func GetByDate(date string) time.Time {
	return GetByFormatAndString(FormatDate, date)
}
func GetByDatetime(datetime string) time.Time {
	return GetByFormatAndString(FormatDatetime, datetime)
}
func GetByDatetimeMilli(datetimeMilli string) time.Time {
	return GetByFormatAndString(FormatDatetimeMilli, datetimeMilli)
}
func init() {
	loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}

package _time

import (
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_string"
	"strings"
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
	FormatYmdHis        = "20060102150405"
	FormatYmd           = "20060102"
	FormatYm            = "200601"
)

var loc *time.Location
var err error

type Formatted struct {
	Year   string `json:"year"`   // 0001~9999
	Month  string `json:"month"`  // 01~12
	Week   string `json:"week"`   // 01~07
	Day    string `json:"day"`    // 01~31
	Hour   string `json:"hour"`   // 00~23
	Minute string `json:"minute"` // 00~59
	Second string `json:"second"` // 00~59
}

func WaitUntilMilli(target int64) int64 {
	now := time.Now().UnixMilli()
	for now < target {
		time.Sleep(time.Millisecond)
		now = time.Now().UnixMilli()
	}
	return now
}
func WaitNextMilli(target int64) int64 {
	now := time.Now().UnixMilli()
	for now <= target {
		time.Sleep(time.Millisecond)
		now = time.Now().UnixMilli()
	}
	return now
}
func FormatByTime(t time.Time) (formatted *Formatted) {
	formatted = &Formatted{}
	formatted.Year = _as.String(t.Year())
	formatted.Month = _string.PadLeft(_as.String(int(t.Month())), 2, "0")
	week := int(t.Weekday())
	if week == 0 {
		week = 7
	}
	formatted.Week = _string.PadLeft(_as.String(week), 2, "0")
	formatted.Day = _string.PadLeft(_as.String(t.Day()), 2, "0")
	formatted.Hour = _string.PadLeft(_as.String(t.Hour()), 2, "0")
	formatted.Minute = _string.PadLeft(_as.String(t.Minute()), 2, "0")
	formatted.Second = _string.PadLeft(_as.String(t.Second()), 2, "0")
	return formatted
}
func FormatByCron(cron string) (formatted *Formatted) {
	tList := []string{}
	for _, v := range strings.Split(strings.TrimSpace(cron), " ") {
		v = strings.TrimSpace(v)
		if "" != v {
			tList = append(tList, v)
		}
	}
	_interceptor.
		Insure(len(tList) == 7).
		Message("参数错误：cron").
		Data(map[string]interface{}{
			"cron": cron,
		}).
		Do()
	formatted = &Formatted{
		Year:   tList[0],
		Month:  tList[1],
		Week:   tList[2],
		Day:    tList[3],
		Hour:   tList[4],
		Minute: tList[5],
		Second: tList[6],
	}
	return formatted
}
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
		_interceptor.Insure(false).Message(err).Do()
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
func Format(t time.Time, format string) string {
	return t.Format(format)
}
func init() {
	loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
}

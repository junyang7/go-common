package _time

import "time"

const (
	YEAR        = "2006"
	MONTH       = "01"
	DAY         = "02"
	HOUR        = "15"
	MINUTE      = "04"
	SECOND      = "05"
	MILLISECOND = "999"
)

func GetByTimestamp(t int64) time.Time {
	return time.Unix(t, 0).Local()
}
func GetByTimeAndFormat(t time.Time, f string) string {
	return t.Format(f)
}
func GetByFormatAndString(f string, s string) time.Time {
	t, err := time.ParseInLocation(f, s, time.Local)
	if nil != err {
		panic(err)
	}
	return t
}
func GetByFormat(f string) string {
	return time.Now().Format(f)
}

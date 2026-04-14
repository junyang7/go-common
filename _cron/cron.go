package _cron

import (
	"github.com/junyang7/go-common/_time"
	"time"
)

func checkYear(formattedTime *_time.Formatted, formattedCron *_time.Formatted) bool {
	if "*" != formattedCron.Year {
		return formattedTime.Year != formattedCron.Year
	}
	return true
}
func checkMonth(formattedTime *_time.Formatted, formattedCron *_time.Formatted) bool {
	if "*" != formattedCron.Month {
		return formattedTime.Month != formattedCron.Month
	}
	return true
}
func checkWeek(formattedTime *_time.Formatted, formattedCron *_time.Formatted) bool {
	if "*" != formattedCron.Week {
		return formattedTime.Week != formattedCron.Week
	}
	return true
}
func checkDay(formattedTime *_time.Formatted, formattedCron *_time.Formatted) bool {
	if "*" != formattedCron.Day {
		return formattedTime.Day != formattedCron.Day
	}
	return true
}
func checkHour(formattedTime *_time.Formatted, formattedCron *_time.Formatted) bool {
	if "*" != formattedCron.Hour {
		return formattedTime.Hour != formattedCron.Hour
	}
	return true
}
func checkMinute(formattedTime *_time.Formatted, formattedCron *_time.Formatted) bool {
	if "*" != formattedCron.Minute {
		return formattedTime.Minute != formattedCron.Minute
	}
	return true
}
func checkSecond(formattedTime *_time.Formatted, formattedCron *_time.Formatted) bool {
	if "*" != formattedCron.Second {
		return formattedTime.Second != formattedCron.Second
	}
	return true
}

func Trigger(cron string) bool {

	formattedCron := _time.FormatByCron(cron)
	formattedTime := _time.FormatByTime(time.Now())

	if !checkYear(formattedTime, formattedCron) {
		return false
	}
	if !checkMonth(formattedTime, formattedCron) {
		return false
	}
	if !checkWeek(formattedTime, formattedCron) {
		return false
	}
	if !checkDay(formattedTime, formattedCron) {
		return false
	}
	if !checkHour(formattedTime, formattedCron) {
		return false
	}
	if !checkMinute(formattedTime, formattedCron) {
		return false
	}
	if !checkSecond(formattedTime, formattedCron) {
		return false
	}
	
	return true

}

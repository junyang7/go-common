package _datetimeSE

import (
	"fmt"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_time"
)

type DatetimeSE struct {
	DatetimeS string
	DatetimeE string
}

func BuildByY(datetimeS string, datetimeE string) (o []DatetimeSE) {
	yS := _as.Int(_time.GetByDatetime(datetimeS).Format(_time.Year))
	yE := _as.Int(_time.GetByDatetime(datetimeE).Format(_time.Year))
	if yE == yS {
		datetimeSE := DatetimeSE{DatetimeS: datetimeS, DatetimeE: datetimeE}
		o = append(o, datetimeSE)
	} else {
		for i := yS; i <= yE; i++ {
			datetimeSE := DatetimeSE{}
			if i == yS {
				datetimeSE.DatetimeS = datetimeS
				datetimeSE.DatetimeE = fmt.Sprintf("%v-12-31 23:59:59", i)
			} else if i == yE {
				datetimeSE.DatetimeS = fmt.Sprintf("%v-01-01 00:00:00", i)
				datetimeSE.DatetimeE = datetimeE
			} else {
				datetimeSE.DatetimeS = fmt.Sprintf("%v-01-01 00:00:00", i)
				datetimeSE.DatetimeE = fmt.Sprintf("%v-12-31 23:59:59", i)
			}
			o = append(o, datetimeSE)
		}
	}
	return o
}
func BuildByYm(datetimeS string, datetimeE string) (o []DatetimeSE) {
	ymMin := _as.Int(_time.GetByDatetime(datetimeS).Format(_time.FormatYm) + "01")
	ymMax := _as.Int(_time.GetByDatetime(datetimeE).Format(_time.FormatYm) + "01")
	ymTep := ymMin
	for {
		if ymTep > ymMax {
			break
		}
		datetimeSE := DatetimeSE{}
		if ymTep == ymMin {
			datetimeSE.DatetimeS = datetimeS
			datetimeSE.DatetimeE = _time.GetByFormatAndString(_time.FormatYmd, _as.String(ymTep)).AddDate(0, 1, 0).AddDate(0, 0, -1).Format(_time.FormatDate) + " 23:59:59"
		} else if ymTep == ymMax {
			datetimeSE.DatetimeS = _time.GetByFormatAndString(_time.FormatYmd, _as.String(ymTep)).Format(_time.FormatDate + " 00:00:00")
			datetimeSE.DatetimeE = datetimeE
		} else {
			datetimeSE.DatetimeS = _time.GetByFormatAndString(_time.FormatYmd, _as.String(ymTep)).Format(_time.FormatDate + " 00:00:00")
			datetimeSE.DatetimeE = _time.GetByFormatAndString(_time.FormatYmd, _as.String(ymTep)).AddDate(0, 1, 0).AddDate(0, 0, -1).Format(_time.FormatDate) + " 23:59:59"
		}
		o = append(o, datetimeSE)
		ymTep = _as.Int(_time.GetByFormatAndString(_time.FormatYmd, _as.String(ymTep)).AddDate(0, 1, 0).Format(_time.FormatYm) + "01")
	}
	return o
}

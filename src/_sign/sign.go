package _sign

import (
	"github.com/junyang7/go-common/src/_hash"
	"github.com/junyang7/go-common/src/_slice"
	"sort"
)

func Md5(dataList map[string]string, key string, filterLabelList []string, filterValueList []string) string {
	var dataLabelList []string
	filterDataList := map[string]string{}
	for dataLabel, dataValue := range dataList {
		if len(filterLabelList) > 0 {
			if _slice.In(dataLabel, filterLabelList) {
				continue
			}
		}
		if len(filterValueList) > 0 {
			for _, filterValue := range filterValueList {
				if filterValue == dataValue {
					continue
				}
			}
		}
		dataLabelList = append(dataLabelList, dataLabel)
		filterDataList[dataLabel] = dataValue
	}
	sort.Strings(dataLabelList)
	s := ""
	for _, dataLabel := range dataLabelList {
		s += "&" + dataLabel + "=" + filterDataList[dataLabel]
	}
	s = s[1:]
	s += key
	sign := _hash.Md5(s)
	return sign
}

package _sign

import (
	"github.com/junyang7/go-common/_hash"
	"sort"
	"strings"
)

func Md5(dataList map[string]string, key string, filterLabelList []string, filterValueList []string) string {
	var dataLabelList []string
	filterDataList := map[string]string{}
	for dataLabel, dataValue := range dataList {
		dataLabel = strings.TrimSpace(dataLabel)
		dataValue = strings.TrimSpace(dataValue)
		if len(filterLabelList) > 0 {
			for _, filterLabel := range filterLabelList {
				if filterLabel == dataLabel {
					continue
				}
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
func HamcSha1(dataList map[string]string, key string, filterLabelList []string, filterValueList []string) string {
	var dataLabelList []string
	filterDataList := map[string]string{}
	for dataLabel, dataValue := range dataList {
		dataLabel = strings.TrimSpace(dataLabel)
		dataValue = strings.TrimSpace(dataValue)
		if len(filterLabelList) > 0 {
			for _, filterLabel := range filterLabelList {
				if filterLabel == dataLabel {
					continue
				}
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
	sign := _hash.HmacSha1(s, key)
	return sign
}

package _slice

import "github.com/junyang7/go-common/src/_as"

type Slice interface {
	string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | bool
}

func Unique[K Slice](elementList []K) []K {
	var res []K
	filter := map[K]bool{}
	for _, element := range elementList {
		if _, ok := filter[element]; !ok {
			res = append(res, element)
			filter[element] = true
		}
	}
	return res
}
func In[K Slice](element K, elementList []K) bool {
	for _, v := range elementList {
		if v == element {
			return true
		}
	}
	return false
}
func Implode[K Slice](elementList []K, separator string) string {
	res := ""
	for i, j := 0, len(elementList); i < j; i++ {
		res += _as.String(elementList[i])
		if i < j-1 {
			res += separator
		}
	}
	return res
}
func Group[K Slice, V Slice](mapList []map[K]V, key K) map[V]map[K]V {
	res := map[V]map[K]V{}
	for _, v := range mapList {
		res[v[key]] = v
	}
	return res
}
func Column[K Slice, V Slice](mapList []map[K]V, key K) []V {
	var res []V
	for _, v := range mapList {
		res = append(res, v[key])
	}
	return res
}

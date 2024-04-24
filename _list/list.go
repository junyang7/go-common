package _list

import "git.ziji.fun/junyang/go-common/_as"

type List interface {
	string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | bool
}

func Unique[K List](elementList []K) []K {
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
func In[K List](element K, elementList []K) bool {
	for _, v := range elementList {
		if v == element {
			return true
		}
	}
	return false
}
func Implode[K List](separator string, elementList []K) string {
	res := ""
	for i, j := 0, len(elementList); i < j; i++ {
		res += _as.String(elementList[i])
		if i < j-1 {
			res += separator
		}
	}
	return res
}
func Column[K List, V List](mapList []map[K]V, key K) []V {
	var res []V
	for _, v := range mapList {
		if _, ok := v[key]; ok {
			res = append(res, v[key])
		}
	}
	return res
}
func Group[K List, V List](mapList []map[K]V, key K) map[V][]map[K]V {
	res := map[V][]map[K]V{}
	for _, _map := range mapList {
		if _, ok := _map[key]; ok {
			if _, ok := res[_map[key]]; !ok {
				res[_map[key]] = []map[K]V{}
			}
			res[_map[key]] = append(res[_map[key]], _map)
		}
	}
	return res
}

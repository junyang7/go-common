package _list

import (
	"fmt"
	"reflect"
)

type List interface {
	string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | bool
}

func Unique[V comparable](vList []V) []V {
	var res []V
	filter := map[V]bool{}
	for _, v := range vList {
		if _, ok := filter[v]; !ok {
			res = append(res, v)
			filter[v] = true
		}
	}
	return res
}
func UniqueDeep[V any](vList []V) []V {
	var oList []V
	for _, n := range vList {
		found := false
		for _, o := range oList {
			if reflect.DeepEqual(n, o) {
				found = true
				break
			}
		}
		if !found {
			oList = append(oList, n)
		}
	}
	return oList
}

func In[V comparable](v V, vList []V) bool {
	for _, one := range vList {
		if v == one {
			return true
		}
	}
	return false
}
func InDeep[V any](v V, vList []V) bool {
	for _, one := range vList {
		if reflect.DeepEqual(v, one) {
			return true
		}
	}
	return false
}
func Implode[K any](separator string, elementList []K) string {
	res := ""
	for i, j := 0, len(elementList); i < j; i++ {
		res += fmt.Sprintf("%v", elementList[i])
		if i < j-1 {
			res += separator
		}
	}
	return res
}
func Column[V any, K any](vList []V, f func(V) K) []K {
	var res []K
	for _, v := range vList {
		res = append(res, f(v))
	}
	return res
}
func GroupListBy[V any, K comparable](vList []V, f func(V) K) map[K][]V {
	res := make(map[K][]V)
	for _, v := range vList {
		k := f(v)
		res[k] = append(res[k], v)
	}
	return res
}
func GroupBy[V any, K comparable](vList []V, f func(V) K) map[K]V {
	res := make(map[K]V)
	for _, v := range vList {
		k := f(v)
		res[k] = v
	}
	return res
}

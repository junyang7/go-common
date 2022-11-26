package _map

type _k interface {
	string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | bool
}
type _v interface {
	interface{}
}

func KeyList[k _k, v _v](elementList map[k]v) []k {
	var res []k
	for k, _ := range elementList {
		res = append(res, k)
	}
	return res
}
func ValueList[k _k, v _v](elementList map[k]v) []v {
	var res []v
	for _, v := range elementList {
		res = append(res, v)
	}
	return res
}

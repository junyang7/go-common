package _mapList

type _k interface {
	string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | bool
}
type _v interface {
	string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | bool
}

func ArrayColumn[k _k, v _v](elementList []map[k]v, column k) []v {
	var res []v
	for _, element := range elementList {
		if value, ok := element[column]; ok {
			res = append(res, value)
		}
	}
	return res
}

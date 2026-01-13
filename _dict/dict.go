package _dict

type K interface {
	string | int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64 | bool
}
type V interface{}

func KeyList[k K, v V](elementList map[k]v) []k {
	res := make([]k, 0, len(elementList))
	for key, _ := range elementList {
		res = append(res, key)
	}
	return res
}
func ValueList[k K, v V](elementList map[k]v) []v {
	res := make([]v, 0, len(elementList))
	for _, v := range elementList {
		res = append(res, v)
	}
	return res
}

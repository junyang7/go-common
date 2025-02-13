package _string

import (
	"strings"
	"unicode"
)

const L int = 1
const R int = 2

func Pad(s string, length int, sp string, t int) string {
	res := s
	diff := length - len(s)
	if diff > 0 {
		if t == L {
			res = strings.Repeat(sp, diff) + s
		} else {
			res = s + strings.Repeat(sp, diff)
		}
	}
	return res
}
func PadLeft(s string, l int, ps string) string {
	return Pad(s, l, ps, L)
}
func PadRight(s string, l int, ps string) string {
	return Pad(s, l, ps, R)
}
func ReplaceAll(s string, old string, new string) string {
	return strings.Replace(s, old, new, -1)
}
func ToUpperCamelCase(s string) string {
	partList := strings.Split(s, "_")
	for index, part := range partList {
		if len(part) > 0 && unicode.IsLower(rune(part[0])) {
			partList[index] = strings.ToUpper(part[0:1]) + part[1:]
		}
	}
	return strings.Join(partList, "")
}

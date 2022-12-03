package _is

import "strings"

func Empty(value string) bool {
	return "" == strings.TrimSpace(value)
}

func NotEmpty(value string) bool {
	return "" != strings.TrimSpace(value)
}

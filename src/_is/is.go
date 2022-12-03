package _is

import "strings"

func Empty(value string) bool {
	return "" == strings.TrimSpace(value)
}

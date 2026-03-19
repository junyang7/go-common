package _sqlHelper

import "strings"

func BuildWhereInTemplate[T any](sList []T) string {
	if len(sList) == 0 {
		return ""
	}
	return strings.Repeat(",?", len(sList))[1:]
}

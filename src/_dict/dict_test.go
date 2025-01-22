package _dict

import (
	"github.com/junyang7/go-common/src/_assert"
	"sort"
	"testing"
)

func TestKeyList(t *testing.T) {
	{
		var give map[string]string = map[string]string{
			"A": "a",
			"B": "b",
			"C": "c",
		}
		var expect []string = []string{"A", "B", "C"}
		get := KeyList(give)
		sort.Strings(get)
		_assert.EqualByList(t, expect, get)
	}
}
func TestValueList(t *testing.T) {
	{
		var give map[string]string = map[string]string{
			"A": "a",
			"B": "b",
			"C": "c",
		}
		var expect []string = []string{"a", "b", "c"}
		get := ValueList(give)
		sort.Strings(get)
		_assert.EqualByList(t, expect, get)
	}
}

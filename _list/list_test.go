package _list

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestUnique(t *testing.T) {
	{
		var give []int = []int{1, 1, 2, 2, 3, 3}
		var expect []int = []int{1, 2, 3}
		get := Unique(give)
		_assert.EqualByList(t, expect, get)
	}
}
func TestIn(t *testing.T) {
	{
		var expect bool = true
		get := In(1, []int{1, 2, 3})
		_assert.EqualByList(t, expect, get)
	}
	{
		var expect bool = false
		get := In(4, []int{1, 2, 3})
		_assert.EqualByList(t, expect, get)
	}
}
func TestImplode(t *testing.T) {
	{
		var expect string = "1,2,3"
		get := Implode(",", []int{1, 2, 3})
		_assert.EqualByList(t, expect, get)
	}
	{
		var expect string = "1,2,3"
		get := Implode(",", []string{"1", "2", "3"})
		_assert.EqualByList(t, expect, get)
	}
}

//func TestColumn(t *testing.T) {
//	{
//		dataList := []map[string]string{
//			{
//				"AZ": "A",
//				"az": "a",
//			},
//			{
//				"AZ": "B",
//				"az": "b",
//			},
//			{
//				"AZ": "C",
//				"az": "c",
//			},
//		}
//		var expect []string = []string{"A", "B", "C"}
//		get := Column(dataList, "AZ")
//		_assert.EqualByList(t, expect, get)
//	}
//}
//func TestGroup(t *testing.T) {
//	{
//		dataList := []map[string]string{
//			{
//				"AZ": "A",
//			},
//			{
//				"AZ": "B",
//			},
//			{
//				"AZ": "C",
//			},
//		}
//		var expect map[string][]map[string]string = map[string][]map[string]string{
//			"A": []map[string]string{{"AZ": "A"}},
//			"B": []map[string]string{{"AZ": "B"}},
//			"C": []map[string]string{{"AZ": "C"}},
//		}
//		get := Group(dataList, "AZ")
//		_assert.EqualByList(t, expect, get)
//	}
//}

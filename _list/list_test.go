package _list

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestUnique(t *testing.T) {
	{
		input := []int{1, 2, 2, 3, 4, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := Unique(input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []int{}
		expected := []int{}
		result := Unique(input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := Unique(input)
		_assert.Equal(t, expected, result)
	}
}
func TestUniqueDeep(t *testing.T) {
	{
		input := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
			{"Alice", 30},
		}
		expected := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
		}
		result := UniqueDeep(input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []struct {
			Name string
			Age  int
		}{}
		expected := []struct {
			Name string
			Age  int
		}{}
		result := UniqueDeep(input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
		}
		expected := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
		}
		result := UniqueDeep(input)
		_assert.Equal(t, expected, result)
	}
}
func TestIn(t *testing.T) {
	{
		input := []int{1, 2, 3, 4, 5}
		expected := true
		result := In(3, input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []int{1, 2, 3, 4, 5}
		expected := false
		result := In(6, input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []int{}
		expected := false
		result := In(1, input)
		_assert.Equal(t, expected, result)
	}
}
func TestInDeep(t *testing.T) {
	{
		input := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
		}
		expected := true
		result := InDeep(struct {
			Name string
			Age  int
		}{"Alice", 30}, input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
		}
		expected := false
		result := InDeep(struct {
			Name string
			Age  int
		}{"Charlie", 35}, input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []struct {
			Name string
			Age  int
		}{}
		expected := false
		result := InDeep(struct {
			Name string
			Age  int
		}{"Alice", 30}, input)
		_assert.Equal(t, expected, result)
	}
}
func TestImplode(t *testing.T) {
	{
		input := []string{"a", "b", "c", "d"}
		expected := "a,b,c,d"
		result := Implode(",", input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []int{1, 2, 3, 4}
		expected := "1234"
		result := Implode("", input)
		_assert.Equal(t, expected, result)
	}
	{
		input := []string{}
		expected := ""
		result := Implode(",", input)
		_assert.Equal(t, expected, result)
	}
}
func TestColumn(t *testing.T) {
	{
		input := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
		}
		expected := []int{30, 25}
		result := Column(input, func(v struct {
			Name string
			Age  int
		}) int {
			return v.Age
		})
		_assert.Equal(t, expected, result)
	}
	{
		input := []struct {
			Name string
			Age  int
		}{}
		expected := []int{}
		result := Column(input, func(v struct {
			Name string
			Age  int
		}) int {
			return v.Age
		})
		_assert.Equal(t, expected, result)
	}
}
func TestGroupListBy(t *testing.T) {
	{
		input := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 30},
		}
		expected := map[int][]struct {
			Name string
			Age  int
		}{
			30: {
				{"Alice", 30},
				{"Charlie", 30},
			},
			25: {
				{"Bob", 25},
			},
		}
		result := GroupListBy(input, func(v struct {
			Name string
			Age  int
		}) int {
			return v.Age
		})
		_assert.Equal(t, expected, result)
	}
}
func TestGroupBy(t *testing.T) {
	{
		input := []struct {
			Name string
			Age  int
		}{
			{"Alice", 30},
			{"Bob", 25},
			{"Alice", 35},
		}
		expected := map[string]struct {
			Name string
			Age  int
		}{
			"Alice": {"Alice", 35},
			"Bob":   {"Bob", 25},
		}
		result := GroupBy(input, func(v struct {
			Name string
			Age  int
		}) string {
			return v.Name
		})
		_assert.Equal(t, expected, result)
	}
}

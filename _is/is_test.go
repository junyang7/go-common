package _is

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestEmptyBool(t *testing.T) {
	{
		var give bool
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give bool = false
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give bool = true
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyInt(t *testing.T) {
	{
		var give int8
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int8 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int8 = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int8 = -1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int16
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int16 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int16 = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int16 = -1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int32
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int32 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int32 = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int32 = -1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int64
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int64 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int64 = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int64 = -1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give int = -1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyUint(t *testing.T) {
	{
		var give uint8
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint8 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint8 = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint16
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint16 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint16 = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint32
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint32 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint32 = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint64
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint64 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint64 = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give uint = 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyFloat(t *testing.T) {
	{
		var give float32
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give float32 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give float32 = 0.1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give float32 = -0.1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give float64
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give float64 = 0
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give float64 = 0.1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give float64 = -0.1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyString(t *testing.T) {
	{
		var give string
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give string = ""
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give string = " "
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give string = "Hello world!"
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyInterface(t *testing.T) {
	{
		var give any
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give interface{}
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyNil(t *testing.T) {
	{
		var give interface{} = nil
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmpty_Slice(t *testing.T) {
	{
		var give []int
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give []int = []int{}
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give []int = []int{1}
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmpty_Map(t *testing.T) {
	{
		var give map[string]string
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give map[string]string = map[string]string{}
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give map[string]string = map[string]string{"a": "A"}
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyChannel(t *testing.T) {
	{
		var give chan int
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give chan int = make(chan int)
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give chan int = make(chan int, 2)
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give chan int = make(chan int, 2)
		give <- 1
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give chan int = make(chan int, 2)
		give <- 1
		<-give
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyArray(t *testing.T) {
	{
		var give [3]int
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give [3]int = [3]int{}
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give [3]int = [3]int{1}
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyPtr(t *testing.T) {
	{
		var give *int
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give = new(int)
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEmptyStructure(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	{
		var give Person
		var expect bool = true
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give Person = Person{Name: "Alice"}
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give Person = Person{Name: "Alice", Age: 25}
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
	{
		var give Person = Person{Age: 25}
		var expect bool = false
		get := Empty(give)
		_assert.Equal(t, expect, get)
	}
}
func TestNumeric(t *testing.T) {
	{
		result := Numeric("1234567890")
		_assert.True(t, result)
	}
	{
		result := Numeric("123abc456")
		_assert.False(t, result)
	}
	{
		result := Numeric("")
		_assert.False(t, result)
	}
}
func TestAlpha(t *testing.T) {
	{
		result := Alpha("HelloWorld")
		_assert.True(t, result)
	}
	{
		result := Alpha("Hello123")
		_assert.False(t, result)
	}
	{
		result := Alpha("")
		_assert.False(t, result)
	}
}
func TestAlphaLower(t *testing.T) {
	{
		result := AlphaLower("helloworld")
		_assert.True(t, result)
	}
	{
		result := AlphaLower("HelloWorld")
		_assert.False(t, result)
	}
	{
		result := AlphaLower("hello123")
		_assert.False(t, result)
	}
	{
		result := AlphaLower("")
		_assert.False(t, result)
	}
}
func TestAlphaUpper(t *testing.T) {
	{
		result := AlphaUpper("HELLOWORLD")
		_assert.True(t, result)
	}
	{
		result := AlphaUpper("HelloWorld")
		_assert.False(t, result)
	}
	{
		result := AlphaUpper("HELLO123")
		_assert.False(t, result)
	}
	{
		result := AlphaUpper("")
		_assert.False(t, result)
	}
}

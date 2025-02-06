package _parameter

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestNew(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestParameter_Int(t *testing.T) {
	var give int = 1
	var expect string = "*_validator.Int"
	get := fmt.Sprintf("%T", New("int", give).Int())
	_assert.Equal(t, expect, get)
}
func TestParameter_String(t *testing.T) {
	var give string = "1"
	var expect string = "*_validator.String"
	get := fmt.Sprintf("%T", New("string", give).String())
	_assert.Equal(t, expect, get)
}
func TestParameter_Bool(t *testing.T) {
	var give bool = false
	var expect string = "*_validator.Bool"
	get := fmt.Sprintf("%T", New("bool", give).Bool())
	_assert.Equal(t, expect, get)
}
func TestParameter_Float64(t *testing.T) {
	var give float64 = 3.1415926
	var expect string = "*_validator.Float64"
	get := fmt.Sprintf("%T", New("float64", give).Float64())
	_assert.Equal(t, expect, get)
}

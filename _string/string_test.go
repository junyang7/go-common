package _string

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestPad(t *testing.T) {
	// no need to test
}
func TestPadLeft(t *testing.T) {
	{
		var base string = "hello"
		var padLen int = 10
		var padChar string = "a"
		var expect string = "aaaaahello"
		get := PadLeft(base, padLen, padChar)
		_assert.Equal(t, expect, get)
	}
}
func TestPadRight(t *testing.T) {
	{
		var base string = "hello"
		var padLen int = 10
		var padChar string = "a"
		var expect string = "helloaaaaa"
		get := PadRight(base, padLen, padChar)
		_assert.Equal(t, expect, get)
	}
}
func TestToUpperCamelCase(t *testing.T) {
	fmt.Println(ToUpperCamelCase("__world_example_A_0"))
}

package _aes

import (
	"github.com/junyang7/go-common/src/_assert"
	"testing"
)

const k32 string = "b841b78d016df9dea4fc49e13d11199d"
const i16 string = "d41d8cd98f00b204"

func TestEncode(t *testing.T) {
	{
		var expect string = "hl4S_PnxLPmEB5L8lj2ZAQ"
		var give string = "hello world!"
		get := Encode(give, k32, i16)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "6FmGhzAbjxh0D48n0WVP-3xrmFHssuDzm_FFprfOUxk"
		var give string = "您好，中国！"
		get := Encode(give, k32, i16)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "aWj6kY_hAmN9ftn6T5dUFg"
		var give string = "hello 中国"
		get := Encode(give, k32, i16)
		_assert.Equal(t, expect, get)
	}
}
func TestDecode(t *testing.T) {
	{
		var expect string = "hello world!"
		var give string = "hl4S_PnxLPmEB5L8lj2ZAQ"
		get := Decode(give, k32, i16)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "您好，中国！"
		var give string = "6FmGhzAbjxh0D48n0WVP-3xrmFHssuDzm_FFprfOUxk"
		get := Decode(give, k32, i16)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "hello 中国"
		var give string = "aWj6kY_hAmN9ftn6T5dUFg"
		get := Decode(give, k32, i16)
		_assert.Equal(t, expect, get)
	}
}

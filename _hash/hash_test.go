package _hash

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestMd5(t *testing.T) {
	{
		var give string = "hello world!"
		var expect string = "fc3ff98e8c6a0d3087d515c0473f8677"
		get := Md5(give)
		_assert.Equal(t, expect, get)
	}
}
func TestSha1(t *testing.T) {
	{
		var give string = "hello world!"
		var expect string = "430ce34d020724ed75a196dfc2ad67c77772d169"
		get := Sha1(give)
		_assert.Equal(t, expect, get)
	}
}
func TestSha256(t *testing.T) {
	{
		var give string = "hello world!"
		var expect string = "7509e5bda0c762d2bac7f90d758b5b2263fa01ccbc542ab5e3df163be08e6ca9"
		get := Sha256(give)
		_assert.Equal(t, expect, get)
	}
}
func TestSha512(t *testing.T) {
	{
		var give string = "hello world!"
		var expect string = "db9b1cd3262dee37756a09b9064973589847caa8e53d31a9d142ea2701b1b28abd97838bb9a27068ba305dc8d04a45a1fcf079de54d607666996b3cc54f6b67c"
		get := Sha512(give)
		_assert.Equal(t, expect, get)
	}
}
func TestHmacSha1(t *testing.T) {
	{
		var key string = "secret"
		var give string = "hello world!"
		var expect string = "a4df5f9d237ab0ca3241f042bcf6059a4ef491c4"
		get := HmacSha1(give, key)
		_assert.Equal(t, expect, get)
	}
}
func TestHmacSha256(t *testing.T) {
	{
		var key string = "secret"
		var give string = "hello world!"
		var expect string = "72069731bf291b463aecb218bc227abce3d403d76da67faef2d48d3cb43b2f54"
		get := HmacSha256(give, key)
		_assert.Equal(t, expect, get)
	}
}
func TestHmacSha512(t *testing.T) {
	{
		var key string = "secret"
		var give string = "hello world!"
		var expect string = "563069fb7c8512ffe6ced927289ac5e6f30a360c1099c61b62e3a91636a2563c95524ab5a0f4fe41f86e990a9f732dbf60d4f6c85761dafbd4953c24c758f936"
		get := HmacSha512(give, key)
		_assert.Equal(t, expect, get)
	}
}
func TestDecodeString(t *testing.T) {
	{
		hexStr := "48656c6c6f20576f726c64"
		expected := []byte("Hello World")
		result := DecodeString(hexStr)
		_assert.Equal(t, expected, result)
	}
	{
		hexStr := "48656c6c6fG25756f726c64"
		defer func() {
			r := recover()
			_assert.NotNil(t, r)
		}()
		DecodeString(hexStr)
	}
	{
		hexStr := ""
		expected := []byte{}
		result := DecodeString(hexStr)
		_assert.Equal(t, expected, result)
	}
}

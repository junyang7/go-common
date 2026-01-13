package _base64Format

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestEncode(t *testing.T) {
	{
		var expect string = "aGVsbG8gd29ybGQh"
		var give string = "hello world!"
		get := Encode(give)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "5oKo5aW977yM5Lit5Zu977yB"
		var give string = "您好，中国！"
		get := Encode(give)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "aGVsbG8g5Lit5Zu9"
		var give string = "hello 中国"
		get := Encode(give)
		_assert.Equal(t, expect, get)
	}
}
func TestDecode(t *testing.T) {
	{
		var expect string = "hello world!"
		var give string = "aGVsbG8gd29ybGQh"
		get := Decode(give)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "您好，中国！"
		var give string = "5oKo5aW977yM5Lit5Zu977yB"
		get := Decode(give)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "hello 中国"
		var give string = "aGVsbG8g5Lit5Zu9"
		get := Decode(give)
		_assert.Equal(t, expect, get)
	}
}
func TestEncodeEmpty(t *testing.T) {
	var give string = ""
	encoded := Encode(give)
	decoded := Decode(encoded)
	_assert.Equal(t, give, decoded)
}
func TestEncodeDecodeReversibility(t *testing.T) {
	testCases := []string{
		"",
		"a",
		"ab",
		"abc",
		"abcd",
		"hello world!",
		"您好，中国！",
		"mixed 混合 123 !@#$%",
		"包含特殊字符: + / = ? & # @",
		"very long text that needs to be encoded and decoded properly to ensure the implementation is correct for all lengths",
	}
	for _, original := range testCases {
		encoded := Encode(original)
		decoded := Decode(encoded)
		_assert.Equal(t, original, decoded)
	}
}
func TestURLSafeCharacters(t *testing.T) {
	testData := []string{
		"test+plus",
		"test/slash",
		"test=equal",
		string([]byte{0, 1, 2, 3, 4, 5}),
	}
	for _, data := range testData {
		encoded := Encode(data)
		_assert.NotContains(t, encoded, "+")
		_assert.NotContains(t, encoded, "/")
		_assert.NotContains(t, encoded, "=")
		decoded := Decode(encoded)
		_assert.Equal(t, data, decoded)
	}
}

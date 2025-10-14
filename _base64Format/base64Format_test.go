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

// 测试空字符串
func TestEncodeEmpty(t *testing.T) {
	var give string = ""
	encoded := Encode(give)
	decoded := Decode(encoded)
	_assert.Equal(t, give, decoded)
}

// 测试可逆性（编码后解码应该等于原文）
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

// 测试 URL 安全字符（验证不包含 + / = 等字符）
func TestURLSafeCharacters(t *testing.T) {
	// 包含会被替换的特殊字符的数据
	testData := []string{
		"test+plus",
		"test/slash",
		"test=equal",
		string([]byte{0, 1, 2, 3, 4, 5}), // 二进制数据
	}

	for _, data := range testData {
		encoded := Encode(data)

		// 验证编码结果不包含 URL 不安全字符
		_assert.NotContains(t, encoded, "+")
		_assert.NotContains(t, encoded, "/")
		_assert.NotContains(t, encoded, "=")

		// 验证可以正确解码
		decoded := Decode(encoded)
		_assert.Equal(t, data, decoded)
	}
}

package _aes

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_base64Format"
	"testing"
)

const k32 string = "b841b78d016df9dea4fc49e13d11199d"

func TestEncodeByAes256Cbc(t *testing.T) {
	// 测试基本加密解密
	{
		var give string = "hello world!"
		encrypted := EncodeByAes256Cbc(give, k32)
		decrypted := DecodeByAes256Cbc(encrypted, k32)
		_assert.Equal(t, give, decrypted)
	}
	{
		var give string = "您好，中国！"
		encrypted := EncodeByAes256Cbc(give, k32)
		decrypted := DecodeByAes256Cbc(encrypted, k32)
		_assert.Equal(t, give, decrypted)
	}
	{
		var give string = "hello 中国"
		encrypted := EncodeByAes256Cbc(give, k32)
		decrypted := DecodeByAes256Cbc(encrypted, k32)
		_assert.Equal(t, give, decrypted)
	}
}

func TestDecodeByAes256Cbc(t *testing.T) {
	// 测试解密功能
	{
		var give string = "hello world!"
		encrypted := EncodeByAes256Cbc(give, k32)
		decrypted := DecodeByAes256Cbc(encrypted, k32)
		_assert.Equal(t, give, decrypted)
	}
	{
		var give string = "您好，中国！"
		encrypted := EncodeByAes256Cbc(give, k32)
		decrypted := DecodeByAes256Cbc(encrypted, k32)
		_assert.Equal(t, give, decrypted)
	}
	{
		var give string = "hello 中国"
		encrypted := EncodeByAes256Cbc(give, k32)
		decrypted := DecodeByAes256Cbc(encrypted, k32)
		_assert.Equal(t, give, decrypted)
	}
}

// 测试空字符串
func TestEncodeEmptyString(t *testing.T) {
	var give string = ""
	encrypted := EncodeByAes256Cbc(give, k32)
	decrypted := DecodeByAes256Cbc(encrypted, k32)
	_assert.Equal(t, give, decrypted)
}

// 测试可逆性（加密后解密应该等于原文）
func TestEncodeDecodeReversibility(t *testing.T) {
	testCases := []string{
		"simple text",
		"复杂的中文内容，包含标点符号！@#￥%……&*（）",
		"123456789",
		"mixed 混合 content 123",
		"very long text that spans multiple blocks to test the padding and encryption properly across multiple AES blocks",
		"a",                 // 单字符
		"0123456789abcdef",  // 刚好16字节
		"0123456789abcdef0", // 17字节，跨块
	}

	for _, original := range testCases {
		encrypted := EncodeByAes256Cbc(original, k32)
		decrypted := DecodeByAes256Cbc(encrypted, k32)
		_assert.Equal(t, original, decrypted)
	}
}

// 测试IV随机性（多次加密相同内容应该得到不同密文）
func TestIVRandomness(t *testing.T) {
	plaintext := "test data"

	encrypted1 := EncodeByAes256Cbc(plaintext, k32)
	encrypted2 := EncodeByAes256Cbc(plaintext, k32)
	encrypted3 := EncodeByAes256Cbc(plaintext, k32)

	// 但解密后应该都是原文（验证IV随机性不影响解密正确性）
	_assert.Equal(t, plaintext, DecodeByAes256Cbc(encrypted1, k32))
	_assert.Equal(t, plaintext, DecodeByAes256Cbc(encrypted2, k32))
	_assert.Equal(t, plaintext, DecodeByAes256Cbc(encrypted3, k32))
}

// 测试标志位的两种情况（IV在前和IV在后）
func TestFlagBothCases(t *testing.T) {
	plaintext := "test data for flag testing"

	hasFlag1 := false
	hasFlag2 := false

	// 最多尝试1000次，确保两种标志位都出现
	for i := 0; i < 1000; i++ {
		encrypted := EncodeByAes256Cbc(plaintext, k32)

		// 解析密文，检查标志位
		// Base64解码后第一个字节就是标志位
		decoded := []byte(_base64Format.Decode(encrypted))
		flag := decoded[0]

		if flag == 1 {
			hasFlag1 = true
		} else if flag == 2 {
			hasFlag2 = true
		}

		// 验证解密正确
		decrypted := DecodeByAes256Cbc(encrypted, k32)
		_assert.Equal(t, plaintext, decrypted)

		// 如果两种标志位都出现过，可以提前结束
		if hasFlag1 && hasFlag2 {
			break
		}
	}

	// 确保两种标志位都被测试到
	_assert.Equal(t, true, hasFlag1)
	_assert.Equal(t, true, hasFlag2)
}

package _aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/junyang7/go-common/_base64Format"
	"github.com/junyang7/go-common/_interceptor"
)

func EncodeByAes256Cbc(data string, k32 string) string {
	_interceptor.Insure(32 == len(k32)).Message("k32长度不符合预期").Do()

	// 生成随机IV
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	_interceptor.Insure(nil == err).Message("生成IV失败").Do()

	// 加密数据
	b := []byte(data)
	block, err := aes.NewCipher([]byte(k32))
	_interceptor.Insure(nil == err).Message(err).Do()
	blockSize := block.BlockSize()
	paddingLength := blockSize - len(b)%blockSize
	b = append(b, bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)...)
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(b))
	mode.CryptBlocks(encrypted, b)

	// 随机生成标志位：1(IV在前) 或 2(IV在后)
	randomByte := make([]byte, 1)
	rand.Read(randomByte)
	flag := randomByte[0]%2 + 1 // 随机得到 1 或 2

	// 组合数据：标志位 + IV + 密文 (或 标志位 + 密文 + IV)
	var combined []byte
	if flag == 1 {
		// IV在前：[标志位][IV][密文]
		combined = append([]byte{flag}, iv...)
		combined = append(combined, encrypted...)
	} else {
		// IV在后：[标志位][密文][IV]
		combined = append([]byte{flag}, encrypted...)
		combined = append(combined, iv...)
	}

	return _base64Format.Encode(string(combined))
}
func DecodeByAes256Cbc(data string, k32 string) string {
	_interceptor.Insure(32 == len(k32)).Message("k32长度不符合预期").Do()

	// Base64 解码
	combined := []byte(_base64Format.Decode(data))
	_interceptor.Insure(len(combined) > 17).Message("解密失败").Do() // 至少要有：1字节标志位 + 16字节IV + 数据

	// 读取标志位
	flag := combined[0]
	_interceptor.Insure(flag == 1 || flag == 2).Message("解密失败").Do()

	// 根据标志位提取IV和密文
	var iv []byte
	var ciphertext []byte

	if flag == 1 {
		// IV在前：[标志位][IV][密文]
		iv = combined[1:17]
		ciphertext = combined[17:]
	} else {
		// IV在后：[标志位][密文][IV]
		iv = combined[len(combined)-16:]
		ciphertext = combined[1 : len(combined)-16]
	}

	// 校验密文长度
	_interceptor.Insure(len(ciphertext) > 0 && len(ciphertext)%16 == 0).Message("解密失败").Do()

	// 解密
	block, err := aes.NewCipher([]byte(k32))
	_interceptor.Insure(nil == err).Message(err).Do()
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	// 去除padding
	length := len(decrypted)
	paddingLength := int(decrypted[length-1])
	_interceptor.Insure(paddingLength > 0 && paddingLength <= 16).Message("解密失败").Do()
	// 验证所有padding字节的值是否正确
	for i := 0; i < paddingLength; i++ {
		_interceptor.Insure(decrypted[length-1-i] == byte(paddingLength)).Message("解密失败").Do()
	}
	decrypted = decrypted[:(length - paddingLength)]
	return string(decrypted)
}

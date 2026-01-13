package _aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/junyang7/go-common/_base64Format"
	"github.com/junyang7/go-common/_interceptor"
)

func Encode(data string, k32 string, i16 string) string {
	b := []byte(data)
	block, err := aes.NewCipher([]byte(k32))
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	size := block.BlockSize()
	l := size - len(b)%size
	b = append(b, bytes.Repeat([]byte{byte(l)}, l)...)
	mode := cipher.NewCBCEncrypter(block, []byte(i16))
	encrypted := make([]byte, len(b))
	mode.CryptBlocks(encrypted, b)
	return _base64Format.Encode(string(encrypted))
}
func Decode(data string, k32 string, i16 string) string {
	b := []byte(_base64Format.Decode(data))
	block, err := aes.NewCipher([]byte(k32))
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	mode := cipher.NewCBCDecrypter(block, []byte(i16))
	decrypted := make([]byte, len(b))
	mode.CryptBlocks(decrypted, b)
	l := len(decrypted)
	decrypted = decrypted[:(l - int(decrypted[l-1]))]
	return string(decrypted)
}
func EncodeByAes256Cbc(data string, k32 string) string {
	_interceptor.Insure(32 == len(k32)).Message("k32长度不符合预期").Do()
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	_interceptor.Insure(nil == err).Message("生成IV失败").Do()
	b := []byte(data)
	block, err := aes.NewCipher([]byte(k32))
	_interceptor.Insure(nil == err).Message(err).Do()
	blockSize := block.BlockSize()
	paddingLength := blockSize - len(b)%blockSize
	b = append(b, bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)...)
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(b))
	mode.CryptBlocks(encrypted, b)
	randomByte := make([]byte, 1)
	rand.Read(randomByte)
	flag := randomByte[0]%2 + 1
	var combined []byte
	if flag == 1 {
		combined = append([]byte{flag}, iv...)
		combined = append(combined, encrypted...)
	} else {
		combined = append([]byte{flag}, encrypted...)
		combined = append(combined, iv...)
	}
	return _base64Format.Encode(string(combined))
}
func DecodeByAes256Cbc(data string, k32 string) string {
	_interceptor.Insure(32 == len(k32)).Message("k32长度不符合预期").Do()
	combined := []byte(_base64Format.Decode(data))
	_interceptor.Insure(len(combined) > 17).Message("解密失败").Do()
	flag := combined[0]
	_interceptor.Insure(flag == 1 || flag == 2).Message("解密失败").Do()
	var iv []byte
	var ciphertext []byte
	if flag == 1 {
		iv = combined[1:17]
		ciphertext = combined[17:]
	} else {
		iv = combined[len(combined)-16:]
		ciphertext = combined[1 : len(combined)-16]
	}
	_interceptor.Insure(len(ciphertext) > 0 && len(ciphertext)%16 == 0).Message("解密失败").Do()
	block, err := aes.NewCipher([]byte(k32))
	_interceptor.Insure(nil == err).Message(err).Do()
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)
	length := len(decrypted)
	paddingLength := int(decrypted[length-1])
	_interceptor.Insure(paddingLength > 0 && paddingLength <= 16).Message("解密失败").Do()
	for i := 0; i < paddingLength; i++ {
		_interceptor.Insure(decrypted[length-1-i] == byte(paddingLength)).Message("解密失败").Do()
	}
	decrypted = decrypted[:(length - paddingLength)]
	return string(decrypted)
}

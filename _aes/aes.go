package _aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
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

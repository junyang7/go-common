package _hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
)

func Md5(data string) string {
	h := md5.New()
	_, err := h.Write([]byte(data))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHashHashWrite).
		Message(err).
		Do()
	return hex.EncodeToString(h.Sum(nil))
}
func Sha1(data string) string {
	h := sha1.New()
	_, err := h.Write([]byte(data))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHashHashWrite).
		Message(err).
		Do()
	return hex.EncodeToString(h.Sum(nil))
}
func Sha256(data string) string {
	h := sha256.New()
	_, err := h.Write([]byte(data))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHashHashWrite).
		Message(err).
		Do()
	return hex.EncodeToString(h.Sum(nil))
}
func Sha512(data string) string {
	h := sha512.New()
	_, err := h.Write([]byte(data))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHashHashWrite).
		Message(err).
		Do()
	return hex.EncodeToString(h.Sum(nil))
}
func HmacSha1(data string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	_, err := h.Write([]byte(data))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHashHashWrite).
		Message(err).
		Do()
	return hex.EncodeToString(h.Sum(nil))
}
func HmacSha256(data string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	_, err := h.Write([]byte(data))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHashHashWrite).
		Message(err).
		Do()
	return hex.EncodeToString(h.Sum(nil))
}

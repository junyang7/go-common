package _hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Md5(data string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(data))
	if nil != err {
		panic(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
func Sha1(data string) string {
	hash := sha1.New()
	_, err := hash.Write([]byte(data))
	if nil != err {
		panic(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
func Sha256(data string) string {
	hash := sha256.New()
	_, err := hash.Write([]byte(data))
	if nil != err {
		panic(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
func Sha512(data string) string {
	hash := sha512.New()
	_, err := hash.Write([]byte(data))
	if nil != err {
		panic(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

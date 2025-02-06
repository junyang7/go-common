package _rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/junyang7/go-common/_base64Format"
	"github.com/junyang7/go-common/_interceptor"
)

type Rsa struct {
	Pri string
	Pub string
}

func Generate() *Rsa {
	this := &Rsa{}
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	{
		b, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		block := pem.Block{Type: "PRIVATE KEY", Bytes: b}
		this.Pri = string(pem.EncodeToMemory(&block))
	}
	publicKey := privateKey.PublicKey
	{
		b, err := x509.MarshalPKIXPublicKey(&publicKey)
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		block := pem.Block{Type: "PUBLIC KEY", Bytes: b}
		this.Pub = string(pem.EncodeToMemory(&block))
	}
	return this
}
func Encode(data string, pub string) string {
	block, _ := pem.Decode([]byte(pub))
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	pubKey := pubInterface.(*rsa.PublicKey)
	b, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(data))
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return _base64Format.Encode(string(b))
}
func Decode(data string, pri string) string {
	block, _ := pem.Decode([]byte(pri))
	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	priKey := priInterface.(*rsa.PrivateKey)
	b, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, []byte(_base64Format.Decode(data)))
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return string(b)
}
func Sign(hashed []byte, pri string) string {
	block, _ := pem.Decode([]byte(pri))
	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	priKey := priInterface.(*rsa.PrivateKey)
	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed)
	if err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
	return string(signature)
}
func Verify(hashed []byte, sign string, pub string) bool {
	block, _ := pem.Decode([]byte(pub))
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	pubKey := pubInterface.(*rsa.PublicKey)
	if err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed, []byte(sign)); nil != err {
		return false
	}
	return true
}

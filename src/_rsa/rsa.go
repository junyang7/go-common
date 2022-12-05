package _rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/junyang7/go-common/src/_base64Format"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
)

type Rsa struct {
	Pri string
	Pub string
}

func Generate() *Rsa {
	this := &Rsa{}
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrRsaGenerateKey).
		Message(err).
		Do()
	{
		b, err := x509.MarshalPKCS8PrivateKey(privateKey)
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrX509MarshalPKCS8PrivateKey).
			Message(err).
			Do()
		block := pem.Block{Type: "PRIVATE KEY", Bytes: b}
		this.Pri = string(pem.EncodeToMemory(&block))
	}
	publicKey := privateKey.PublicKey
	{
		b, err := x509.MarshalPKIXPublicKey(&publicKey)
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrX509MarshalPKIXPublicKey).
			Message(err).
			Do()
		block := pem.Block{Type: "PUBLIC KEY", Bytes: b}
		this.Pub = string(pem.EncodeToMemory(&block))
	}
	return this
}
func Encode(data string, pub string) string {
	block, _ := pem.Decode([]byte(pub))
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	pubKey := pubInterface.(*rsa.PublicKey)
	b, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(data))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrRsaEncryptPKCS1v15).
		Message(err).
		Do()
	return _base64Format.Encode(string(b))
}
func Decode(data string, pri string) string {
	block, _ := pem.Decode([]byte(pri))
	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrX509ParsePKCS8PrivateKey).
		Message(err).
		Do()
	priKey := priInterface.(*rsa.PrivateKey)
	b, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, []byte(_base64Format.Decode(data)))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrRsaDecryptPKCS1v15).
		Message(err).
		Do()
	return string(b)
}

package _rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/junyang7/go-common/src/_base64Format"
)

type Rsa struct {
	Pri string
	Pub string
}

func Generate() *Rsa {
	this := &Rsa{}
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	{
		b, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if nil != err {
			panic(err)
		}
		block := pem.Block{Type: "PRIVATE KEY", Bytes: b}
		this.Pri = string(pem.EncodeToMemory(&block))
	}
	publicKey := privateKey.PublicKey
	{
		b, err := x509.MarshalPKIXPublicKey(&publicKey)
		if err != nil {
			panic(err)
		}
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
	if err != nil {
		panic(err)
	}
	return _base64Format.Encode(string(b))
}
func Decode(data string, pri string) string {
	block, _ := pem.Decode([]byte(pri))
	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	priKey := priInterface.(*rsa.PrivateKey)
	if err != nil {
		panic(err)
	}
	b, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, []byte(_base64Format.Decode(data)))
	if err != nil {
		panic(err)
	}
	return string(b)
}

package _rsa

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_hash"
	"testing"
)

func TestGenerate(t *testing.T) {
	{
		rsaObj := Generate()
		_assert.NotNil(t, rsaObj.Pri)
		_assert.NotNil(t, rsaObj.Pub)
		_assert.True(t, len(rsaObj.Pri) > 0)
		_assert.True(t, len(rsaObj.Pub) > 0)
	}
}
func TestEncodeAndDecode(t *testing.T) {
	{
		rsaObj := Generate()
		originalData := "Hello, RSA!"
		encryptedData := Encode(originalData, rsaObj.Pub)
		_assert.NotNil(t, encryptedData)
		_assert.True(t, len(encryptedData) > 0)
		decryptedData := Decode(encryptedData, rsaObj.Pri)
		_assert.NotNil(t, decryptedData)
		_assert.Equal(t, originalData, decryptedData)
	}
}
func TestSignAndVerify(t *testing.T) {
	{
		rsaObj := Generate()
		message := "hello world"
		hashed := _hash.DecodeString(_hash.Sha256(message))
		signature := Sign(hashed, rsaObj.Pri)
		_assert.NotNil(t, signature)
		_assert.True(t, len(signature) > 0)
		isVerified := Verify(hashed, signature, rsaObj.Pub)
		_assert.True(t, isVerified)
	}
}

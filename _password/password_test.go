package _password

import (
	"github.com/junyang7/go-common/_assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHash(t *testing.T) {
	{
		password := "mySuperSecretPassword"
		hashedPassword := Hash(password)
		_assert.NotEmpty(t, hashedPassword)
		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		_assert.Nil(t, err)
	}
}
func TestVerify(t *testing.T) {
	{
		password := "mySuperSecretPassword"
		hashedPassword := Hash(password)
		isValid := Verify(hashedPassword, password)
		_assert.True(t, isValid)
	}
	{
		password := "mySuperSecretPassword"
		hashedPassword := Hash(password)
		invalidPassword := "wrongPassword"
		isValid := Verify(hashedPassword, invalidPassword)
		_assert.False(t, isValid)
	}
}

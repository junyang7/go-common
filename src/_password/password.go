package _password

import (
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_interceptor.Insure(nil != err).
		CodeMessage(_codeMessage.ErrBcryptGenerateFromPassword).
		Message(err.Error()).
		Do()
	return string(hash)
}

func Verify(passwordHashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
	return nil == err
}

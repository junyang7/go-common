package _password

import "golang.org/x/crypto/bcrypt"

func Hash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if nil != err {
		panic(err)
	}
	return string(hash)
}

func Verify(passwordHashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
	return nil == err
}

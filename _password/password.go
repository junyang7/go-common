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
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password)); nil != err {
		return false
	}
	return true
}

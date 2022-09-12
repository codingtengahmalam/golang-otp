package helper

import "golang.org/x/crypto/bcrypt"

func NewPassword(password string) (string, error) {
	bytePass := []byte(password)

	hashPass, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

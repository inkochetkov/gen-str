package gen

import "golang.org/x/crypto/bcrypt"

func GenPassword(password string) (string, error) {

	hashedPassByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPass := string(hashedPassByte)

	return hashedPass, nil

}

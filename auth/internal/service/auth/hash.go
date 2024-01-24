package auth

import "golang.org/x/crypto/bcrypt"

func hashPassword(psw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.DefaultCost)
	return string(bytes), err
}
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

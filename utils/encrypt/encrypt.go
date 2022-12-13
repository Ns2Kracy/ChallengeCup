package encrypt

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encrypts the password using bcrypt
func EncryptPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// ComparePassword compares the password with the hash
func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
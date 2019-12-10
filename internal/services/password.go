package services

import "golang.org/x/crypto/bcrypt"

// Password service
type Password struct {
	cost int
}

// NewPassword initialize the service
func NewPassword() *Password {
	return &Password{cost: 14}
}

// SecurePassword generate a hash to store into the database
func (service *Password) SecurePassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), service.cost)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// CheckPasswordHash check a hash and a raw password
func (service *Password) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

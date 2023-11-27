package cryptography

import "golang.org/x/crypto/bcrypt"

type BcryptAdapter struct{}

func NewBcryptAdapter() *BcryptAdapter {
	return &BcryptAdapter{}
}

func (b *BcryptAdapter) Hash(password string, salt int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

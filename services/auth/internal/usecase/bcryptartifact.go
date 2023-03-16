package usecase

import (
	domain "grpcrest/services/auth/domain/usecase"

	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
}

func (*Bcrypt) CompareHashAndPassword(hashed []byte, pass []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, pass)
}

func (*Bcrypt) GenerateFromPassword(pass []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pass, cost)
}

func NewBcrypt() domain.Bcrypt_ {
	return &Bcrypt{}
}

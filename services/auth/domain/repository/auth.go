package repository

import (
	"grpcrest/services/auth/domain/entity"
)

type AuthRepository interface {
	CreateAuth(auth *entity.Auth) error
	DeleteAuth(auth *entity.Auth) error
	FindPasswordByEmail(auth *entity.Auth) (string, error)
}

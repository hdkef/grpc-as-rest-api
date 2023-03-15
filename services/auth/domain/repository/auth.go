package repository

import (
	"grpcrest/services/auth/domain/entity"
)

type AuthRepository interface {
	UpdateAuth(auth *entity.Auth) error
	CreateAuth(auth *entity.Auth) error
	DeleteAuth(auth *entity.Auth) error
	FindUserIdPasswordByEmail(auth *entity.Auth) (string, string, error)
	IsExistEmail(email *string) bool
}

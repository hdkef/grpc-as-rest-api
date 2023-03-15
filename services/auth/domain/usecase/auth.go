package usecase

import "grpcrest/services/auth/domain/entity"

type AuthUseCase interface {
	CreateAuth(auth *entity.Auth) error
	DeleteAuth(auth *entity.Auth) error
	FindPasswordByEmail(auth *entity.Auth) (string, error)
}

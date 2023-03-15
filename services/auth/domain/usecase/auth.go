package usecase

import "grpcrest/services/auth/domain/entity"

type AuthUseCase interface {
	UpdateAuth(auth *entity.Auth) error
	CreateAuth(auth *entity.Auth) error
	DeleteAuth(auth *entity.Auth) error
	FindPasswordByEmail(auth *entity.Auth) (string, error)
}

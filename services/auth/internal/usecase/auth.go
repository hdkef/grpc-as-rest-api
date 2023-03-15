package usecase

import (
	"grpcrest/services/auth/domain/entity"
	authRepo "grpcrest/services/auth/domain/repository"
	domain "grpcrest/services/auth/domain/usecase"
)

type AuthArtifact struct {
	authRepo authRepo.AuthRepository
}

// CreateAuth implements usecase.AuthUseCase
func (a *AuthArtifact) CreateAuth(auth *entity.Auth) error {
	return a.authRepo.CreateAuth(auth)
}

// DeleteAuth implements usecase.AuthUseCase
func (a *AuthArtifact) DeleteAuth(auth *entity.Auth) error {
	return a.authRepo.DeleteAuth(auth)
}

// FindPasswordByEmail implements usecase.AuthUseCase
func (a *AuthArtifact) FindPasswordByEmail(auth *entity.Auth) (string, error) {
	return a.authRepo.FindPasswordByEmail(auth)
}

func NewAuthUsecase(repo authRepo.AuthRepository) domain.AuthUseCase {
	return &AuthArtifact{
		authRepo: repo,
	}
}

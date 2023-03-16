package usecase

import "grpcrest/services/user/domain/entity"

// GetAll implements usecase.UserUsecase
func (u *UserArtifact) GetAll(lastID *int, limit *int) ([]entity.User, error) {
	return u.userrepo.GetAll(lastID, limit)
}

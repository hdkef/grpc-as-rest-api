package usecase

import (
	"errors"
	"grpcrest/services/user/domain/entity"
)

// GetAll implements usecase.UserUsecase
func (u *UserArtifact) GetAll(page *int, limit *int) ([]entity.User, error) {
	if *page <= 0 || *limit <= 0 {
		return nil, errors.New("limit or page query params is invalid")
	}
	return u.userrepo.GetAll(page, limit)
}

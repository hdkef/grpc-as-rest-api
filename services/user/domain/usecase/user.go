package usecase

import "grpcrest/services/user/domain/entity"

type UserUsecase interface {
	Create(*entity.User) error
	Update(*entity.User) error
	Delete(*entity.User) error
	GetAll(page *int, limit *int) ([]entity.User, error)
}

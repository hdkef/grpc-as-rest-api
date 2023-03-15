package repository

import "grpcrest/services/user/domain/entity"

type UserRepository interface {
	Create(*entity.User) error
	Update(*entity.User) error
	Delete(*entity.User) error
	GetAll(lastID *int, limit *int) ([]entity.User, error)
	IsExistEmail(email *string) bool
}

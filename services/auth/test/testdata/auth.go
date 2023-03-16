package testdata

import (
	"grpcrest/services/auth/domain/entity"

	"github.com/bxcodec/faker/v4"
)

func NewAuthEntity() *entity.Auth {
	return &entity.Auth{
		ID:       faker.UUIDDigit(),
		UserID:   faker.UUIDDigit(),
		Password: faker.Password(),
		Email:    faker.Email(),
	}
}

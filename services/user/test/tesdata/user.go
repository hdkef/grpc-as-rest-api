package tesdata

import (
	"grpcrest/services/user/domain/entity"

	"github.com/bxcodec/faker/v4"
)

func NewUserEntity() entity.User {
	return entity.User{
		ID:       faker.UUIDDigit(),
		Name:     faker.Name(),
		Email:    faker.Email(),
		Address:  faker.Word(),
		Password: faker.Password(),
	}
}

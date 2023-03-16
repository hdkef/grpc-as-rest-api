package mapper

import (
	"grpcrest/services/user/domain/entity"
	"grpcrest/services/user/internal/repository/models"
)

func ToModels(e *entity.User) models.User {
	return models.User{
		ID:       e.ID,
		Email:    e.Email,
		Name:     e.Name,
		Address:  e.Address,
		Password: e.Password,
	}
}

func ToEntity(m *models.User) entity.User {
	return entity.User{
		ID:       m.ID,
		Email:    m.Email,
		Name:     m.Name,
		Address:  m.Address,
		Password: m.Password,
	}
}

func ToEntities(m []models.User) []entity.User {
	var data []entity.User
	for _, v := range m {
		data = append(data, ToEntity(&v))
	}
	return data
}

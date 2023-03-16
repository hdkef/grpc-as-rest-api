package mapper

import (
	"grpcrest/services/auth/domain/entity"
	"grpcrest/services/auth/internal/repository/models"
)

func ToModels(e *entity.Auth) models.Auth {
	return models.Auth{
		ID:       e.ID,
		UserID:   e.UserID,
		Email:    e.Email,
		Password: e.Password,
	}
}

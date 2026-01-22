package mappers

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db/models"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/entity"
)

func ToUserEntity(model models.UserModel) entity.User {

	return entity.User{
		Id:       model.Id,
		Email:    model.Email,
		Name:     model.Name,
		Password: model.Password,
		TenantId: model.TenantId,
	}
}

func ToUserModel(user *entity.User) models.UserModel {

	return models.UserModel{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		TenantId: user.TenantId,
	}
}

func ToUserEntities(models []models.UserModel) []entity.User {
	users := make([]entity.User, len(models))
	for i, model := range models {
		users[i] = ToUserEntity(model)
	}

	return users
}

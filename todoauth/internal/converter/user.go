package converter

import (
	"todoauth/internal/dto"
	"todoauth/internal/model"
	"todoauth/internal/modeldb"
)

func UserDBModelToModel(user *modeldb.User) *model.User {
	return &model.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Dates:    DatesDBModelToModel(user.Dates),
	}
}

func UserModelToDTOResponse(user *model.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

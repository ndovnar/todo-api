package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todolib/db"
	"todolib/ginhelper"

	"todoauth/internal/converter"
	"todoauth/internal/dto"
	"todoauth/internal/service"
)

type Users struct {
	userService service.UserService
}

func NewUsers(userService service.UserService) *Users {
	return &Users{
		userService: userService,
	}
}

func (h *Users) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	user, err := h.userService.CreateUser(ctx, &service.CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if err == db.ErrDuplicateKey {
			ctx.Error(ginhelper.NewHttpErrorWithDescription(http.StatusUnprocessableEntity, "user already exists"))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.JSON(http.StatusOK, converter.UserModelToDTOResponse(user))
}

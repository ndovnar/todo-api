package handlers

import (
	"auth/internal/model"
	"auth/internal/service"
	"lib/db"
	"lib/ginhelper"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var req registerUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	user, err := h.userService.CreateUser(ctx, req.Email, req.Password)
	if err != nil {
		if err == db.ErrDuplicateKey {
			ctx.Error(ginhelper.NewHttpErrorWithDescription(http.StatusUnprocessableEntity, "user already exists"))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	resp := newUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type registerUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func newUserResponse(user *model.User) *userResponse {
	return &userResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

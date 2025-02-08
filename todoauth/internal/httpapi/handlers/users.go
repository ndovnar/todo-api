package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todolib/db"
	"todolib/ginhelper"

	"todoauth/internal/model"
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
	var req createUserRequest
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

	resp := newCreateUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func newCreateUserResponse(user *model.User) *createUserResponse {
	return &createUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

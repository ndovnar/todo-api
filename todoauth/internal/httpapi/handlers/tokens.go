package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todoauth/internal/converter"
	"todoauth/internal/dto"
	"todoauth/internal/service"

	"todolib/auth"
	"todolib/ginhelper"
)

type Tokens struct {
	authService service.AuthService
}

func NewTokens(authService service.AuthService) *Tokens {
	return &Tokens{
		authService: authService,
	}
}

func (h *Tokens) HandleLogin(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	tokenPair, err := h.authService.Login(ctx, &service.LoginParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusUnauthorized))
		return
	}

	ctx.JSON(http.StatusOK, converter.TokenPairModelToDTOResponse(tokenPair))
}

func (h *Tokens) HandleLogout(ctx *gin.Context) {
	claims := auth.GetClaimsFromContext(ctx)

	err := h.authService.Logout(ctx, claims.ID)
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusUnauthorized))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Tokens) HandleRenewAccessToken(ctx *gin.Context) {
	var req dto.RenewTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	accessToken, err := h.authService.RenewAccessToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusUnauthorized))
		return
	}

	ctx.JSON(http.StatusOK, converter.AccessTokenToDTOResponse(accessToken))
}

func (h *Tokens) HandleRenewRefreshToken(ctx *gin.Context) {
	var req dto.RenewTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	tokenPair, err := h.authService.RenewRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusUnauthorized))
		return
	}

	ctx.JSON(http.StatusOK, converter.TokenPairModelToDTOResponse(tokenPair))
}

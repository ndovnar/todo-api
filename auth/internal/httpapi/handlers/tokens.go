package handlers

import (
	"auth/internal/service"
	"lib/auth"
	"lib/ginhelper"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	accessToken, refreshToken, err := h.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusUnauthorized))
		return
	}

	resp := newTokenPairResponse(accessToken, refreshToken)
	ctx.JSON(http.StatusOK, resp)
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
	var req renewTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	accessToken, err := h.authService.RenewAccessToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusUnauthorized))
		return
	}

	resp := newRenewAccessTokenResponse(accessToken)
	ctx.JSON(http.StatusOK, resp)
}

func (h *Tokens) HandleRenewRefreshToken(ctx *gin.Context) {
	var req renewTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	accessToken, refreshToken, err := h.authService.RenewRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusUnauthorized))
		return
	}

	resp := newTokenPairResponse(accessToken, refreshToken)
	ctx.JSON(http.StatusOK, resp)
}

type renewTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

func newRenewAccessTokenResponse(accessToken string) *renewAccessTokenResponse {
	return &renewAccessTokenResponse{
		AccessToken: accessToken,
	}
}

type tokenPairResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func newTokenPairResponse(accessToken string, refreshToken string) *tokenPairResponse {
	return &tokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

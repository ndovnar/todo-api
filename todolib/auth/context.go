package auth

import "github.com/gin-gonic/gin"

const claimsContextKey = "authorization_payload"

func GetClaimsFromContext(ctx *gin.Context) *Claims {
	return ctx.MustGet(claimsContextKey).(*Claims)
}

func SetClaimsToContext(ctx *gin.Context, claims *Claims) {
	ctx.Set(claimsContextKey, claims)
}

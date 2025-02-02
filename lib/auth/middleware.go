package auth

import (
	"crypto/rsa"
	"fmt"
	"lib/ginhelper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func AuthMiddleware(publicKey *rsa.PublicKey) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			ctx.Error(ginhelper.NewHttpErrorWithDescription(http.StatusUnauthorized, "authorization header is not provided"))
			ctx.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.Error(ginhelper.NewHttpErrorWithDescription(http.StatusUnauthorized, "invalid authorization header format"))
			ctx.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.Error(ginhelper.NewHttpErrorWithDescription(http.StatusUnauthorized, fmt.Sprintf("unsupported authorization type %s", authorizationType)))
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		claims, err := VerifyToken(accessToken, publicKey)
		if err != nil {
			ctx.Error(ginhelper.NewHttpError(http.StatusUnauthorized))
			ctx.Abort()
			return
		}

		SetClaimsToContext(ctx, claims)
		ctx.Next()
	}
}

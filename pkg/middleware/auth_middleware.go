package middleware

import (
	"learnyscape-backend-mono/pkg/constant"
	"learnyscape-backend-mono/pkg/httperror"
	jwtutil "learnyscape-backend-mono/pkg/util/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwt jwtutil.JWTUtil, allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := parseAccessToken(ctx)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		claims, err := jwt.ParseAccess(accessToken)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		if claims.UserID == 0 {
			ctx.Error(httperror.NewUnauthorizedError())
			ctx.Abort()
			return
		}

		// TODO: change to forbidden error
		if !isRoleAllowed(claims.Role, allowedRoles...) {
			ctx.Error(httperror.NewUnauthorizedError())
			ctx.Abort()
			return
		}

		ctx.Set(constant.CTX_USER_ID, claims.UserID)
		ctx.Set(constant.CTX_USER_ROLE, claims.Role)
		ctx.Next()
	}
}

func parseAccessToken(ctx *gin.Context) (string, error) {
	accessToken := ctx.GetHeader("Authorization")
	if accessToken == "" || len(accessToken) < 7 {
		return "", httperror.NewUnauthorizedError()
	}

	splitToken := strings.Split(accessToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", httperror.NewUnauthorizedError()
	}

	return splitToken[1], nil
}

func isRoleAllowed(role string, allowedRoles ...string) bool {
	if len(allowedRoles) == 0 {
		return true
	}

	for _, allowedRoles := range allowedRoles {
		if role == allowedRoles {
			return true
		}
	}

	return false
}

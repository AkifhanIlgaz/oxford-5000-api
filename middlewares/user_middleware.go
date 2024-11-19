package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
	tokenService services.TokenService
}

func NewUserMiddleware(tokenService services.TokenService) UserMiddleware {
	return UserMiddleware{
		tokenService: tokenService,
	}
}

func parseIdTokenFromHeader(header http.Header) (string, error) {
	authorizationHeader := header.Get(api.AuthHeader)
	fields := strings.Fields(authorizationHeader)

	if len(fields) == 0 {
		return "", fmt.Errorf("authorization header is empty")
	}
	if len(fields) > 2 {
		return "", fmt.Errorf("authorization header is invalid")
	}
	if fields[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header scheme")
	}

	return fields[1], nil
}

func (m UserMiddleware) AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token from header
		token, err := parseIdTokenFromHeader(ctx.Request.Header)
		if err != nil {
			response.WithError(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		// Parse and validate token
		uid, err := m.tokenService.ParseToken("access", token)
		if err != nil {
			response.WithError(ctx, http.StatusUnauthorized, message.InvalidToken)
			return
		}

		ctx.Set(api.ContextUid, uid)
		ctx.Next()
	}
}

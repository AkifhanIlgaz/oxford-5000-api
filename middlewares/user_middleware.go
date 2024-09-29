package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
	userService services.UserService
}

func NewUserMiddleware(userService services.UserService) UserMiddleware {
	return UserMiddleware{
		userService: userService,
	}
}

func (middleware UserMiddleware) GetUserFromIdToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken, err := parseIdTokenFromHeader(ctx.Request.Header)
		if err != nil {
			log.Println(err.Error())
			response.WithError(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := middleware.userService.GetUserFromIdToken(idToken)
		if err != nil {
			log.Println(err.Error())
			response.WithError(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set(api.ContextUser, user)
		ctx.Next()
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

func GetUserFromContext(ctx *gin.Context) (*auth.UserRecord, error) {
	ctxUser, exists := ctx.Get(api.ContextUser)
	if !exists {
		return nil, errors.New(message.UserNotFound)
	}

	user, ok := ctxUser.(*auth.UserRecord)
	if !ok {
		return nil, errors.New(message.UnableParseUser)
	}

	return user, nil
}

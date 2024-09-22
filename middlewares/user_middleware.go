package middlewares

import (
	"github.com/AkifhanIlgaz/dictionary-api/services"
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

	}
}

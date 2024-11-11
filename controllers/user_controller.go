package controllers

import (
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/gin-gonic/gin"
)

const UserPath = "/user"

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (controller UserController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(UserPath)

    // Create, delete API KEY
    // Get API key usage
    _ = router
}

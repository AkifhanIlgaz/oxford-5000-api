package controllers

import (
	"github.com/gin-gonic/gin"
)

const AuthPath string = "/auth"

type AuthController struct {
}

func NewAuthController() AuthController {
	return AuthController{}
}

func (controller AuthController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(AuthPath)

	// Login
	// Register
	// Refresh Access Token
	_ = router
}

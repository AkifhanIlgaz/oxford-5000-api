package controllers

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

const AuthPath string = "/auth"

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (controller AuthController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(AuthPath)

	router.POST("/signup", controller.Signup)
	router.POST("/signin", controller.Signin)
	router.POST("/refresh", controller.Refresh)
}

func (controller AuthController) Signup(ctx *gin.Context) {
	var req models.SignupRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, message.MissingField)
		return
	}

	if err := req.Validate(); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	uid, err := controller.authService.Create(req)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}


	// Create Tokens
}

func (controller AuthController) Signin(ctx *gin.Context) {

}

func (controller AuthController) Refresh(ctx *gin.Context) {

}

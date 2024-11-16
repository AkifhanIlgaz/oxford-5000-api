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
const ApiKeyPath string = "/api-key"

type AuthController struct {
	authService  services.AuthService
	tokenService services.TokenService
}

func NewAuthController(authService services.AuthService, tokenService services.TokenService) AuthController {
	return AuthController{
		authService:  authService,
		tokenService: tokenService,
	}
}

func (controller AuthController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(AuthPath)

	router.POST("/signup", controller.Signup)
	router.POST("/signin", controller.Signin)
	router.POST("/refresh", controller.Refresh)

	apiKey := router.Group(ApiKeyPath)

	apiKey.GET("/", controller.GetApiKey)
	apiKey.POST("/", controller.CreateApiKey)
	apiKey.DELETE("/", controller.DeleteApiKey)

}

func (controller AuthController) Signup(ctx *gin.Context) {
	var req models.AuthRequest

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

	tokens, err := controller.tokenService.CreateTokens(uid.Hex())
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, "user created", tokens)
}

func (controller AuthController) Signin(ctx *gin.Context) {
	var req models.AuthRequest

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

	uid, err := controller.authService.AuthenticateUser(req)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	tokens, err := controller.tokenService.CreateTokens(uid.Hex())
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, "logged in", tokens)
}

func (controller AuthController) Refresh(ctx *gin.Context) {
	var req models.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusBadRequest, message.MissingField)
		return
	}

	uid, err := controller.tokenService.ParseToken("refresh", req.RefreshToken)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	tokens, err := controller.tokenService.CreateTokens(uid)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.WithSuccess(ctx, http.StatusOK, "tokens refreshed", tokens)
}

func (controller AuthController) GetApiKey(ctx *gin.Context) {}

func (controller AuthController) CreateApiKey(ctx *gin.Context) {}

func (controller AuthController) DeleteApiKey(ctx *gin.Context) {}

package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/dictionary-api/middlewares"
	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserPath = "/user"
const ApiKeyPath = "/api-key"

type UserController struct {
	userService    services.UserService
	userMiddleware middlewares.UserMiddleware
}

func NewUserController(userService services.UserService, userMiddleware middlewares.UserMiddleware) UserController {
	return UserController{
		userService:    userService,
		userMiddleware: userMiddleware,
	}
}

func (controller UserController) SetupRoutes(rg *gin.RouterGroup) {
	router := rg.Group(UserPath)

	apiKey := router.Group(ApiKeyPath)
	apiKey.Use(controller.userMiddleware.AuthenticateUser())

	apiKey.GET("/", controller.GetApiKey)
	apiKey.POST("/", controller.CreateApiKey)
	apiKey.DELETE("/", controller.DeleteApiKey)
}

func (controller UserController) GetApiKey(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)

	apiKey, err := controller.userService.GetApiKey(uid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			response.WithError(ctx, http.StatusNotFound, "No API key found for this user")
			return
		}
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.ApiKeyRetrieved, models.APIKeyResponse{
		APIKey: apiKey,
	})
}

func (controller UserController) CreateApiKey(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)
	name := ctx.Query(api.NameParam)

	apiKey, err := controller.userService.CreateApiKey(uid, name)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
		return
	}

	response.WithSuccess(ctx, http.StatusCreated, message.ApiKeyCreated, models.APIKeyResponse{
		APIKey: apiKey,
	})
}

func (controller UserController) DeleteApiKey(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)

	if err := controller.userService.DeleteApiKey(uid); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.ApiKeyDeleted, nil)
}

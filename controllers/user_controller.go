package controllers

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

const UserPath = "/user"
const ApiKeyPath = "/api-key"

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

	apiKey := router.Group(ApiKeyPath)
	apiKey.GET("/:uid", controller.GetApiKey)
	apiKey.POST("/", controller.CreateApiKey)
	apiKey.DELETE("/", controller.DeleteApiKey)
}

func (controller UserController) GetApiKey(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)

	apiKey, err := controller.userService.GetApiKey(uid)
	if err != nil {
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
	uid := ctx.GetString("uid")

	if err := controller.userService.DeleteApiKey(uid); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.ApiKeyDeleted, nil)
}

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

	apiKey.GET("", controller.GetAPIKey)
	apiKey.POST("", controller.GenerateAPIKey)
	apiKey.DELETE("", controller.RevokeAPIKey)
	apiKey.GET("/usage/today", controller.GetTodayUsage)
	apiKey.GET("/usage/total", controller.GetTotalUsage)
}

// @Summary Get API Key
// @Description Retrieves the API key for the authenticated user
// @Tags API Keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.APIKeyResponse "API key retrieved successfully"
// @Failure 404 {object} response.Response "No API key found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/api-key [get]
func (controller UserController) GetAPIKey(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)

	apiKey, err := controller.userService.GetApiKey(uid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			response.WithSuccess(ctx, http.StatusOK, message.ApiKeyRetrieved, models.APIKeyResponse{
				APIKey: nil,
			})
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

// @Summary Create API Key
// @Description Creates a new API key for the authenticated user
// @Tags API Keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string true "Name for the API key"
// @Success 201 {object} models.APIKeyResponse "API key created successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/api-key [post]
func (controller UserController) GenerateAPIKey(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)

	apiKey, err := controller.userService.CreateApiKey(uid)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
		return
	}

	response.WithSuccess(ctx, http.StatusCreated, message.ApiKeyCreated, models.APIKeyResponse{
		APIKey: apiKey,
	})
}

// @Summary Delete API Key
// @Description Deletes the API key for the authenticated user
// @Tags API Keys
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response "API key deleted successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/api-key [delete]
func (controller UserController) RevokeAPIKey(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)

	if err := controller.userService.DeleteApiKey(uid); err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.ApiKeyDeleted, nil)
}

// @Summary Get Today Usage
// @Description Retrieves the today usage for the authenticated user
// @Tags Usage
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} gin.H "Today usage retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/api-key/usage/today [get]
func (controller UserController) GetTodayUsage(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)

	usage, err := controller.userService.GetTodayUsage(uid)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.UsageRetrieved, gin.H{
		"usage": usage,
	})
}

// @Summary Get Total Usage
// @Description Retrieves the total usage for the authenticated user
// @Tags Usage
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} gin.H "Total usage retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/api-key/usage/total [get]
func (controller UserController) GetTotalUsage(ctx *gin.Context) {
	uid := ctx.GetString(api.UidParam)

	usage, err := controller.userService.GetTotalUsage(uid)
	if err != nil {
		log.Println(err.Error())
		response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, message.UsageRetrieved, gin.H{
		"usage": usage,
	})
}

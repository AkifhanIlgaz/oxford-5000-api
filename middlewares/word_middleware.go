package middlewares

import (
	"net/http"

	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

type WordMiddleware struct {
	wordService services.WordService
	userService services.UserService
	authService services.AuthService
}

func NewWordMiddleware(wordService services.WordService, userService services.UserService, authService services.AuthService) WordMiddleware {
	return WordMiddleware{
		wordService: wordService,
		userService: userService,
		authService: authService,
	}
}

func (m *WordMiddleware) TrackUsage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Query(api.ApiKeyParam)
		if key == "" {
			response.WithError(ctx, http.StatusBadRequest, message.ApiKeyRequired)
			return
		}

		apiKeyDoc, err := m.userService.GetByKey(key)
		if err != nil {
			response.WithError(ctx, http.StatusUnauthorized, message.InvalidApiKey)
			return
		}

		dailyUsage, err := m.userService.IncrementUsage(apiKeyDoc.Key)
		if err != nil {
			response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
			return
		}

		plan, err := m.authService.GetUserPlan(apiKeyDoc.Uid)
		if err != nil {
			response.WithError(ctx, http.StatusInternalServerError, message.ApiKeyError)
			return
		}

		usageLimit := api.PlanUsageLimits[plan]
		if dailyUsage > usageLimit {
			response.WithError(ctx, http.StatusPaymentRequired, message.UsageLimitReached)
			return
		}

		ctx.Next()
	}
}

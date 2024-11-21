package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AkifhanIlgaz/dictionary-api/services"
	"github.com/AkifhanIlgaz/dictionary-api/utils/api"
	"github.com/AkifhanIlgaz/dictionary-api/utils/message"
	"github.com/AkifhanIlgaz/dictionary-api/utils/response"
	"github.com/gin-gonic/gin"
)

// UserMiddleware handles user authentication middleware operations
// using a token service for validation.
type UserMiddleware struct {
	tokenService services.TokenService
}

// NewUserMiddleware creates a new UserMiddleware instance with the provided token service.
// Parameters:
//   - tokenService: Service handling token operations
//
// Returns:
//   - UserMiddleware: New middleware instance
func NewUserMiddleware(tokenService services.TokenService) UserMiddleware {
	return UserMiddleware{
		tokenService: tokenService,
	}
}

// parseIdTokenFromHeader extracts and validates the Bearer token from the HTTP header.
// The token should be in the format: "Bearer <token>"
//
// Parameters:
//   - header: HTTP header containing the authorization token
//
// Returns:
//   - string: The extracted token
//   - error: Error if the token format is invalid or missing
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

// AuthenticateUser returns a Gin middleware handler that authenticates users
// by validating their access token. If authentication is successful, the user's
// UID is added to the Gin context.
//
// Returns:
//   - gin.HandlerFunc: Middleware handler for authentication
//
// Context Sets:
//   - api.ContextUid: User ID extracted from the token
func (m UserMiddleware) AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token from header
		token, err := parseIdTokenFromHeader(ctx.Request.Header)
		if err != nil {
			response.WithError(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		// Parse and validate token
		uid, err := m.tokenService.ParseToken("access", token)
		if err != nil {
			response.WithError(ctx, http.StatusUnauthorized, message.InvalidToken)
			return
		}

		ctx.Set(api.ContextUid, uid)
		ctx.Next()
	}
}

package response

import "github.com/gin-gonic/gin"

const (
	statusSuccess string = "success"
	statusError   string = "error"
)

// Response represents the standard API response structure
// @Description Standard API response format
type Response struct {
	// Success indicates if the request was successful
	// @Description Indicates if the request was successful
	Success bool        `json:"success"`
	
	// Message contains a human-readable response message
	// @Description Human-readable response message
	Message string      `json:"message"`
	
	// Data contains the actual response payload
	// @Description Response payload data
	Data    interface{} `json:"data"`
	
	// Error contains error details if Success is false
	// @Description Error details when Success is false
	Error   string      `json:"error,omitempty"`
}

// WithSuccess sends a JSON response with a success status.
// Parameters:
//   - ctx: Gin context for the HTTP response
//   - statusCode: HTTP status code to be sent
//   - message: Human-readable message describing the success
//   - data: Any data to be included in the response
func WithSuccess(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(statusCode, gin.H{
		"status":  statusSuccess,
		"message": message,
		"result":  data,
	})
}

// WithError sends a JSON error response and aborts the request chain.
// Parameters:
//   - ctx: Gin context for the HTTP response
//   - statusCode: HTTP status code to be sent
//   - message: Human-readable error message
func WithError(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status":  statusError,
		"message": message,
	})
}

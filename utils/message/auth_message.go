package message

const (
	MissingField  string = "missing field"
	WrongPassword string = "wrong password"

	// API Key related messages
	ApiKeyRetrieved string = "api key retrieved"
	ApiKeyCreated   string = "api key created"
	ApiKeyDeleted   string = "api key deleted"
	ApiKeyError     string = "error processing api key"
	ApiKeyRequired  string = "api key is required"

	InvalidApiKey string = "invalid api key"
	InvalidToken  string = "invalid or expired token"
)

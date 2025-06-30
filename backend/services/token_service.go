package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
)

// TokenService handles JWT token generation and parsing operations
type TokenService struct {
	client *redis.Client
	config config.Config
}

// NewTokenService creates a new TokenService instance with the provided configuration
func NewTokenService(config config.Config, client *redis.Client) TokenService {
	return TokenService{
		config: config,
		client: client,
	}
}

// CreateTokens generates both access and refresh tokens for a given user ID
// Returns a models.Tokens struct containing both tokens, or an error if token generation fails
func (service TokenService) CreateTokens(uid string) (models.Tokens, error) {
	accessToken, err := service.generateToken("access", uid)
	if err != nil {
		return models.Tokens{}, fmt.Errorf("create tokens: %w", err)
	}

	refreshToken, err := service.generateToken("refresh", uid)
	if err != nil {
		return models.Tokens{}, fmt.Errorf("create tokens: %w", err)
	}

	err = service.client.Set(service.client.Context(), "refresh_tokens:"+refreshToken, uid, time.Duration(service.config.RefreshTokenExpiry)*time.Hour).Err()
	if err != nil {
		return models.Tokens{}, fmt.Errorf("store refresh token on redis: %w", err)
	}

	return models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (service TokenService) IsRefreshTokenExpired(token string) (bool, error) {
	// Check if the token exists in Redis
	_, err := service.client.Get(service.client.Context(), "refresh_tokens:"+token).Result()
	if err != nil {
		if err == redis.Nil {
			return false, errors.New("refresh token not found")
		}
		return false, fmt.Errorf("check refresh token in redis: %w", err)
	}

	// If found, return the user ID
	return true, nil
}

// generateToken creates a specific type of JWT token (access or refresh) for a given user ID
// Parameters:
//   - tokenType: "access" or "refresh"
//   - uid: user identifier
//
// Returns the generated token as a string or an error if generation fails
func (service TokenService) generateToken(tokenType string, uid string) (string, error) {
	switch tokenType {
	case "access":
		return generateToken(service.config.AccessTokenPrivateKey, uid, service.config.AccessTokenExpiry)
	case "refresh":
		return generateToken(service.config.RefreshTokenPrivateKey, uid, service.config.RefreshTokenExpiry)
	default:
		return "", errors.New("unsupported token type")
	}
}

// ParseToken validates and extracts the user ID from a given token
// Parameters:
//   - tokenType: "access" or "refresh"
//   - token: the JWT token string to parse
//
// Returns the user ID (uid) from the token claims or an error if parsing fails
func (service TokenService) ParseToken(tokenType string, token string) (string, error) {
	switch tokenType {
	case "access":
		return parseToken(service.config.AccessTokenPublicKey, token)
	case "refresh":
		return parseToken(service.config.RefreshTokenPublicKey, token)
	default:
		return "", errors.New("unsupported token type")
	}
}

// generateToken creates a JWT token using RSA private key encryption
// Parameters:
//   - privKey: base64 encoded RSA private key
//   - uid: user identifier to be included in token claims
//   - expiryHour: token expiration time in hours
//
// Returns the signed JWT token as a string or an error if generation fails
func generateToken(privKey string, uid string, expiryHour int) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	claims := jwt.RegisteredClaims{
		Subject:   uid,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiryHour) * time.Hour)),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return token, nil
}

// parseToken validates and decodes a JWT token using RSA public key
// Parameters:
//   - publicKey: base64 encoded RSA public key
//   - token: JWT token string to parse
//
// Returns the user ID from the token claims or an error if validation fails
func parseToken(publicKey string, token string) (string, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", fmt.Errorf("extract uid from access token: %w", err)
	}

	parsedToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return jwt.RegisteredClaims{}, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	})

	if err != nil {
		return "", fmt.Errorf("parse token: %w", err)
	}

	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)
	if !ok || !parsedToken.Valid {
		return "", fmt.Errorf("validate: invalid token")
	}

	return claims.Subject, nil
}

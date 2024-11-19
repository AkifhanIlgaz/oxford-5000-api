package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/AkifhanIlgaz/dictionary-api/models"
	"github.com/golang-jwt/jwt/v4"
)

// TokenService handles JWT token generation and parsing operations
type TokenService struct {
	config config.Config
}

// NewTokenService creates a new TokenService instance with the provided configuration
func NewTokenService(config config.Config) TokenService {
	return TokenService{
		config: config,
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

	return models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
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

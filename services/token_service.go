package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/dictionary-api/config"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
	config config.Config
}

func NewTokenService(config config.Config) TokenService {
	return TokenService{
		config: config,
	}
}

func (service TokenService) GenerateToken(tokenType string, uid string) (string, error) {
	switch tokenType {
	case "access":
		return generateToken(service.config.AccessTokenPrivateKey, uid, service.config.AccessTokenExpiry)
	case "refresh":
		return generateToken(service.config.RefreshTokenPrivateKey, uid, service.config.RefreshTokenExpiry)
	default:
		return "", errors.New("unsupported token type")
	}
}

// returns uid
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

package services

import "github.com/AkifhanIlgaz/dictionary-api/config"

type TokenService struct {
	config config.Config
}

func NewTokenService(config config.Config) TokenService {
	return TokenService{
		config: config,
	}
}

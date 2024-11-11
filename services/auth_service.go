package services

import (
	"context"
)

type AuthService struct {
	ctx context.Context
}

func NewAuthService(ctx context.Context) AuthService {
	return AuthService{
		ctx: ctx,
	}
}

package models

import (
	"fmt"
	"net/mail"
)

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (req AuthRequest) Validate() error {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return fmt.Errorf("validate signup request: %w", err)
	}

	return nil
}

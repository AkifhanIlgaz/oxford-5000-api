package models

type AuthRequest struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required,min=6"`
}

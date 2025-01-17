package model

import "time"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshToken struct {
	Refresh    string
	UserID     string
	Expires_at time.Time
	Revoked    bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RequestResetPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type RequestResetPasswordResponse struct {
	Message string `json:"message"`
}

type ResetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required"`
}

type VerifyEmailResponse struct {
	Message string `json:"message"`
}

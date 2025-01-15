package model

import "time"

type User struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	Password        string `json:"-"`
	IsEmailVerified bool   `json:"-"`
	WalletAddress   string `json:"walletAddress"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CreateUserReqesut struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

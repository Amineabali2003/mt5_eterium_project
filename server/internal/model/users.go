package model

import "time"

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CreateUserReqesut struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

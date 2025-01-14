package services

import (
	"fmt"
	"os"

	"github.com/idir-44/ethereum/internal/jwttoken"
	"github.com/idir-44/ethereum/internal/model"
	"github.com/idir-44/ethereum/pkg/utils"
)

func (s service) Login(req model.LoginRequest) (model.User, string, error) {
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return model.User{}, "", err
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return model.User{}, "", fmt.Errorf("invalid password: %s", err)
	}

	key := os.Getenv("jwt_secret")
	if key == "" {
		return model.User{}, "", fmt.Errorf("jwt secret is not set")
	}

	token, err := jwttoken.CreateToken(user, key)
	if err != nil {
		return model.User{}, "", err
	}

	return user, token, nil

}

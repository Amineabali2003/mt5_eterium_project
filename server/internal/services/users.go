package services

import (
	"fmt"
	"os"

	"github.com/idir-44/ethereum/internal/jwttoken"
	"github.com/idir-44/ethereum/internal/model"
	"github.com/idir-44/ethereum/pkg/utils"
)

func (s service) CreateUser(req model.CreateUserReqesut) (model.User, error) {

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return model.User{}, err
	}

	user, err := s.repository.CreateUser(model.User{Email: req.Email, Password: password})

	key := os.Getenv("jwt_secret")
	if key == "" {
		return model.User{}, fmt.Errorf("jwt secret is not set")
	}

	token, err := jwttoken.CreateToken(user, key, jwttoken.TokenTypeEmailValidation)
	if err != nil {
		return model.User{}, fmt.Errorf("cannot create verification token")
	}

	// TODO: move front url to env
	// send email with link
	fmt.Printf("\nhttp://localhost:3000/verify-email/%s\n", token)

	return user, err
}

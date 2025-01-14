package services

import (
	"github.com/idir-44/ethereum/internal/model"
	"github.com/idir-44/ethereum/pkg/utils"
)

func (s service) CreateUser(req model.CreateUserReqesut) (model.User, error) {

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return model.User{}, err
	}

	return s.repository.CreateUser(model.User{Email: req.Email, Password: password})
}

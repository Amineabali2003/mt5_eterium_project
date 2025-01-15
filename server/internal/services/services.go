package services

import (
	"github.com/idir-44/ethereum/internal/model"
	"github.com/idir-44/ethereum/internal/repositories"
)

type service struct {
	repository repositories.Repository
}

func NewService(repo repositories.Repository) Service {
	return service{repo}
}

type Service interface {
	CreateUser(req model.CreateUserReqesut) (model.User, error)
	Login(req model.LoginRequest) (model.User, string, error)
	RequestResetPassword(email string) error
	ResetPassword(req model.ResetPasswordRequest) error
	VerifyEmail(req model.VerifyEmailRequest) error
	UpdateWallet(userID, walletAddress string) (model.User, error)
}

func (s service) UpdateWallet(userID, walletAddress string) (model.User, error) {
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return model.User{}, err
	}
	user.WalletAddress = walletAddress
	return s.repository.UpdateUser(userID, user)
}

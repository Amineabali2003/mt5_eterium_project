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

	if !user.IsEmailVerified {
		return model.User{}, "", fmt.Errorf("Verify email before login")
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return model.User{}, "", fmt.Errorf("invalid password: %s", err)
	}

	key := os.Getenv("jwt_secret")
	if key == "" {
		return model.User{}, "", fmt.Errorf("jwt secret is not set")
	}

	token, err := jwttoken.CreateToken(user, key, jwttoken.TokenTypeAccess)
	if err != nil {
		return model.User{}, "", err
	}

	return user, token, nil

}

func (s service) RequestResetPassword(email string) error {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	key := os.Getenv("jwt_secret")
	if key == "" {
		return fmt.Errorf("jwt secret is not set")
	}

	token, err := jwttoken.CreateToken(user, key, jwttoken.TokenTypeResetPassword)
	if err != nil {
		return err
	}

	// TODO: move front url to env
	// send email with link
	fmt.Printf("\nhttp://localhost:3000/reset-password/%s\n", token)

	return nil
}

func (s service) ResetPassword(req model.ResetPasswordRequest) error {

	key := os.Getenv("jwt_secret")
	if key == "" {
		return fmt.Errorf("jwt secret is not set")
	}

	userClaims, err := jwttoken.ParseToken(req.Token, key)
	if err != nil {
		return err
	}

	pwd, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil
	}

	user, err := s.repository.GetUser(userClaims.ID)
	if err != nil {
		return err
	}

	if user.Email != userClaims.Email {
		return fmt.Errorf("whatchudoing ?")
	}

	user.Password = pwd

	_, err = s.repository.UpdateUser(userClaims.ID, user)
	if err != nil {
		return err
	}

	return nil

}

func (s service) VerifyEmail(req model.VerifyEmailRequest) error {

	key := os.Getenv("jwt_secret")
	if key == "" {
		return fmt.Errorf("jwt secret is not set")
	}

	userClaims, err := jwttoken.ParseToken(req.Token, key)
	if err != nil {
		return err
	}

	user, err := s.repository.GetUserByEmail(userClaims.Email)
	if err != nil {
		return err
	}

	if user.IsEmailVerified {
		return fmt.Errorf("Email already verified")
	}

	user.IsEmailVerified = true

	_, err = s.repository.UpdateUser(user.ID, user)

	return err
}

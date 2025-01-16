package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/idir-44/ethereum/internal/jwttoken"
	"github.com/idir-44/ethereum/internal/model"
	"github.com/idir-44/ethereum/pkg/utils"
)

func (s service) SendEmail(to []string, content string, subject string) error {

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	if from == "" || password == "" {
		return fmt.Errorf("email or password not set")
	}
	host := "smtp.gmail.com"
	port := "587"
	contentType := "Content-Type: text/html; charset=\"utf-8\"\n\n"

	body := []byte(subject + contentType + content)

	auth := smtp.PlainAuth("", from, password, host)

	return smtp.SendMail(host+":"+port, auth, from, to, body)
}

const htmlContent = `
<!DOCTYPE html>
<html>
<head>
    <style>
        .btn {
            display: inline-block;
            background-color: #4CAF50;
            color: white;
            padding: 10px 20px;
            text-decoration: none;
            border-radius: 5px;
        }
    </style>
</head>
<body>
    <h1>Verify Your Email</h1>
    <p>Click the button below to verify your email:</p>
    <a href="http://localhost:3000/verify-email/%s" class="btn">Verify Email</a>
</body>
</html>
`

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
	fmt.Printf("\nhttp://localhost:3000/verify-email/%s\n", token)

	emailErr := s.SendEmail([]string{user.Email}, fmt.Sprintf(htmlContent, token), "Subject: Verify Your Email\n")
	if emailErr != nil {
		fmt.Println(emailErr)
	}

	return user, err
}

func (s service) UpdateWallet(userID, walletAddress string) (model.User, error) {
	user, err := s.repository.GetUser(userID)
	if err != nil {
		return model.User{}, err
	}
	user.WalletAddress = walletAddress
	return s.repository.UpdateUser(userID, user)
}

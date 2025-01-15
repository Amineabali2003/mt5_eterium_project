package controllers

import (
	"net/http"
	"time"

	"github.com/idir-44/ethereum/internal/model"
	"github.com/labstack/echo/v4"
)

func (r controller) login(c echo.Context) error {
	req := model.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, token, err := r.service.Login(req)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, user)
}

func (r controller) requestResetPassword(c echo.Context) error {
	req := model.RequestResetPasswordRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	err := r.service.RequestResetPassword(req.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.RequestResetPasswordResponse{Message: "A password reset link has been sent to your email address. Please check your inbox to proceed with resetting your password."})

}

func (r controller) resetPassword(c echo.Context) error {
	req := model.ResetPasswordRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := r.service.ResetPassword(req); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "")

}

func (r controller) verifyEmail(c echo.Context) error {
	req := model.VerifyEmailRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	err := r.service.VerifyEmail(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.VerifyEmailResponse{Message: "Your email has been verified"})

}

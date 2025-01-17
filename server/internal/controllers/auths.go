package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/idir-44/ethereum/internal/jwttoken"
	"github.com/idir-44/ethereum/internal/model"
	"github.com/labstack/echo/v4"
)

func (r controller) login(c echo.Context) error {
	req := model.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, token, refreshToken, err := r.service.Login(req)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	refreshCookie := new(http.Cookie)
	refreshCookie.Name = "refreshToken"
	refreshCookie.Value = refreshToken
	refreshCookie.Expires = time.Now().Add(72 * time.Hour)
	refreshCookie.HttpOnly = true
	c.SetCookie(refreshCookie)

	return c.JSON(http.StatusOK, user)
}

func (r controller) verifyRefreshToken(c echo.Context) error {
	cookies, err := c.Cookie("refreshToken")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	key := os.Getenv("jwt_secret")
	if key == "" {
		return fmt.Errorf("jwt secret is not set")
	}
	user, err := jwttoken.ParseToken(cookies.Value, key)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	accessToken, refreshToken, err := r.service.VerifyRefreshToken(user, cookies.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = accessToken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	refreshCookie := new(http.Cookie)
	refreshCookie.Name = "refreshToken"
	refreshCookie.Value = refreshToken
	refreshCookie.Expires = time.Now().Add(72 * time.Hour)
	refreshCookie.HttpOnly = true
	c.SetCookie(refreshCookie)

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

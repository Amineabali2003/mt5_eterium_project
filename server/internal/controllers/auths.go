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

package controllers

import (
	"net/http"

	"github.com/idir-44/ethereum/internal/model"
	"github.com/labstack/echo/v4"
)

func (r controller) createUser(c echo.Context) error {
	req := model.CreateUserReqesut{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := r.service.CreateUser(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)

}

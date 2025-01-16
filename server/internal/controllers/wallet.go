package controllers

import (
	"net/http"

	"github.com/idir-44/ethereum/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func (r controller) getWalletData(c echo.Context) error {
	user, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}

	res, err := r.service.GetWalletData(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)

}

func (r controller) getWalletTransactions(c echo.Context) error {
	user, err := middlewares.GetUser(c)
	if err != nil {
		return err
	}

	res, err := r.service.GetWalletTransactions(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)

}

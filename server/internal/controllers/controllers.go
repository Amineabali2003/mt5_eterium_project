package controllers

import (
	"github.com/idir-44/ethereum/internal/middlewares"
	"github.com/idir-44/ethereum/internal/services"
	"github.com/idir-44/ethereum/pkg/server"
)

type controller struct {
	service services.Service
}

func RegisterHandlers(routerGroup *server.Router, srv services.Service) {
	c := controller{srv}

	routerGroup.Use(middlewares.AddCurentUser)

	routerGroup.POST("/users", c.createUser)
	routerGroup.POST("/login", c.login)

	routerGroup.GET("/me", c.getCurrentUser)

}

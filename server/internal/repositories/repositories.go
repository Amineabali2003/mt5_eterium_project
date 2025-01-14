package repositories

import (
	"github.com/idir-44/ethereum/internal/model"
	"github.com/idir-44/ethereum/pkg/database"
)

type repository struct {
	db *database.DBConnection
}

func NewRepository(db *database.DBConnection) Repository {
	return repository{db}
}

type Repository interface {
	CreateUser(user model.User) (model.User, error)
}

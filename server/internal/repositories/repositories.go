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
	GetUserByEmail(email string) (model.User, error)
	GetUser(id string) (model.User, error)
	UpdateUser(id string, user model.User) (model.User, error)
	CreateRefreshToken(user model.User, refresh string) (model.User, error)
	GetRefreshToken(refresh string) (model.RefreshToken, error)
}

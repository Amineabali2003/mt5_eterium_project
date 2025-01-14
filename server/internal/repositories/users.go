package repositories

import (
	"context"
	"time"

	"github.com/idir-44/ethereum/internal/model"
)

func (r repository) CreateUser(user model.User) (model.User, error) {

	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	_, err := r.db.NewInsert().Model(&user).ExcludeColumn("id").Returning("*").Exec(context.TODO())
	if err != nil {
		return model.User{}, err
	}

	return user, nil

}

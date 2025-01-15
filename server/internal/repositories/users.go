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

func (r repository) GetUserByEmail(email string) (model.User, error) {
	user := model.User{}
	err := r.db.NewSelect().Model(&user).Where("email = ?", email).Scan(context.TODO())
	return user, err
}

func (r repository) GetUser(id string) (model.User, error) {
	user := model.User{}
	err := r.db.NewSelect().Model(&user).Where("id = ?", id).Scan(context.TODO())
	return user, err
}

func (r repository) UpdateUser(id string, user model.User) (model.User, error) {
	updateUser := map[string]interface{}{
		"updated_at": time.Now().UTC(),
	}
	query := r.db.NewUpdate().Model(&updateUser).TableExpr("users")
	if user.Password != "" {
		updateUser["password"] = user.Password
	}
	if user.IsEmailVerified {
		updateUser["is_email_verified"] = user.IsEmailVerified
	}
	if user.WalletAddress != "" {
		updateUser["wallet_address"] = user.WalletAddress
	}
	_, err := query.Where("id = ?", id).Returning("*").Exec(context.TODO())
	return user, err
}

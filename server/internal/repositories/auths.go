package repositories

import (
	"context"
	"time"

	"github.com/idir-44/ethereum/internal/model"
)

func (r repository) CreateRefreshToken(user model.User, refresh string) (model.User, error) {
	refreshToken := model.RefreshToken{}

	refreshToken.CreatedAt = time.Now().UTC()
	refreshToken.UpdatedAt = time.Now().UTC()
	refreshToken.UserID = user.ID
	refreshToken.Refresh = refresh
	refreshToken.Expires_at = time.Now().Add(72 * time.Hour).UTC()

	_, err := r.db.NewInsert().Model(&refreshToken).ExcludeColumn("id").Returning("*").Exec(context.TODO())
	if err != nil {
		return model.User{}, err
	}
	return user, nil

}

func (r repository) GetRefreshToken(refresh string) (model.RefreshToken, error) {
	refreshToken := model.RefreshToken{}

	err := r.db.NewSelect().Model(&refreshToken).Where("refresh = ?", refresh).Scan(context.TODO())

	return refreshToken, err
}

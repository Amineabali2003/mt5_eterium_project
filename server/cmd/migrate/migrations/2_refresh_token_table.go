package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

func init() {

	up := []string{`
		CREATE TABLE IF NOT EXISTS refresh_tokens (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id VARCHAR(255) NOT NULL UNIQUE,
		refresh TEXT NOT NULL UNIQUE,
		expires_at TIMESTAMP NOT NULL,
		revoked bool DEFAULT false,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

		FOREIGN KEY (user_id) REFERENCES users(id)

	);
	`}

	down := []string{`
			DROP TABLE IF EXISTS refresh_token CASCADE;
	`}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print("create refresh_tokens table")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print("drop refresh_tokens table")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

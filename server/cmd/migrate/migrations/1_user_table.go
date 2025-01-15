package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

func init() {

	up := []string{`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		is_email_verified BOOLEAN DEFAULT false,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	`}

	down := []string{`
			DROP TABLE IF EXISTS users CASCADE;
	`}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print("create users table")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print("drop users table")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

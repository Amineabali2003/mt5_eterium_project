package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DBConnection = bun.DB

func checkConnection(db *bun.DB) error {
	_, err := db.ExecContext(context.TODO(), "SELECT 1")
	return err
}

func Connect() (*DBConnection, error) {

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithUser(os.Getenv("PGUSER")),
		pgdriver.WithAddr(fmt.Sprintf("%s:%s", os.Getenv("PGHOST"), os.Getenv("PGPORT"))),
		pgdriver.WithDatabase(os.Getenv("PGDATABASE")),
		pgdriver.WithPassword(os.Getenv("PGPASSWORD")),
		pgdriver.WithInsecure(true),
	))

	db := bun.NewDB(sqldb, pgdialect.New())

	if err := checkConnection(db); err != nil {
		return &DBConnection{}, err
	}

	return db, nil
}

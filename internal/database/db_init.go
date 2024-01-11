package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func DB_Init(dbUrl string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

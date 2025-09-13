package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func Init() error {
	connStr := "postgres://dustinmeyer@localhost:5432/timewarp"

	var err error
	pool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatalf("unable to create Postgres pool: %v", err)
	}

	pool.Config().MaxConns = 25
	pool.Config().MinConns = 5

	err = pool.Ping(context.Background())

	if err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	log.Println("Postgres connection pool initialized")

	return nil

}

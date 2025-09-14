package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	AppEnv           string
	ConnectionString string
}

var pool *pgxpool.Pool

func LoadDatabaseConfig(appEnv string) (*Config, error) {

	var connectionString string

	switch appEnv {
	case "prod":
		connectionString = os.Getenv("PROD_DATABASE_URL")
	case "tst":
		connectionString = os.Getenv("TEST_DATABASE_URL")
	case "dev":
		connectionString = os.Getenv("DATABASE_URL")
	default:
		return nil, fmt.Errorf("unknow APP_ENV: %s", appEnv)
	}

	if connectionString == "" {
		return nil, fmt.Errorf("database URL not set for enviroment: %s", appEnv)
	}

	return &Config{ConnectionString: connectionString, AppEnv: appEnv}, nil
}

func (c *Config) Init() error {

	var err error
	pool, err = pgxpool.New(context.Background(), c.ConnectionString)

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

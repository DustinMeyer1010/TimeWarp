package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	AppEnv           string
	ConnectionString string
}

var pool *pgxpool.Pool

// Loads the correct database based on the app enviroment that you are in tst, dev, prod
func LoadDatabaseConfig(appEnv string) (*Config, error) {

	var connectionString string

	switch appEnv {
	case "prod":
		connectionString = os.Getenv("PROD_DATABASE_URL")
	case "tst":
		connectionString = os.Getenv("TEST_DATABASE_URL")
	case "dev":
		connectionString = os.Getenv("DEV_DATABASE_URL")
	default:
		return nil, fmt.Errorf("unknow APP_ENV: %s", appEnv)
	}

	if connectionString == "" {
		return nil, fmt.Errorf("database URL not set for enviroment: %s", appEnv)
	}

	return &Config{ConnectionString: connectionString, AppEnv: appEnv}, nil
}

// Creates a connection pool and looks for all Up mirgrations to run to make sure database is fully created
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

	createAllTables()

	log.Println("Postgres connection pool initialized")

	return nil

}

// Runs through all Up migrations and creates the tables and other queries to make sure database has everything to run properly
func createAllTables() {
	dir := "./migrations"
	files, err := os.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		content, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			panic(err)
		}
		query := string(content)
		createTable(query)
	}
}

// Runs a query that is given to create tables
func createTable(query string) {
	_, err := pool.Exec(
		context.Background(),
		string(query),
	)

	if err != nil {
		panic(err)
	}
}

// Given a table name it will clear the specifc tables (usually only used for testing)
func ClearTables(tableNames ...string) {
	for _, name := range tableNames {
		query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", name)
		pool.Exec(
			context.Background(),
			query,
		)
	}
}

func ClearAllTables() {
	rows, err := pool.Query(
		context.Background(),
		`
		SELECT tablename
		FROM pg_catalog.pg_tables
		WHERE schemaname = 'public';
		`,
	)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var table string
		err := rows.Scan(&table)

		if err != nil {
			panic(err)
		}

		ClearTables(table)
	}

}

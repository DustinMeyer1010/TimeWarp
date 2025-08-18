package db

import (
	"context"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func GetAccountByUsername(username string) (*types.Account, error) {

	var account types.Account

	row := pool.QueryRow(
		context.Background(),
		"SELECT id, username, password FROM account WHERE username = $1",
		username,
	)

	err := row.Scan(&account.ID, &account.Username, &account.Password)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func GetAccountByID(id int) (*types.Account, error) {
	var account types.Account

	row := pool.QueryRow(
		context.Background(),
		"SELECT id, username FROM account WHERE id = $1",
		id,
	)

	err := row.Scan(&account.ID, &account.Username)

	if err != nil {
		return nil, err
	}

	return &account, nil
}

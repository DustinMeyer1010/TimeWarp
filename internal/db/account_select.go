package db

import (
	"context"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func GetAccountByUsername(username string) (*types.Account, error) {

	var account types.Account

	row := pool.QueryRow(
		context.Background(),
		"SELECT username, password FROM account WHERE username = $1",
		username,
	)

	err := row.Scan(&account.Username, &account.Password)

	if err != nil {
		return nil, nil
	}

	return &account, nil
}

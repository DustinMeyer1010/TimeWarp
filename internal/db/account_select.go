package db

import (
	"context"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

// Returns account ID, Username, Password for account matching username
// If account not found returns empty account with error
func GetAccountByUsername(username string) (types.Account, error) {

	var account types.Account

	row := pool.QueryRow(
		context.Background(),
		"SELECT id, username, password FROM account WHERE username = $1",
		username,
	)

	err := row.Scan(&account.ID, &account.Username, &account.Password)

	if err != nil {
		return types.Account{}, err
	}

	return account, nil
}

// Pulls account with the id given
func GetAccountByID(id int) (types.Account, error) {
	var account types.Account

	err := pool.QueryRow(
		context.Background(),
		"SELECT id, username, email, password, creation_date FROM account WHERE id = $1",
		id,
	).Scan(&account.ID, &account.Username, &account.Email, &account.Password, &account.CreationDate)

	return account, err
}

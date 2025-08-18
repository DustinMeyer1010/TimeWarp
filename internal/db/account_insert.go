package db

import (
	"context"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateAccount(account types.Account) error {

	err := account.EncryptPassword()

	if err != nil {
		return err
	}

	_, err = pool.Exec(
		context.Background(),
		"INSERT INTO account (username, password) VALUES ($1, $2)",
		account.Username, account.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

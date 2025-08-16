package db

import (
	"context"
	"fmt"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateAccount(account types.Account) error {

	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO account (username, password) VALUES ($1, $2)",
		account.Username, account.Password,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

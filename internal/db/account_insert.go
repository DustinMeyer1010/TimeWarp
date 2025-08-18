package db

import (
	"context"
	"fmt"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateAccount(account types.Account) error {

	err := account.EncryptPassword()

	if err != nil {
		return err
	}

	_, err = pool.Exec(
		context.Background(),
		"INSERT INTO account (username, password, email, creation_date) VALUES ($1, $2, $3, $4)",
		account.Username, account.Password, account.Email, time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func AddRefreshToken(id int, refreshToken string) error {

	_, err := pool.Exec(
		context.Background(),
		"UPDATE account SET refresh_token = $1 WHERE id = $2",
		refreshToken, id,
	)

	fmt.Println(id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

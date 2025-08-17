package db

import "context"

func DeleteAccount(username string) error {

	_, err := pool.Exec(
		context.Background(),
		"DELETE FROM account WHERE username=$1",
		username,
	)

	if err != nil {
		return err
	}

	return nil
}

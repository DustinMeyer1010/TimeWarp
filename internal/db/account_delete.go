package db

import "context"

func DeleteAccountByUsername(username string) (int, error) {
	var accountID int = -1

	err := pool.QueryRow(
		context.Background(),
		"DELETE FROM account WHERE username=$1 RETURNING id",
		username,
	).Scan(&accountID)

	return accountID, err
}

func DeleteAccountById(id int) (int, error) {
	var accountID int = -1
	err := pool.QueryRow(
		context.Background(),
		"DELETE FROM account WHERE id=$1 RETURNING id",
		id,
	).Scan(&accountID)

	return accountID, err
}

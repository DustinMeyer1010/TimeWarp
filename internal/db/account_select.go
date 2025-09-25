package db

import (
	"context"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
)

// GetAccountByUsername retrieves an account from the database that matches the provided username.
//
// It performs a SQL SELECT query on the `account` table to fetch the `id`, `username`, and `password`
// fields for the account with the specified username. The result is scanned into a `models.Account`
// struct and returned.
//
// Parameters:
//   - username: the username of the account to retrieve.
//
// Returns:
//   - models.Account: a populated account struct if found; otherwise, an empty account struct.
//   - error: an error object if the query fails or no account is found.
//
// Example:
//
//	account, err := GetAccountByUsername("johndoe")
//	if err != nil {
//	    log.Printf("Account not found or error occurred: %v", err)
//	} else {
//	    log.Printf("Account ID: %d, Username: %s", account.ID, account.Username)
//	}
func GetAccountByUsername(username string) (models.Account, error) {

	var account models.Account

	row := pool.QueryRow(
		context.Background(),
		"SELECT id, username, password FROM account WHERE username = $1",
		username,
	)

	err := row.Scan(&account.ID, &account.Username, &account.Password)

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

// GetAccountByID retrieves an account from the database using the provided account ID.
//
// It performs a SQL SELECT query on the `account` table to fetch the fields: `id`, `username`,
// `email`, `password`, and `creation_date` for the account with the specified ID. The result is
// scanned into a `models.Account` struct and returned.
//
// Parameters:
//   - id: the unique identifier of the account to retrieve.
//
// Returns:
//   - models.Account: a populated account struct containing the account details.
//   - error: an error object if the query fails or no account is found.
//
// Example:
//
//	account, err := GetAccountByID(123)
//	if err != nil {
//	    log.Printf("Failed to retrieve account: %v", err)
//	} else {
//	    log.Printf("Account retrieved: %+v", account)
//	}
func GetAccountByID(id int) (models.Account, error) {
	var account models.Account

	err := pool.QueryRow(
		context.Background(),
		"SELECT id, username, email, password, creation_date FROM account WHERE id = $1",
		id,
	).Scan(&account.ID, &account.Username, &account.Email, &account.Password, &account.CreationDate)

	return account, err
}

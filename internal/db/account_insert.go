package db

import (
	"context"
	"fmt"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
)

// CreateAccount attempts to create a new account in the database.
//
// It returns the ID of the newly created account on success.
// If the account could not be created, it returns -1 as the ID along with a non-nil error.
//
// Parameters:
//   - account: The account information to be created.
//
// Returns:
//   - id: The ID of the newly created account, or -1 if the creation failed.
//   - err: An error describing why the account could not be created, or nil if successful.
func CreateAccount(account models.Account) (id int, err error) {
	id = -1
	err = account.EncryptPassword()

	if err != nil {
		return -1, err
	}

	err = pool.QueryRow(
		context.Background(),
		`INSERT INTO account (username, password, email, creation_date) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`,
		account.Username, account.Password, account.Email, time.Now(),
	).Scan(&id)

	return
}

// AddRefreshToken updates the refresh token for a specific account in the database.
//
// It executes an SQL UPDATE statement on the `account` table, setting the `refresh_token`
// field to the provided value for the account with the given ID.
//
// Parameters:
//   - id: the unique identifier of the account to update.
//   - refreshToken: the new refresh token string to be stored.
//
// Returns:
//   - error: returns an error if the update operation fails; otherwise, returns nil.
//
// Example:
//
//	err := AddRefreshToken(123, "new-refresh-token")
//	if err != nil {
//	    log.Printf("Failed to update refresh token: %v", err)
//	} else {
//	    log.Println("Refresh token updated successfully")
//	}
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

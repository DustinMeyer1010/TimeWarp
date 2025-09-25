package db

import "context"

// DeleteAccountByUsername deletes an account from the database using the provided username.
//
// It executes a SQL DELETE statement on the `account` table where the `username` matches
// the input parameter. The query uses the RETURNING clause to retrieve the `id` of the
// deleted account.
//
// Parameters:
//   - username: the username of the account to be deleted.
//
// Returns:
//   - int: the ID of the deleted account. If no account is deleted, returns -1.
//   - error: an error object if the deletion fails or the username does not exist.
//
// Example:
//
//   id, err := DeleteAccountByUsername("johndoe")
//   if err != nil {
//       log.Printf("Failed to delete account: %v", err)
//   } else {
//       log.Printf("Deleted account with ID: %d", id)
//   }
func DeleteAccountByUsername(username string) (int, error) {
	var accountID int = -1

	err := pool.QueryRow(
		context.Background(),
		"DELETE FROM account WHERE username=$1 RETURNING id",
		username,
	).Scan(&accountID)

	return accountID, err
}

// DeleteAccountById deletes an account from the database using the provided account ID.
//
// It performs a SQL DELETE operation on the `account` table where the `id` matches
// the input parameter. The query uses the RETURNING clause to retrieve the `id` of the
// deleted account.
//
// Parameters:
//   - id: the unique identifier of the account to be deleted.
//
// Returns:
//   - int: the ID of the deleted account. If no account is deleted, returns -1.
//   - error: an error object if the deletion fails or the ID does not exist.
//
// Example:
//
//   deletedID, err := DeleteAccountById(123)
//   if err != nil {
//       log.Printf("Error deleting account: %v", err)
//   } else {
//       log.Printf("Successfully deleted account with ID: %d", deletedID)
//   }
func DeleteAccountById(id int) (int, error) {
	var accountID int = -1
	err := pool.QueryRow(
		context.Background(),
		"DELETE FROM account WHERE id=$1 RETURNING id",
		id,
	).Scan(&accountID)

	return accountID, err
}

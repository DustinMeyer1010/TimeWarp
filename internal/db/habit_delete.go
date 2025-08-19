package db

import "context"

func DeleteHabit(habitID int, accountID int) error {

	_, err := pool.Exec(
		context.Background(),
		"DELETE FROM habit WHERE id = $1 AND account_id = $2",
		habitID, accountID,
	)

	return err
}

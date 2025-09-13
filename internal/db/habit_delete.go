package db

import (
	"context"
	"time"
)

func DeleteHabit(habitID int, accountID int) error {

	_, err := pool.Exec(
		context.Background(),
		"DELETE FROM habit WHERE id = $1 AND account_id = $2",
		habitID, accountID,
	)

	return err
}

func DeleteExtraHabitCompletion(habitID int, date time.Time, extraCompletionCount int) error {
	_, err := pool.Exec(
		context.Background(),
		`DELETE FROM habits_completed
				WHERE id IN (
					SELECT id
					FROM habits_completed
					WHERE habit_id = $1
					AND completion_date = DATE($2)
					ORDER BY id ASC
					LIMIT $3
				)
				RETURNING id;`,
		habitID, date, extraCompletionCount,
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteHabitTimeLogs(timeLogID int) error {
	var habitID int
	var date time.Time

	err := pool.QueryRow(
		context.Background(),
		`DELETE FROM habits_time_logs WHERE id = $1 RETURNING habit_id, "current_time"`,
		timeLogID,
	).Scan(&habitID, &date)

	UpdateCompletion(habitID, date)

	if err != nil {
		return err
	}

	return nil
}

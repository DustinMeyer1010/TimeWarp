package db

import (
	"context"
	"time"
)

type DeletedHabit struct {
	HabitId   int
	AccountId int
}

func NewDeletedHabit() DeletedHabit {
	return DeletedHabit{
		HabitId:   -1,
		AccountId: -1,
	}
}

// Given id for both account and habit and it will delete habit.
// WARNING: will cascade delete all related information
// Time logs and completions logs will be deleted when habit is removed
func DeleteHabitWithTime(habitId int, accountId int) (DeletedHabit, error) {
	deletedHabit := NewDeletedHabit()

	err := pool.QueryRow(
		context.Background(),
		"DELETE FROM habits_with_time WHERE id = $1 AND account_id = $2	RETURNING id, account_id",
		habitId, accountId,
	).Scan(&deletedHabit.HabitId, &deletedHabit.AccountId)

	return deletedHabit, err
}

// If time logs or habits times are updated and completion are higher than they should
// this function will remove all extra completion for specific day
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

// Removes single time log from database
// Automatically updates completion table
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

func DeleteHabitWithoutTime(id, accountID int) (DeletedHabit, error) {
	deletedHabit := NewDeletedHabit()

	err := pool.QueryRow(
		context.Background(),
		"DELETE FROM habits_without_time WHERE id = $1 AND account_id = $2 RETURNING id, account_id",
		id, accountID,
	).Scan(&deletedHabit.HabitId, &deletedHabit.AccountId)

	return deletedHabit, err
}

package db

import (
	"context"
	"time"
)

// DeletedHabit represents a habit that has been deleted from the database.
// It contains the habit's unique identifier and the associated account ID.
//
// Fields:
//   - HabitId: The ID of the deleted habit.
//   - AccountId: The ID of the account that owned the habit.
type DeletedHabit struct {
	HabitId   int
	AccountId int
}

// NewDeletedHabit initializes a DeletedHabit struct with default placeholder values.
// This is typically used to create an empty DeletedHabit before populating it
// with actual data from a database operation.
//
// Returns:
//   - DeletedHabit: A struct with HabitId and AccountId set to -1,
//     indicating an uninitialized or placeholder state.
//
// Example:
//
//	dh := NewDeletedHabit()
//	fmt.Println(dh.HabitId) // Output: -1
func NewDeletedHabit() DeletedHabit {
	return DeletedHabit{
		HabitId:   -1,
		AccountId: -1,
	}
}

// DeleteHabitWithTime deletes a habit entry from the `habits_with_time` table
// using the provided habit ID and account ID. It returns the deleted habit's
// details if the deletion is successful.
//
// Parameters:
//   - habitId: The unique identifier of the habit to delete.
//   - accountId: The identifier of the account that owns the habit.
//
// Returns:
//   - DeletedHabit: A struct containing the ID and account ID of the deleted habit.
//   - error: An error if the deletion fails or no matching record is found.
//
// Behavior:
//
//	This function executes a SQL DELETE statement with a RETURNING clause,
//	which allows it to retrieve the deleted habit's ID and account ID directly.
//	If no matching record is found, the Scan will return an error.
//
// Example:
//
//	deleted, err := DeleteHabitWithTime(321, 654)
//	if err != nil {
//	    log.Printf("Deletion failed: %v", err)
//	} else {
//	    fmt.Printf("Deleted habit: %+v\n", deleted)
//	}
func DeleteHabitWithTime(habitId int, accountId int) (DeletedHabit, error) {
	deletedHabit := NewDeletedHabit()

	err := pool.QueryRow(
		context.Background(),
		"DELETE FROM habits_with_time WHERE id = $1 AND account_id = $2	RETURNING id, account_id",
		habitId, accountId,
	).Scan(&deletedHabit.HabitId, &deletedHabit.AccountId)

	return deletedHabit, err
}

// DeleteExtraHabitCompletion removes a specified number of extra habit completion records
// for a given habit on a specific date from the `habits_completed` table.
//
// Parameters:
//   - habitID: The unique identifier of the habit.
//   - date: The date for which the extra completions should be removed.
//   - extraCompletionCount: The number of extra completion records to delete.
//
// Returns:
//   - error: An error if the deletion fails; otherwise, nil.
//
// Behavior:
//
//	This function executes a SQL DELETE statement that targets a limited number of
//	completion records matching the habit ID and date. It uses a subquery to select
//	the oldest matching records (ordered by ID ascending) and deletes up to the
//	specified limit.
//
// Notes:
//   - The `RETURNING id` clause is included but its result is not used.
//   - If no matching records are found, the function returns nil (no error).
//
// Example:
//
//	err := DeleteExtraHabitCompletion(101, time.Now(), 2)
//	if err != nil {
//	    log.Printf("Failed to delete extra completions: %v", err)
//	}
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

// DeleteHabitTimeLogs deletes a time log entry from the `habits_time_logs` table
// using the provided time log ID, and updates the habit's completion status
// for the associated date.
//
// Parameters:
//   - timeLogID: The unique identifier of the time log entry to delete.
//
// Returns:
//   - error: An error if the deletion or the completion update fails; otherwise, nil.
//
// Behavior:
//   - Executes a SQL DELETE statement with a RETURNING clause to retrieve
//     the associated habit ID and timestamp (`current_time`) of the deleted log.
//   - If the deletion is successful, it calls UpdateCompletion(habitID, date)
//     to refresh the habit's completion status.
//   - Any error from either the deletion or the update is returned.
//
// Notes:
//   - The function ensures that UpdateCompletion is only called if the deletion succeeds.
//   - The returned timestamp is cast to a Go `time.Time` object.
//
// Example:
//
//	err := DeleteHabitTimeLogs(789)
//	if err != nil {
//	    log.Printf("Failed to delete time log or update completion: %v", err)
//	}
func DeleteHabitTimeLogs(timeLogID int) error {
	var habitID int
	var date time.Time

	err := pool.QueryRow(
		context.Background(),
		`DELETE FROM habits_time_logs WHERE id = $1 RETURNING habit_id, "current_time"`,
		timeLogID,
	).Scan(&habitID, &date)

	if err != nil {
		return err
	}

	err = UpdateCompletion(habitID, date)

	return err
}

// DeleteHabitWithoutTime deletes a habit record from the `habits_without_time` table
// based on the provided habit ID and account ID. It returns the deleted habit's
// details and any error encountered during the operation.
//
// Parameters:
//   - id: The unique identifier of the habit to be deleted.
//   - accountID: The identifier of the account that owns the habit.
//
// Returns:
//   - DeletedHabit: A struct containing the ID and account ID of the deleted habit.
//   - error: An error object if the deletion fails or no matching record is found.
//
// Usage:
//
//	deleted, err := DeleteHabitWithoutTime(42, 1001)
//	if err != nil {
//	    log.Fatalf("Failed to delete habit: %v", err)
//	}
//	fmt.Printf("Deleted habit: %+v\n", deleted)
func DeleteHabitWithoutTime(id, accountID int) (DeletedHabit, error) {
	deletedHabit := NewDeletedHabit()

	err := pool.QueryRow(
		context.Background(),
		"DELETE FROM habits_without_time WHERE id = $1 AND account_id = $2 RETURNING id, account_id",
		id, accountID,
	).Scan(&deletedHabit.HabitId, &deletedHabit.AccountId)

	return deletedHabit, err
}

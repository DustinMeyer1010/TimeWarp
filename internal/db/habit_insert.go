package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
)

// CreateHabitWithTime inserts a new habit record into the database.
//
// This function creates a habit with its associated details, such as name,
// description, frequency, and any other fields defined in the models.Habit struct.
// It returns the newly created habit's unique ID on success.
//
// Parameters:
//   - habit: A models.Habit struct containing the data to be stored in the database.
//
// Returns:
//   - int: The ID of the newly inserted habit.
//   - error: An error if the insert operation fails.
func CreateHabitWithTime(habit models.Habit) (int, error) {
	var habitID int

	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO habits_with_time (name, description, account_id, completion_time) VALUES ($1, $2, $3, $4) RETURNING id",
		habit.Name, habit.Description, habit.AccountID, habit.CompletionTime,
	).Scan(&habitID)

	if err != nil {
		return -1, err
	}

	return habitID, nil
}

// CreateHabitTimeLog inserts a single time log record for a habit into the database.
//
// This function records the amount of time spent on a specific habit for a given date.
// It creates a new entry in the habit time log table, associating the provided duration
// with the specified habit and date.
//
// Parameters:
//   - timespent: The amount of time spent on the habit, represented as a models.Duration value.
//   - habitID: The unique identifier of the habit to which the time log belongs.
//   - date: The date on which the time was spent.
//
// Returns an error if the database insert operation fails.
func CreateHabitTimeLog(timespent models.Duration, habitID int, date time.Time) error {
	var HabitsTimeLogs int

	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO habits_time_logs (habit_id, current_time, time_spent) VALUES ($1, $2, $3) RETURNING id",
		habitID, date, timespent,
	).Scan(&HabitsTimeLogs)

	if err != nil {
		return err
	}

	err = UpdateCompletion(habitID, date)

	if err != nil {
		return err
	}

	return nil
}

// CreateHabitCompletion inserts multiple completion records for a given habit on a specific date.
//
// Given a habit ID, a date, and the number of times the habit was completed (`timesCompleted`),
// this function will insert that many rows into the `completion` table. Each inserted row
// represents one completion instance for the provided habit and date.
//
// Parameters:
//   - habitID: The unique identifier of the habit to which the completions belong.
//   - date: The date for which the completions are being recorded.
//   - timesCompleted: The number of completion records to insert.
//
// Returns an error if the database insert operation fails.
func CreateHabitCompletion(habitID int, date time.Time, timesCompleted int) error {

	query := "INSERT INTO habits_with_time_completed (habit_id, completion_date) VALUES "
	placeholders := []string{}
	args := []any{}

	for i := 1; i <= timesCompleted*2; i += 2 {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", i, i+1))
		args = append(args, habitID, date)
	}

	query += strings.Join(placeholders, ",")

	_, err := pool.Exec(
		context.Background(),
		query,
		args...,
	)

	if err != nil {
		return err
	}

	return nil
}

func CreateHabitWithoutTime(habit models.Habit) (int, error) {
	var habitID int = -1
	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO habits_without_time (name, description, account_id) VALUES ($1, $2, $3) RETURNING id",
		habit.Name, habit.Description, habit.AccountID,
	).Scan(&habitID)

	return habitID, err
}

// CreateHabitWithoutTimeCompletion inserts multiple completion records for a habit without time tracking.
//
// This function adds one row per completion into the "habits_without_time_completed" table
// for the specified habit ID and date. The number of rows inserted corresponds to the
// timesCompleted value.
//
// Parameters:
//   - habitID: The unique identifier of the habit.
//   - date: The date on which the completions occurred.
//   - timesCompleted: The number of times the habit was completed on the given date.
//
// Returns:
//   - error: An error if the insert operation fails.
//
// Notes:
//   - The function dynamically builds the SQL INSERT statement with multiple value placeholders.
//   - Each completion is recorded as a separate row with the same habit ID and date.
//   - Uses parameterized queries to prevent SQL injection.
//
// Example:
//
//	err := CreateHabitWithoutTimeCompletion(202, time.Now(), 3)
//	if err != nil {
//	    log.Fatal(err)
//	}
func CreateHabitWithoutTimeCompletion(habitID int, date time.Time, timesCompleted int) error {
	query := "INSERT INTO habits_without_time_completed (habit_id, completion_date) VALUES "
	placeholders := []string{}
	args := []any{}

	for i := 1; i <= timesCompleted*2; i += 2 {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", i, i+1))
		args = append(args, habitID, date)
	}

	query += strings.Join(placeholders, ",")

	_, err := pool.Exec(
		context.Background(),
		query,
		args...,
	)

	if err != nil {
		return err
	}

	return nil
}

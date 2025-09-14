package db

import (
	"context"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

// Add a single habit to the datebase
func CreateHabitWithTime(habit types.Habit) error {
	var habitID int

	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO habits (name, description, account_id, completion_time) VALUES ($1, $2, $3, $4) RETURNING id",
		habit.Name, habit.Description, habit.AccountID, habit.CompletionTime,
	).Scan(&habitID)

	if err != nil {
		return err
	}

	return nil
}

// Add a single habit log to the datebase
func CreateHabitTimeLog(timespent types.Duration, habitID int, date time.Time) error {
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

// Given a number of completions times it will add that many completions rows to completion table for hibit and date
func CreateHabitCompletion(habitID int, date time.Time, timesCompleted int) error {
	println(timesCompleted)
	for i := 0; i < timesCompleted; i++ {
		_, err := pool.Exec(
			context.Background(),
			"INSERT INTO habits_completed (habit_id, completion_date) VALUES ($1, $2)",
			habitID, date,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

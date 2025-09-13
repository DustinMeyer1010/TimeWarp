package db

import (
	"context"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateHabit(habit types.Habit) error {
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

func CreateHabitLog(timespent types.Duration, habitID int, date time.Time) error {
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

func CreateCompletionForHabit(habitID int, date time.Time, addCount int) error {
	for i := 0; i < addCount; i++ {
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

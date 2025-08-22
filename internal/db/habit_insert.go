package db

import (
	"context"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateHabit(habit types.Habit) error {
	var habit_id int

	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO habits (name, description, account_id, completion_time) VALUES ($1, $2, $3, $4) RETURNING id",
		habit.Name, habit.Description, habit.Account_id, habit.CompletionTime,
	).Scan(&habit_id)

	if err != nil {
		return err
	}

	_, err = pool.Exec(
		context.Background(),
		"INSERT INTO habits_time_logs (task_id) VALUES ($1)",
		habit_id,
	)

	if err != nil {
		return err
	}

	return nil
}

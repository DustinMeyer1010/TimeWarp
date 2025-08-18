package db

import (
	"context"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateHabit(habit types.Habit) error {

	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO habit (name, description, account_id) values ($1, $2, $3)",
		habit.Name, habit.Description, habit.Account_id,
	)

	if err != nil {
		return err
	}

	return nil
}

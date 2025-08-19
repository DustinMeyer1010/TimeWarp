package db

import (
	"context"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func GetAllHabitsForUser(id int) (habits []types.Habit, err error) {

	rows, err := pool.Query(
		context.Background(),
		"SELECT * FROM habit WHERE account_id = $1",
		id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var habit types.Habit
		rows.Scan(&habit.ID, &habit.Name, &habit.Description, &habit.Account_id)
		habits = append(habits, habit)
	}
	return
}

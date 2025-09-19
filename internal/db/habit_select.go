package db

import (
	"context"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

// Pulls all habits related for a single users
func GetAllHabitsForUser(UserID int) (habits []types.Habit, err error) {

	rows, err := pool.Query(
		context.Background(),
		"SELECT * FROM habit WHERE account_id = $1",
		UserID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var habit types.Habit
		rows.Scan(&habit.ID, &habit.Name, &habit.Description, &habit.AccountID)
		habits = append(habits, habit)
	}
	return
}

// Gets all the time spent for habit on a given day
func GetHabitTotalTimeSpent(habitID int, date time.Time) (time.Duration, error) {
	var timeSpent time.Duration = time.Duration(time.Second * 0)

	err := pool.QueryRow(
		context.Background(),
		`SELECT COALESCE(SUM(time_spent), INTERVAL '0 HOURS') FROM habits_time_logs WHERE habit_id = $1 AND DATE("current_time") = DATE($2)`,
		habitID, date,
	).Scan(&timeSpent)

	return timeSpent, err
}

// Retrives the completion time for a specific habit
func GetHabitCompletionTime(habitID int) (time.Duration, error) {
	var completionTime time.Duration = time.Duration(time.Second * 0)

	err := pool.QueryRow(
		context.Background(),
		"SELECT completion_time FROM habits WHERE id = $1",
		habitID,
	).Scan(&completionTime)

	return completionTime, err
}

// Retrives the current number of completion for habit on given day
func GetHabitCompletionCount(habitID int, date time.Time) (int, error) {
	var completionCount int = 0

	err := pool.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM habits_completed WHERE habit_id = $1 AND DATE(completion_date) = DATE($2)",
		habitID, date,
	).Scan(&completionCount)

	return completionCount, err
}

func GetHabitWithTime(id, accountID int) (types.Habit, error) {
	var habit = types.Habit{}
	var completionTime time.Duration = time.Duration(time.Second * 0)

	err := pool.QueryRow(
		context.Background(),
		"SELECT id, name, description, completion_time, account_id FROM habits_with_time WHERE id = $1 AND account_id = $2",
		id, accountID,
	).Scan(&habit.ID, &habit.Name, &habit.Description, &completionTime, &habit.AccountID)

	if err != nil {
		return habit, err
	}

	habit.CompletionTime = types.Duration{Duration: completionTime}

	return habit, err
}

func GetHabitWithoutTime(id, accountID int) (types.Habit, error) {
	var habit = types.Habit{}

	err := pool.QueryRow(
		context.Background(),
		"SELECT id, name, description, account_id FROM habits_without_time WHERE id = $1 AND account_id = $2",
		id, accountID,
	).Scan(&habit.ID, &habit.Name, &habit.Description, &habit.AccountID)

	return habit, err
}

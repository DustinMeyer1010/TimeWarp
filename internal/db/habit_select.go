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
	var timeSpent time.Duration

	err := pool.QueryRow(
		context.Background(),
		`SELECT COALESCE(SUM(time_spent), INTERVAL '0 HOURS') FROM habits_time_logs WHERE habit_id = $1 AND DATE("current_time") = DATE($2)`,
		habitID, date,
	).Scan(&timeSpent)

	if err != nil {
		return 0, err
	}

	return timeSpent, nil
}

// Retrives the completion time for a specific habit
func GetHabitCompletionTime(habitID int) (time.Duration, error) {
	var completionTime time.Duration

	err := pool.QueryRow(
		context.Background(),
		"SELECT completion_time FROM habits WHERE id = $1",
		habitID,
	).Scan(&completionTime)

	if err != nil {
		return 0, nil
	}

	return completionTime, nil
}

// Retrives the current number of completion for habit on given day
func GetHabitCompletionCount(habitID int, date time.Time) (int, error) {
	var completionCount int

	err := pool.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM habits_completed WHERE habit_id = $1 AND DATE(completion_date) = DATE($2)",
		habitID, date,
	).Scan(&completionCount)

	if err != nil {
		return 0, err
	}

	return completionCount, nil
}

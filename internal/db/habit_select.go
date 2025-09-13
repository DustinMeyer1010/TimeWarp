package db

import (
	"context"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

// @parameters id - User id to find the habits for
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
		rows.Scan(&habit.ID, &habit.Name, &habit.Description, &habit.AccountID)
		habits = append(habits, habit)
	}
	return
}

func GetAllHabitLogForDay(id int, day time.Time) (habits []types.Habit, err error) {
	return
}

func GetTotalTimeSpentOnHabit(habitID int, date time.Time) (time.Duration, error) {
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

func GetCompletionTimeOnHabit(habitID int) (time.Duration, error) {
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

func GetHabitCompletionCountOnDate(habitID int, date time.Time) (int, error) {
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

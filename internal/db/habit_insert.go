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

	err = CheckForCompletions(habitID, date)

	if err != nil {
		return err
	}

	return nil
}

func CheckForCompletions(habitID int, date time.Time) error {

	timeSpent, err := GetTotalTimeSpentOnDate(habitID, date)

	if err != nil {
		return err
	}

	completionTime, err := GetCompletionTimeForHabit(habitID)

	if err != nil {
		return err
	}

	newTotalCompletions := int(timeSpent.Seconds() / completionTime.Seconds())
	completionCount, err := GetHabitCompletionCountOnDate(habitID, date)

	if err != nil {
		return err
	}

	println(completionCount, newTotalCompletions)

	if completionCount > newTotalCompletions {
		err := RemoveExtraCompletion(habitID, date, completionCount-newTotalCompletions)

		if err != nil {
			return err
		}
	} else {
		err := AddCompletionForHabit(habitID, date, newTotalCompletions-completionCount)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTotalTimeSpentOnDate(habitID int, date time.Time) (time.Duration, error) {
	var timeSpent time.Duration

	err := pool.QueryRow(
		context.Background(),
		`SELECT SUM(time_spent) FROM habits_time_logs WHERE habit_id = $1 AND DATE("current_time") = DATE($2)`,
		habitID, date,
	).Scan(&timeSpent)

	if err != nil {
		return 0, err
	}

	return timeSpent, nil
}

func GetCompletionTimeForHabit(habitID int) (time.Duration, error) {
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

func RemoveExtraCompletion(habitID int, date time.Time, extraCompletionCount int) error {
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

func AddCompletionForHabit(habitID int, date time.Time, addCount int) error {
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

package db

import (
	"context"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
)

// GetAllHabitsForUser retrieves all habits associated with a specific user.
//
// It queries the "habit" table in the database for records where the account_id
// matches the provided UserID. Each row is scanned into a models.Habit struct,
// and the resulting slice of habits is returned.
//
// Parameters:
//   - UserID: The unique identifier of the user whose habits are to be fetched.
//
// Returns:
//   - habits: A slice of models.Habit containing all habits linked to the user.
//   - err: An error if the query fails or if scanning rows encounters an issue.
//
// Example:
//
//	habits, err := GetAllHabitsForUser(123)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, h := range habits {
//	    fmt.Println(h.Name)
//	}
func GetAllHabitsForUser(UserID int) (habits []models.Habit, err error) {

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
		var habit models.Habit
		rows.Scan(&habit.ID, &habit.Name, &habit.Description, &habit.AccountID)
		habits = append(habits, habit)
	}
	return
}

// GetHabitWithTimeTotalTimeSpent retrieves the total time spent on a specific habit for a given day.
//
// It queries the "habits_time_logs" table to sum the "time_spent" values for the specified habit ID
// and date. If no records are found, it returns a zero duration.
//
// Parameters:
//   - habitID: The unique identifier of the habit.
//   - date: The date for which the total time spent is to be calculated.
//
// Returns:
//   - time.Duration: The total time spent on the habit for the given date.
//   - error: An error if the query fails or scanning the result encounters an issue.
//
// Notes:
//   - Uses COALESCE to ensure a zero duration is returned if no matching records exist.
//   - The query filters by habit_id and compares the date portion of the "current_time" column.
//
// Example:
//
//	duration, err := GetHabitWithTimeTotalTimeSpent(101, time.Now())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total time spent: %v\n", duration)
func GetHabitWithTimeTotalTimeSpent(habitID int, date time.Time) (time.Duration, error) {
	var timeSpent time.Duration = time.Duration(time.Second * 0)

	err := pool.QueryRow(
		context.Background(),
		`SELECT COALESCE(SUM(time_spent), INTERVAL '0 HOURS') FROM habits_time_logs WHERE habit_id = $1 AND DATE("current_time") = DATE($2)`,
		habitID, date,
	).Scan(&timeSpent)

	return timeSpent, err
}

// GetHabitWithTimeCompletionTime retrieves the completion time for a specific habit.
//
// It queries the "habits_with_time" table to fetch the value of the "completion_time"
// column for the habit identified by the given habitID.
//
// Parameters:
//   - habitID: The unique identifier of the habit.
//
// Returns:
//   - time.Duration: The completion time associated with the habit.
//   - error: An error if the query fails or if scanning the result encounters an issue.
//
// Notes:
//   - If no record is found, the returned duration will be zero.
//   - The function assumes that "completion_time" is stored in a format compatible with time.Duration.
//
// Example:
//
//	completion, err := GetHabitWithTimeCompletionTime(101)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Completion time: %v\n", completion)
func GetHabitWithTimeCompletionTime(habitID int) (time.Duration, error) {
	var completionTime time.Duration = time.Duration(time.Second * 0)

	err := pool.QueryRow(
		context.Background(),
		"SELECT completion_time FROM habits_with_time WHERE id = $1",
		habitID,
	).Scan(&completionTime)

	return completionTime, err
}

// GetHabitWithTimeCompletionCount retrieves the number of times a habit was completed on a specific day.
//
// It queries the "habits_with_time_completed" table to count how many completion records exist
// for the given habit ID and date. If no records are found, the function returns 0.
//
// Parameters:
//   - habitID: The unique identifier of the habit.
//   - date: The date for which the completion count is to be retrieved.
//
// Returns:
//   - int: The number of completions for the habit on the specified date.
//   - error: An error if the query fails or scanning the result encounters an issue.
//
// Notes:
//   - The query uses a CASE expression to ensure a zero count is returned when no records match.
//   - DATE comparison is used to match the day portion of the completion_date.
//
// Example:
//
//	count, err := GetHabitWithTimeCompletionCount(101, time.Now())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Completion count: %d\n", count)
func GetHabitWithTimeCompletionCount(habitID int, date time.Time) (int, error) {
	var completionCount int = 0

	err := pool.QueryRow(
		context.Background(),
		"SELECT CASE WHEN COUNT(*) > 0 THEN COUNT(*) ELSE 0 END FROM habits_with_time_completed WHERE habit_id = $1 AND DATE(completion_date) = DATE($2)",
		habitID, date,
	).Scan(&completionCount)

	return completionCount, err
}

// GetHabitWithTime retrieves a habit with its associated completion time for a specific user.
//
// It queries the "habits_with_time" table using the habit ID and account ID to fetch
// the habit's details including its completion time. The completion time is wrapped
// into a models.Duration type before being returned.
//
// Parameters:
//   - id: The unique identifier of the habit.
//   - accountID: The unique identifier of the user/account that owns the habit.
//
// Returns:
//   - models.Habit: A Habit struct populated with the habit's details and completion time.
//   - error: An error if the query fails or scanning the result encounters an issue.
//
// Notes:
//   - The function ensures that the habit belongs to the specified account.
//   - Completion time is stored as a time.Duration and wrapped in models.Duration.
//
// Example:
//
//	habit, err := GetHabitWithTime(101, 42)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Habit: %s, Completion Time: %v\n", habit.Name,
func GetHabitWithTime(id, accountID int) (models.Habit, error) {
	var habit = models.Habit{}
	var completionTime time.Duration = time.Duration(time.Second * 0)

	err := pool.QueryRow(
		context.Background(),
		"SELECT id, name, description, completion_time, account_id FROM habits_with_time WHERE id = $1 AND account_id = $2",
		id, accountID,
	).Scan(&habit.ID, &habit.Name, &habit.Description, &completionTime, &habit.AccountID)

	if err != nil {
		return habit, err
	}

	habit.CompletionTime = models.Duration{Duration: completionTime}

	return habit, err
}

// GetHabitWithoutTime retrieves a habit that does not have an associated completion time.
//
// It queries the "habits_without_time" table using the habit ID and account ID to fetch
// the habit's basic details. This function is intended for habits that are tracked without
// time-based metrics.
//
// Parameters:
//   - id: The unique identifier of the habit.
//   - accountID: The unique identifier of the user/account that owns the habit.
//
// Returns:
//   - models.Habit: A Habit struct populated with the habit's details.
//   - error: An error if the query fails or scanning the result encounters an issue.
//
// Notes:
//   - The function ensures that the habit belongs to the specified account.
//   - This is useful for habits tracked by count or completion status rather than duration.
//
// Example:
//
//	habit, err := GetHabitWithoutTime(202, 42)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Habit: %s\n", habit.Name)
func GetHabitWithoutTime(id, accountID int) (models.Habit, error) {
	var habit = models.Habit{}

	err := pool.QueryRow(
		context.Background(),
		"SELECT id, name, description, account_id FROM habits_without_time WHERE id = $1 AND account_id = $2",
		id, accountID,
	).Scan(&habit.ID, &habit.Name, &habit.Description, &habit.AccountID)

	return habit, err
}

// GetHabitWithoutTimeCompletionCount retrieves the number of times a habit without time tracking
// was completed on a specific day.
//
// It queries the "habits_without_time_completed" table to count how many completion records exist
// for the given habit ID and date. If no records are found, the function returns 0.
//
// Parameters:
//   - habitID: The unique identifier of the habit.
//   - date: The date for which the completion count is to be retrieved.
//
// Returns:
//   - int: The number of completions for the habit on the specified date.
//   - error: An error if the query fails or scanning the result encounters an issue.
//
// Notes:
//   - The query uses a CASE expression to ensure a zero count is returned when no records match.
//   - DATE comparison is used to match the day portion of the completion_date.
//
// Example:
//
//	count, err := GetHabitWithoutTimeCompletionCount(202, time.Now())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Completion count: %d\n", count)
func GetHabitWithoutTimeCompletionCount(habitID int, date time.Time) (int, error) {
	var completionCount int = 0

	err := pool.QueryRow(
		context.Background(),
		"SELECT CASE WHEN COUNT(*) > 0 THEN COUNT(*) ELSE 0 END FROM habits_without_time_completed WHERE habit_id = $1 AND DATE(completion_date) = DATE($2)",
		habitID, date,
	).Scan(&completionCount)

	return completionCount, err
}

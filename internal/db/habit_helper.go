package db

import "time"

func UpdateCompletion(habitID int, date time.Time) error {

	timeSpent, err := GetTotalTimeSpentOnHabit(habitID, date)

	if err != nil {
		return err
	}

	completionTime, err := GetCompletionTimeOnHabit(habitID)

	if err != nil {
		return err
	}

	newTotalCompletions := int(timeSpent.Seconds() / completionTime.Seconds())
	completionCount, err := GetHabitCompletionCountOnDate(habitID, date)

	if err != nil {
		return err
	}

	if completionCount > newTotalCompletions {
		err := DeleteExtraHabitCompletion(habitID, date, completionCount-newTotalCompletions)

		if err != nil {
			return err
		}
	} else {
		err := CreateCompletionForHabit(habitID, date, newTotalCompletions-completionCount)
		if err != nil {
			return err
		}
	}
	return nil
}

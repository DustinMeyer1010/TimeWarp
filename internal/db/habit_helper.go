package db

import (
	"time"
)

// Adjust the completion row amount base on updates to time logs and habit
func UpdateCompletion(habitID int, date time.Time) error {

	timeSpent, err := GetHabitTotalTimeSpent(habitID, date)

	if err != nil {
		return err
	}

	completionTime, err := GetHabitCompletionTime(habitID)

	if err != nil {
		return err
	}

	newTotalCompletions := int(timeSpent.Seconds() / completionTime.Seconds())

	completionCount, err := GetHabitCompletionCount(habitID, date)

	if err != nil {
		return err
	}

	if completionCount > newTotalCompletions {
		err := DeleteExtraHabitCompletion(habitID, date, completionCount-newTotalCompletions)

		if err != nil {
			return err
		}
	} else {

		err := CreateHabitCompletion(habitID, date, newTotalCompletions-completionCount)
		if err != nil {
			return err
		}
	}
	return nil
}

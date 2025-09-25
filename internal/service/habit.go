package service

import (
	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/models"
)

func CreateHabit(habit models.Habit, claims models.Claims) (int, error) {
	habit.AccountID = claims.ID

	return db.CreateHabitWithTime(habit)
}

func DeleteHabitWithTime(id int, account_id int) (db.DeletedHabit, error) {
	return db.DeleteHabitWithTime(id, account_id)
}

func DeleteHabitWithoutTime(id int, account_id int) (db.DeletedHabit, error) {
	return db.DeleteHabitWithoutTime(id, account_id)
}

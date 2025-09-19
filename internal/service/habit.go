package service

import (
	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func CreateHabit(habit types.Habit, claims types.Claims) (int, error) {
	habit.AccountID = claims.ID

	return db.CreateHabitWithTime(habit)
}

func DeleteHabit(id int, account_id int) error {
	return db.DeleteHabitWithTime(id, account_id)
}

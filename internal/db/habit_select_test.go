package db

import (
	"testing"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetHabitWithTime(t *testing.T) {

	var habit types.Habit = types.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)
	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithTime(id, habit.AccountID)

	assert.NoError(t, err)

	assertHabitsAreEqual(t, returnedHabit, habit)
}

func TestGetHabitWithTimeNotExist(t *testing.T) {
	DeleteHabitWithTime(1, Account1.ID) // Make sure there is no habit that exist with this habit

	habit, err := GetHabitWithTime(1, Account1.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, habit)

}

func TestGetHabitWithoutTimeNotExist(t *testing.T) {
	DeleteHabitWithoutTime(1, Account1.ID)

	habit, err := GetHabitWithoutTime(1, Account1.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, habit)

}

func TestGetHabitWithTimeExist(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)

	assert.NotEqual(t, id, -1)
	assert.NoError(t, err)

	returnedHabit, err := GetHabitWithTime(id, habit.AccountID)

	assert.NoError(t, err)
	assertHabitsAreEqual(t, returnedHabit, habit)

}

func TestGetHabitWithoutTimeExist(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:        "Test",
		Description: "This is a valid habit",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithoutTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithoutTime(id, habit.AccountID)

	assert.NoError(t, err)
	assertHabitsAreEqual(t, returnedHabit, habit)

}

func TestGetHabitWithTimeIDAndAccountNoMatch(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithTime(id, Account2.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, returnedHabit)

}

func TestGetHabitWithoutTimeIDandAccountNoMatch(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:        "Test",
		Description: "This is a valid habit",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithoutTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithTime(id, Account2.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, returnedHabit)
}

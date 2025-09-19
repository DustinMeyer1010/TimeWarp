package db

import (
	"testing"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/stretchr/testify/assert"
)

// Get Habit With Time Created
func TestGHWTC(t *testing.T) {
	t.Logf("%s: Get Habit With Time Created", t.Name())

	var habit models.Habit = models.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: models.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)
	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithTime(id, habit.AccountID)

	assert.NoError(t, err)

	assertHabitsAreEqual(t, returnedHabit, habit)
}

// Get Habit Without Time Created
func TestGHWOTC(t *testing.T) {
	t.Logf("%s: Get Habit Without Time Not Created", t.Name())
	var habit models.Habit = models.Habit{
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

// Get Habit With Time Not Created
func TestGHWTNC(t *testing.T) {
	t.Logf("%s: Get Habit With Time Not Created", t.Name())
	DeleteHabitWithTime(1, Account1.ID) // Make sure there is no habit that exist with this habit

	habit, err := GetHabitWithTime(1, Account1.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, habit)

}

// Get Habit Without Time Not Created
func TestGHWOTNC(t *testing.T) {
	t.Logf("%s: Get Habit Without Time Not Created", t.Name())
	DeleteHabitWithoutTime(1, Account1.ID)

	habit, err := GetHabitWithoutTime(1, Account1.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, habit)

}

// Get Habit With Time Id And Account Not Matching
func TestGHWTIAANM(t *testing.T) {
	t.Logf("%s: Get Habit With Time Id And Account Not Matching", t.Name())
	var habit models.Habit = models.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: models.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	returnedHabit, err := GetHabitWithTime(id, Account2.ID)

	assert.Error(t, err)
	assertHabitIsEmpty(t, returnedHabit)

}

// Get Habit Without Time Id And Account Not Matching
func TestGHWOTIAANM(t *testing.T) {
	t.Logf("%s: Get Habit With Time Id And Account Not Matching", t.Name())
	var habit models.Habit = models.Habit{
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

package db

import (
	"testing"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/stretchr/testify/assert"
)

// Delete Habit With Time
func TestDHWT(t *testing.T) {
	t.Logf("%s: Delete Habit With Time", t.Name())

	habit := models.Habit{
		Name:        "DHWT",
		Description: "Delete Habit With Time",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)

	deletehabit, err := DeleteHabitWithTime(id, habit.AccountID)

	assert.NoError(t, err)
	assert.NotEqual(t, deletehabit.AccountId, -1)
	assert.NotEqual(t, deletehabit.HabitId, -1)
	assert.Equal(t, deletehabit.AccountId, habit.AccountID)
	assert.Equal(t, deletehabit.HabitId, deletehabit.HabitId)

}

// Delete Habit With Time Verify All Completions Are Deleted
func TestDHWTVACAD(t *testing.T) {
	t.Logf("%s: Delete Habit With Time Verify All Completions Are Deleted", t.Name())

	habit := models.Habit{
		Name:        "DHWTVACAD",
		Description: "Delete Habit With Time Verify All Completions Are Deleted",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)

	err = CreateHabitCompletion(id, time.Now(), 10)

	assert.NoError(t, err)

	deletehabit, err := DeleteHabitWithTime(id, habit.AccountID)

	assert.NoError(t, err)
	assert.NotEqual(t, deletehabit.AccountId, -1)
	assert.NotEqual(t, deletehabit.HabitId, -1)
	assert.Equal(t, deletehabit.AccountId, habit.AccountID)
	assert.Equal(t, deletehabit.HabitId, deletehabit.HabitId)

	count, err := GetHabitWithTimeCompletionCount(id, time.Now())

	assert.NoError(t, err)
	assert.Equal(t, count, 0)

}

// Delete Habit Without Time Verify All Completions Are Deleted
func TestDHWOTVACAD(t *testing.T) {
	t.Logf("%s: Delete Habit Without Time Verify All Completions Are Deleted", t.Name())

	habit := models.Habit{
		Name:        "DHWOTVACAD",
		Description: "Delete Habit Without Time Verify All Completions Are Deleted",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithoutTime(habit)

	assert.NoError(t, err)

	err = CreateHabitWithoutTimeCompletion(id, time.Now(), 10)

	assert.NoError(t, err, id)

	/*
		deletehabit, err := DeleteHabitWithoutTime(id, habit.AccountID)

		assert.NoError(t, err)
		assert.NotEqual(t, deletehabit.AccountId, -1)
		assert.NotEqual(t, deletehabit.HabitId, -1)
		assert.Equal(t, deletehabit.AccountId, habit.AccountID)
		assert.Equal(t, deletehabit.HabitId, deletehabit.HabitId)


	*/
	count, err := GetHabitWithoutTimeCompletionCount(id, time.Now())
	assert.NoError(t, err)
	assert.Equal(t, count, 0)
}

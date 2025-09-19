package db

import (
	"testing"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateHabitWithTimeValid(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)
}

func TestCreateSameHabitTwiceValid(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:           "TwoHabits",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: types.Duration{Duration: time.Hour * 1},
	}
	id, err := CreateHabitWithTime(habit)
	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	id2, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	assert.NotEqual(t, id, id2)
}

func TestCreateHabitWithoutTime(t *testing.T) {
	var habit types.Habit = types.Habit{
		Name:        "TwoHabits",
		Description: "This is a valid habit",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithoutTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)
}

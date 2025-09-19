package db

import (
	"testing"
	"time"

	"github.com/DustinMeyer1010/TimeWarp/internal/models"
	"github.com/stretchr/testify/assert"
)

// Create Habit With Time
func TestCHWT(t *testing.T) {
	t.Logf("%s: Create Habit With Time", t.Name())
	var habit models.Habit = models.Habit{
		Name:           "Test",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: models.Duration{Duration: time.Hour * 1},
	}

	id, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)
}

// Create Same Habit Twice
func TestCSHT(t *testing.T) {
	t.Logf("%s: Create Same Habit Twice", t.Name())
	var habit models.Habit = models.Habit{
		Name:           "TwoHabits",
		Description:    "This is a valid habit",
		AccountID:      Account1.ID,
		CompletionTime: models.Duration{Duration: time.Hour * 1},
	}
	id, err := CreateHabitWithTime(habit)
	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	id2, err := CreateHabitWithTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)

	assert.NotEqual(t, id, id2)
}

// Create Habit Without Time
func TestCHWOT(t *testing.T) {
	t.Logf("%s: Create Same Habit Twice", t.Name())
	var habit models.Habit = models.Habit{
		Name:        "TwoHabits",
		Description: "This is a valid habit",
		AccountID:   Account1.ID,
	}

	id, err := CreateHabitWithoutTime(habit)

	assert.NoError(t, err)
	assert.NotEqual(t, id, -1)
}

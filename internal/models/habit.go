package models

import (
	"encoding/json"
	"time"
)

type Habit struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	AccountID      int      `json:"account_id"`
	CompletionTime Duration `json:"completion_time"`
}

type HabitTimeLogs struct {
	ID          int       `json:"id"`
	HabitID     int       `json:"habit_id"`
	CurrentTime time.Time `json:"current_time"`
	TimeSpent   Duration  `json:"time_spent"`
}

type HabitCompleted struct {
	ID             int       `json:"id"`
	HabitID        int       `json:"habit_id"`
	TimeLogID      int       `json:"time_log_id"`
	CompletionTime time.Time `json:"completion_date"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	dur, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.Duration = dur
	return nil
}

func (d *Duration) IsZero() bool {
	return d.Duration == 0
}

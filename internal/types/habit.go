package types

import "time"

type Habit struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Time        time.Duration `json:"time"`
}

package types

type Habit struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Account_id  int    `json:"account_id"`
	//Time        time.Duration `json:"time"`
}

package types

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Account) Verify(account *Account) bool {
	return a.Password == account.Password
}

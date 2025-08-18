package types

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID           float64 `json:"id"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	RefreshToken string  `json:"refresh_token"`
}

func (a *Account) CheckPassword(account Account) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(account.Password))

	return err == nil
}

func (a *Account) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("failed to encrypt password: %v", err)
	}

	a.Password = string(hashedPassword)

	return nil

}

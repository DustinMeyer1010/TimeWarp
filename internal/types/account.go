package types

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}

func (a *Account) CheckPassword(account Account) bool {
	fmt.Println(a.Password)
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(account.Password))
	fmt.Println(err)
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

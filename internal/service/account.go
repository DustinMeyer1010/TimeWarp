package service

import (
	"fmt"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func DeleteAccount(id int, requestor_username string) error {
	foundAccount, err := db.GetAccountByID(id)

	if err != nil {
		return err
	}

	fmt.Println(err)

	if requestor_username != foundAccount.Username {
		return fmt.Errorf("username do not match")
	}

	return db.DeleteAccount(foundAccount.Username)

}

func CreateAccount(account types.Account) error {
	return db.CreateAccount(account)
}

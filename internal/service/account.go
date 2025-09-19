package service

import (
	"fmt"

	"github.com/DustinMeyer1010/TimeWarp/internal/db"
	"github.com/DustinMeyer1010/TimeWarp/internal/types"
)

func DeleteAccount(id int, requestor_username string) (int, error) {
	foundAccount, err := db.GetAccountByID(id)

	if err != nil {
		return -1, err
	}

	if requestor_username != foundAccount.Username {
		return -1, fmt.Errorf("Unauthorized")
	}

	return db.DeleteAccountByUsername(foundAccount.Username)

}

func CreateAccount(account types.Account) (int, error) {
	return db.CreateAccount(account)
}

/*
 - account:
 	in - name, unique id, state
 	out - account number
 - list all accounts (Get):
 	out - list of objects: {name, id, account number, account status, account creation date, balance}
 - account deactivation (no delete):
 	in - account number
 - list all transactions for specific account (get):
 	in - account number, transaction amount, transaction type, comment,
	out - transaction id, current balance
 - apply debits or credits to account:
 	in - account number, transaction amount, transaction type
	out - transaction id, current balance

*/

package main

import (
	"fmt"
)

func main() {
	names := []string{"Jeremy", "Doug", "Ryan"}
	fmt.Println()
	accounts := create_accounts(names)
}

type Account struct {
	Name          string
	Id            int
	State         string
	AccountNumber int
	Status        string
	DateCreated   string
	Balance       int
}

func create_accounts(names []string) map[string]Account {
	for _, v := range names {

	}
}

func list() map[string]Account {

}

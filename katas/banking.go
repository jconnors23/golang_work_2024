// implement stack

/*

Write a class Account that offers the following methods:
void deposit(int, name)
void withdraw(int, name)
string List(name) - lists transactions in most recent order when given name

An example statement would be:

Date        Amount  Balance
24.12.2015   +500      500
23.8.2016    -100      400

*/

package bank

import (
	"fmt"
	"os"
	"time"
)

type Account struct {
	name         string
	balance      float64
	transactions map[string]float64
}

func (a *Account) withdraw(amount float64) {
	if a.balance < amount {
		fmt.Println("account has insufficient funds.")
		os.Exit(0)
	}
	a.transactions[time.Now().Format("2006-01-02")] = float64(amount)
}

func (a *Account) deposit(amount float64) {
	a.balance += amount
	a.transactions[time.Now().Format("2006-01-02")] = float64(amount)
}

// func (a *Account) list() map[string]int {
// 	return
// }

func main() {
	fmt.Println(list("John"))
}

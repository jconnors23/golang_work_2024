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

package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Account struct {
	name         string
	balance      float64
	transactions map[string]float64
}

func withdraw(a *Account, amount float64) {
	if a.balance < amount {
		fmt.Println("account has insufficient funds.")
		os.Exit(0)
	}
	a.transactions[time.Now().Format("01/01/2024 00:00:00")] = float64(amount)
}

func deposit(a *Account, amount float64) {
	a.balance += amount
	a.transactions[time.Now().Format("01/01/2024 00:00:00")] = float64(amount)
}

// transaction map is inherently unsorted, using a sort method to implement stack behavior
func stack_transactions(a *Account) []time.Time {
	lifo := make([]time.Time, 0, len(a.transactions))
	timeLayout := "01/01/2024 00:00:00"
	for transaction := range a.transactions {
		parsedTime, err := time.Parse(timeLayout, transaction)
		if err != nil {
			fmt.Println("time parsing error", err)
			os.Exit(0)
		}
		lifo = append(lifo, parsedTime)
	}
	sort.Slice(lifo, func(i, j int) bool {
		return lifo[i].After(lifo[j]) // will sort with most recent transactions at top
	})
	return lifo
}

func list(a *Account) {
	lifo := stack_transactions(a)
	timeLayout := "01/01/2024 00:00:00"
	for _, t := range lifo {
		fmt.Println(t.Format(timeLayout), a.transactions[t.Format(timeLayout)]) // print transactions by most recent
	}
}

func main() {
	a := Account{
		name:    "test",
		balance: 1000,
		transactions: map[string]float64{
			"24.12.2015": 500,
			"23.8.2016":  100,
		},
	}
	list(&a)
}

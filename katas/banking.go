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
	transactions map[time.Time]float64
}

func (a *Account) withdraw(amount float64) {
	if a.balance < amount {
		fmt.Println("account has insufficient funds.")
		os.Exit(0)
	}
	a.transactions[time.Now()] = float64(amount)
}

func (a *Account) deposit(amount float64) {
	a.balance += amount
	a.transactions[time.Now()] = float64(amount)
}

// transaction map is inherently unsorted, using a sort method to implement stack behavior
func (a *Account) stack_transactions() []time.Time {
	lifo := make([]time.Time, 0, len(a.transactions))
	for transaction := range a.transactions {
		lifo = append(lifo, transaction)
	}
	sort.Slice(lifo, func(i, j int) bool {
		return lifo[i].After(lifo[j])
	})
	return lifo
}

func (a *Account) list() {
	lifo := a.stack_transactions()
	for _, t := range lifo {
		fmt.Println(t, a.transactions[t]) // print transactions by most recent
	}
}

// func (a *Account) list() map[string]int {
// 	return
// }

func main_bank() {
	fmt.Println()
}

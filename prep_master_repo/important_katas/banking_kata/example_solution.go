package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	a := Account{
		name:    "test",
		balance: 1000,
		transactions: map[string]Transaction{
			"2015-04-05": {500, "deposit"},
			"2016-09-09": {100, "withdraw"},
			"2013-01-01": {1000, "deposit"},
			"2001-01-01": {108, "deposit"},
			"2023-06-08": {450, "withdraw"},
		},
	}
	// withdraw(&a, 200)
	deposit(&a, 750)
	list(&a)

}

type Transaction struct {
	amount float64
	form   string
}

type Account struct {
	name         string
	balance      float64
	transactions map[string]Transaction
}

func withdraw(a *Account, amount float64) {
	if a.balance < amount {
		fmt.Println("account has insufficient funds.")
		os.Exit(0)
	}
	a.balance -= amount
	transaction := Transaction{amount, "withdraw"}
	a.transactions[time.Now().Format("2006-01-02")] = transaction
}

func deposit(a *Account, amount float64) {
	a.balance += amount
	transaction := Transaction{amount, "deposit"}
	a.transactions[time.Now().Format("2006-01-02")] = transaction
}

// transaction map is inherently unsorted, using a sort method to implement stack behavior
func stack_transactions(a *Account) []time.Time {
	lifo := make([]time.Time, 0, len(a.transactions))
	timeLayout := "2006-01-02"
	for transaction := range a.transactions {
		parsedTime, err := time.Parse(timeLayout, transaction)
		if err != nil {
			fmt.Println("error", err)
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
	timeLayout := "2006-01-02"
	for _, t := range lifo {
		fmt.Println(t.Format(timeLayout), a.transactions[t.Format(timeLayout)]) // print transactions by most recent
	}
}

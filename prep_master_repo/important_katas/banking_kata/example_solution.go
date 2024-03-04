/*
 Open ended design for this Kata. This design features:
- stack for printing transaction history
- deposit, withdraw, list transaction methods
- os exit for specific inputs / errors
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	a := Account{ // example account A
		name:    "Alice",
		balance: 1000,
		transactions: map[string]Transaction{
			"2015-04-05": {500, "deposit"},
			"2016-09-09": {100, "withdraw"},
			"2013-01-01": {1000, "deposit"},
			"2001-01-01": {108, "deposit"},
			"2023-06-08": {450, "withdraw"},
		},
	}
	b := Account{ // example account B
		name:    "Bob",
		balance: 1000,
		transactions: map[string]Transaction{
			"2012-04-05": {5000, "deposit"},
			"2012-09-09": {1030, "withdraw"},
			"2013-01-01": {100, "deposit"},
			"2006-01-01": {450, "deposit"},
			"2021-06-08": {1450, "withdraw"},
		},
	}
	// example methods
	withdraw(&b, 200)
	deposit(&a, 750)
	list(&a)
	list(&b)
}

type Transaction struct { // each transaction will have dollar amounts and a string detailing if a withdrawal or deposit
	amount float64
	form   string
}

type Account struct { // each account will have an individual's name, account balance, and a map with keys of dates and values of transactions
	name         string
	balance      float64
	transactions map[string]Transaction
}

func withdraw(a *Account, amount float64) {
	if a.balance < amount { // cant withdraw more than account balance
		fmt.Println("account has insufficient funds.")
		os.Exit(0)
	}
	a.balance -= amount // modify account balance based on current transaction amount
	transaction := Transaction{amount, "withdraw"}
	a.transactions[time.Now().Format("2006-01-02")] = transaction // add current transaction to transaction history associated with this account
}

func deposit(a *Account, amount float64) {
	a.balance += amount
	transaction := Transaction{amount, "deposit"}
	a.transactions[time.Now().Format("2006-01-02")] = transaction
}

// transaction map is unsorted, using a sort method to implement stack behavior
func stack_transactions(a *Account) []time.Time {
	lifo := make([]time.Time, 0, len(a.transactions)) // make a slice of time values
	timeLayout := "2006-01-02"                        // layout for how dates will be formatted in parsing
	for transaction := range a.transactions {         // iterate among transaction values in account
		parsedTime, err := time.Parse(timeLayout, transaction)
		if err != nil {
			fmt.Println("error", err)
			os.Exit(0)
		}
		lifo = append(lifo, parsedTime) // add formatted time into stack
	}
	sort.Slice(lifo, func(i, j int) bool { // sort the stack to have most recent times on top
		return lifo[i].After(lifo[j])
	})
	return lifo
}

func list(a *Account) {
	fmt.Printf("Account: %s\n", a.name)
	fmt.Printf("Balance: $%.2f\n", a.balance)
	lifo := stack_transactions(a) // get the transactions by most recent
	timeLayout := "2006-01-02"
	for _, date := range lifo {
		// print the date & the value in transaction map associated with that date (the amount and deposit/withdrawal)
		fmt.Println(date.Format(timeLayout), a.transactions[date.Format(timeLayout)])
	}
}

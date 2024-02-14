package main

import "fmt"

type Accounts struct {
	Users map[int][]Transactions
}

type Transactions struct {
	price    float64
	category string
	date     string
}

// iterate filter if Transaction = x category & y date then increment sum for comparisons

func main() {
	// format account
	// need to find way to have multiple Transactions for each user and filter towards them interface?
	accounts := Accounts{
		Users: map[int][]Transactions{
			12389: {
				{12.50, "entertainment", "2006-01-02"},
				{14.78, "entertainment", "2006-09-02"},
				{1.50, "food", "2006-05-02"},
			},
			124673: {
				{12.50, "entertainment", "2006-01-02"},
				{14.78, "entertainment", "2006-09-02"},
				{1.50, "food", "2006-05-02"},
			},
		},
	}
	fmt.Println(spendingAlert(&accounts, 12389))
}

func spendingAlert(a *Accounts, id int) bool {
	//var current_month, previous_month := float64
	for _, transaction := range a.Users[id] {

	}

	return true
}

// func getUser(name string, Account *Accounts) User {
// 	return Account.Users[name]
// }

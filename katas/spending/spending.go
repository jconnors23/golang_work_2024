package main

import "fmt"

type Accounts struct {
	Users map[string]Payments
}

type Payments struct {
	price    float64
	category string
	date     string
}

// iterate filter if payment = x category & y date then increment sum for comparisons

func main() {
	// format account
	// need to find way to have multiple payments for each user and filter towards them interface?
	accounts := Accounts{
		Users: map[string]Payments{
			"Josh": Payments{
				12.50, "entertainment", "2006-01-02",
			},
			"Charlie": Payments{
				10.00, "food", "2008-04-07",
			},
		},
	}
	fmt.Println(spendingAlert(&accounts))
}

func spendingAlert(a *Accounts) {

}

// func getUser(name string, Account *Accounts) User {
// 	return Account.Users[name]
// }

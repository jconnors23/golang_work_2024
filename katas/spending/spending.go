package main

import "fmt"

type User struct {
	Name     string
	Payments map[string]Payment
}

type Payment struct {
	price    float64
	category string
	date     string
}

// iterate filter if payment = x category & y date then increment sum for comparisons

func main() {
	fmt.Println(spendingAlert())
}

func spendingAlert() {

}

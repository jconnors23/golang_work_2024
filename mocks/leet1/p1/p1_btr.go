package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	fmt.Println(invalidTransactions(input))
}

func invalidTransactions(transactions []string) []string {
	flags := []string{}
	for i := 0; i < len(transactions); i++ {
		current := transactions[i]
		if checkAmount(current) {
			if contains(current, flags) == false {
				flags = append(flags, current)
			}
		}
		for j := 0; j < len(transactions); j++ {
			compare := transactions[j]
			fmt.Println(compare)
			if compareName(current, compare) {
				if compareTime(current, compare) {
					if compareCity(current, compare) {
						if contains(current, flags) == false {
							flags = append(flags, current)
						}
						if contains(compare, flags) == false {
							flags = append(flags, compare)
						}
					}
				}
			}
		}
	}
	return flags
}

func findComma(str string, after bool) string {
	comma := strings.Index(str, ",")
	if after {
		// everything after
		return str[comma+1:]
	} else {
		// everything before
		return str[:comma]
	}
}

func compareName(a, b string) bool {
	a = findComma(a, false)
	b = findComma(b, false)
	return a == b
}

func compareTime(a, b string) bool {
	a = findComma(a, true)  // after first comma
	a = findComma(a, false) // before second
	b = findComma(b, true)
	b = findComma(b, false)
	min1, err := strconv.Atoi(a)
	if err != nil {
		_ = err
	}
	min2, err := strconv.Atoi(b)
	if err != nil {
		_ = err
	}
	return math.Abs(float64(min1-min2)) <= 60
}

func compareCity(a, b string) bool {
	a = findComma(a, true) // after first comma
	a = findComma(a, true) // after second
	a = findComma(a, true) // after third
	b = findComma(b, true)
	b = findComma(b, true)
	b = findComma(b, true)
	return a != b
}

func contains(str string, t []string) bool {
	contains := false
	for _, v := range t {
		if str == v {
			contains = true
		}
	}
	return contains
}

func checkAmount(a string) bool {
	a = findComma(a, true)  // after first
	a = findComma(a, true)  // after second
	a = findComma(a, false) // before third
	amount, err := strconv.Atoi(a)
	if err != nil {
		_ = err
	}
	return amount > 1000
}

// package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// 	"strings"
// )

// /*
// 	You are given an array of strings transaction where
// 	transactions[i] consists of comma-separated values representing
// 	the name,
// 	time (in minutes),
// 	amount,
// 	and city of the transaction.

// Return a list of transactions that are possibly invalid.

// str conv handling

// for each transaction, compare it to rest
// have checks for each statement in prompt
// if clear checks & not already in flag array, push it to flag array
// return flag array

// */

// func main() {
// 	input := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
// 	fmt.Println(invalidTransactions(input))
// }

// func invalidTransactions(transactions []string) []string {
// 	flags := []string{}
// 	for i := 0; i < len(transactions); i++ {
// 		current, original := transactions[i], transactions[i]
// 		fmt.Println(current, original)
// 		for j := 0; j < len(transactions); j++ {
// 			compare, originalComp := transactions[j], transactions[j]
// 			fmt.Println(compare, originalComp)
// 			if sameName(findComma(current, false), findComma(compare, false)) {
// 				fmt.Println("same name")
// 				current = findComma(current, true)
// 				compare = findComma(compare, true)
// 			} else {
// 				continue
// 			}
// 			if within60(findComma(current, false), findComma(compare, false)) {
// 				fmt.Println("within 60")
// 				current = findComma(current, true)
// 				compare = findComma(compare, true)
// 			} else {
// 				continue
// 			}

// 			if largeAmount(findComma(current, false)) {
// 				if contains(original, flags) == false {
// 					flags = append(flags, original)
// 				}
// 			} else {
// 				current = findComma(current, true)
// 			}
// 			if largeAmount(findComma(compare, false)) {
// 				if contains(originalComp, flags) == false {
// 					flags = append(flags, compare)
// 				}
// 			} else {
// 				compare = findComma(compare, true)
// 			}

// 			if diffCity(findComma(current, false), findComma(compare, false)) {
// 				if contains(original, flags) == false {
// 					flags = append(flags, original)
// 				}
// 				if contains(originalComp, flags) == false {
// 					flags = append(flags, compare)
// 				}
// 			} else {
// 				continue
// 			}
// 		}
// 	}
// 	return flags
// }

// func findComma(str string, truncate bool) string {
// 	comma := strings.Index(str, ",")
// 	if comma == -1 {
// 		return str
// 	}
// 	if truncate {
// 		// everything after
// 		return str[comma:]
// 	} else {
// 		// everything before
// 		return str[:comma]
// 	}
// }

// func sameName(a, b string) bool {
// 	return a == b
// }

// func within60(a, b string) bool {
// 	min1, err := strconv.Atoi(a)
// 	if err != nil {
// 		_ = err
// 	}
// 	min2, err := strconv.Atoi(b)
// 	if err != nil {
// 		_ = err
// 	}
// 	return math.Abs(float64(min1-min2)) <= 60
// }

// func largeAmount(a string) bool {
// 	amount, err := strconv.Atoi(a)
// 	if err != nil {
// 		_ = err
// 	}
// 	return amount > 1000
// }

// func diffCity(a, b string) bool {
// 	return a != b
// }

// func contains(str string, t []string) bool {
// 	for _, v := range t {
// 		if str == v {
// 			return true
// 		}
// 	}
// 	return false
// }

package main

package main

import "fmt"

func main() {
	inputArray := []string{"aba", "aa", "ad", "vcd", "aba"}
	fmt.Println(solution(inputArray))
}

func solution(inputArray []string) []string {
	longestStrings := []string{} // init empty slice of strings for return
	max := 0
	for _, str := range inputArray { // iterate over our input using range keyword
		if len(str) > max { // check value in input Array, compare to max length
			max = len(str) // set max equal to longest string length
		}
	}
	for i := 0; i < len(inputArray); i++ { // traditional for loop, adding longest strings to return slice
		test := inputArray[i]
		if len(test) == max {
			longestStrings = append(longestStrings, test)
		}
	}
	return longestStrings
}

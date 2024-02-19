package main

import "fmt"

func main() {
	inputArray := []string{"aba", "aa", "ad", "vcd", "aba"}
	fmt.Println(solution(inputArray))
}

func solution(inputArray []string) []string {

	longestStrings := []string{}
	max := 0
	for _, str := range inputArray {
		if len(str) > max {
			max = len(str)
		}
	}
	for i := 0; i < len(inputArray); i++ {
		test := inputArray[i]
		if len(test) == max {
			longestStrings = append(longestStrings, test)
		}
	}
	return longestStrings
}

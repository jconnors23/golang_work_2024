package main

import "fmt"

func main() {
	fmt.Println(solution(2024))
}

func solution(year int) int {
	// for every one hundred years, increment our count by one
	count := 0
	for ; year > 0; year -= 100 { // python equivalent: while year > 0, year -= 100
		count++
	}
	return count
}

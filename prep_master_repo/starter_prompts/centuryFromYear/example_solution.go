package main

import "fmt"

func main() {
	fmt.Println(solution(2024))
}

func solution(year int) int {
	count := 0
	for ; year > 0; year -= 100 {
		count++
	}
	return count
}

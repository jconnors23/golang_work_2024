package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 6, -2, -5, 7, 3}
	fmt.Println(solution(arr))
}

func solution(inputArray []int) int {
	max := -math.MaxInt
	for i := 0; i < len(inputArray)-1; i++ {
		prod := inputArray[i] * inputArray[i+1]
		if prod > max {
			max = prod
		}
	}
	return max
}

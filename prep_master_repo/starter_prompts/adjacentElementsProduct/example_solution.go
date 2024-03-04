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
	max := -math.MaxInt // init max, accounting for negative solutions in array
	for i := 0; i < len(inputArray)-1; i++ {
		prod := inputArray[i] * inputArray[i+1]
		if prod > max { // comparison, make max equal to whatever largest value is in comparison
			max = prod
		}
	}
	return max
}

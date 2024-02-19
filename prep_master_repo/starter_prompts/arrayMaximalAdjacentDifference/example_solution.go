package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 4, 1, 0}
	fmt.Println(solution(arr))
}
func solution(arr []int) int {
	max := 0.0
	for i := 0; i < len(arr); i++ {
		current := arr[i]
		for j := 0; j < len(arr); j++ {
			if math.Abs(float64(i-j)) <= 1 {
				diff := float64(current - arr[j])
				if math.Abs(diff) > max {
					max = diff
				}
			}
		}
	}
	return int(max)
}

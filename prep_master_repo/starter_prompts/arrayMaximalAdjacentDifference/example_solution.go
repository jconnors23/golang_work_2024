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
	var max float64                 // init max as float for comparison purposes
	for i := 0; i < len(arr); i++ { // double for loop
		current := arr[i]
		for j := 0; j < len(arr); j++ {
			if math.Abs(float64(i-j)) <= 1 { // check positions of elements in arr for adjacency
				diff := float64(current - arr[j]) // if adjacent, find diff
				if math.Abs(diff) > max {         // compare absolute value of diff to max difference
					max = diff
				}
			}
		}
	}
	return int(max) // return type of int
}

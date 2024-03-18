package main

import (
	"fmt" // fmt.Println for testing
)

func main() {
	nums := []int{3, 4, 5, 6, 7}
	target := 8
	result := twoSum(nums, target)
	fmt.Printf("Result: %v\n", result)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		j := 1 + i
		for j < len(nums) && nums[i] != nums[j] {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	// If no pair is found return an empty slice
	return []int{}
}

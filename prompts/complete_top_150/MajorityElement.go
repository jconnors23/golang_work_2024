package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(MajorityElement(nums))
}

func MajorityElement(nums []int) int {
	floor := float64(len(nums) / 2)
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		if float64(countOf(curr, nums)) > floor {
			return curr
		}
	}
	return -1
}

func countOf(curr int, nums []int) int {
	count := 0
	for _, v := range nums {
		if v == curr {
			count++
		}
	}
	return count
}

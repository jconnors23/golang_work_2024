// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

package main

import (
	"fmt"
)

func main() {
	nums := []int{1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) []int {
	floorLen := float64(len(nums) / 3)
	majors := []int{}
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		count := 1
		for j := 0; j < len(nums); j++ {
			if nums[j] == current && j != i {
				count++
			}
		}
		if count > int(floorLen) && !contains(majors, current) {
			majors = append(majors, current)
		}
	}
	return majors
}

func contains(majors []int, current int) bool {
	for _, v := range majors {
		if v == current {
			return true
		}
	}
	return false
}

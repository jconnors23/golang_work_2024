package main

import "fmt"

func main() {
	arr := []int{3, 2, 3}
	fmt.Println(majorityElement(arr))
}

func majorityElement(nums []int) []int {
	floorLen := float64(len(nums) / 3) // threshold
	majors := []int{}                  // majority elements return slice
	for i := 0; i < len(nums); i++ {
		current := nums[i] // number in question
		count := 1         // starts at 1 for current num
		for j := 0; j < len(nums); j++ {
			if nums[j] == current && j != i {
				count++ // count all occurences of element
			}
		}
		if count > int(floorLen) && !contains(majors, current) { // return if not already accounted for and greater than threshold
			majors = append(majors, current)
		}
	}
	return majors
}

// helper method to find element that is within array
func contains(majors []int, current int) bool {
	for _, v := range majors {
		if v == current {
			return true
		}
	}
	return false
}

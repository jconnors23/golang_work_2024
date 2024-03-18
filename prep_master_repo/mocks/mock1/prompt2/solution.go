package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	i := 0 // unique elements tracker
	for index, v := range nums {
		_ = v
		fmt.Printf("nums before: %d\n", nums)
		if nums[index] != nums[i] {
			i++
			nums[i] = nums[index] // shift unique element left
		}
	}
	return i + 1
}

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

// func main() {
// 	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	fmt.Println(removeDuplicates(nums))
// }

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

// NIP:
// func removeDuplicates2(nums []int) int {
// 	new := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		ele := nums[i]
// 		if !contains(new, ele) {
// 			new = append(new, ele)
// 		}
// 	}
// 	nums = new
// 	fmt.Printf("%d\n", nums)
// 	return len(new)
// }

// func contains(new []int, target int) bool {
// 	for _, v := range new {
// 		if v == target {
// 			return true
// 		}
// 	}
// 	return false
// }

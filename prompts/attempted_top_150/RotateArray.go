// https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	//nums := []int{-1, -100, 3, 99}
	fmt.Println(rotate(nums, 17))
	/*
		Output
		[11,12,13,1,2,9,4,5,6,7,8,10,3]
		Expected
		[10,11,12,13,1,2,3,4,5,6,7,8,9]
	*/
}

func rotate(nums []int, k int) []int {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k+1:])

	return nums
	//fmt.Println(nums)
}

func reverse(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return i > j
	})
	return nums
}

// func rotate(nums []int, k int) []int {
// 	rotation := []int{}
// 	num_shift := len(nums) - k
// 	rest, index := num_shift, num_shift
// 	if len(nums) - k < 0 {
// 		handlenegatives(nums, k)
// 	}
// 	for k > 0 { // shift back k elements to front
// 		rotation = append(rotation, nums[index])
// 		k--
// 	}
// 	if rest > 0 {
// 		rotation = append(rotation, nums[0:rest]...) // append rest of original slice elements
// 	}
// 	copy(nums, rotation)
// 	return nums
// }

// func handlenegatives(nums, k) {

// }

// fmt.Printf("before rotation %d\n", rotation)
// fmt.Printf("before append %d\n", rotation)
// fmt.Printf("after rotation %d\n", rotation)
// fmt.Printf("after append %d\n", rotation)
// func rotate(nums []int, k int) []int {

// }

/*
 [ 1 8 3 9]  [ 1 8 3 9 _ _ _ _ ]
 can we swap with elements k units behind? - take nums[-k] and swap with nums[-k - k ? ]
element: nums[len(nums)-k] swap to: nums[len(nums)-k]
iterate -

*/

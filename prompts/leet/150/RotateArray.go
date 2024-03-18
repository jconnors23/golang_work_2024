//https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150

package main

import "fmt"

/*
k = 3
1,2,3,4,5,6,7 - > 5,6,7,1,2,3,4

a. create new arr of same size
b. while < k, push back els of nums to new arr
c. push back portion of slice to new arr & return

account for large values of k ?

*/

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6, 7} // want 5,6,7,1,2,3,4
	nums := []int{-1, -100, 3, 99} // want 3 99 s-1 -100
	//fmt.Println(removeDuplicates(nums))
	fmt.Println(rotate(nums, 2))
}

// correct but not in place
func rotate(nums []int, k int) []int {
	rotation := []int{}      // init new slice
	for i := k; i > 0; i-- { // shift back k elements to front
		fmt.Printf("back kth element: %d\n", nums[len(nums)-i])
		rotation = append(rotation, nums[len(nums)-i])
		fmt.Printf("post rotation: %d\n", rotation)
	}
	rest := len(nums) - k
	rotation = append(rotation, nums[0:rest]...) // append rest of original slice elements
	copy(nums, rotation)
	return nums
}

// func copyToK(nums []int, k int) []int {
// 	var ans = make([]int, len(nums))

// 	for i := 0; i < len(nums); i++ {
// 		idx := (i + k) % len(nums) // the modulo operator gives us the overflow remainder and that delivers our idx
// 		ans[idx] = nums[i]
// 	}

// 	return ans
// }

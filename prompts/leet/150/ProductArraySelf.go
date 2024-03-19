package main

import "fmt"

// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150

/*
	iterature through nums
	create slice of arr excluding curr index append product of all indices to ans arr[i]
	return - > not o(n)?
*/

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))  // store left products
	right := make([]int, len(nums)) // store right products
	ans := make([]int, len(nums))   // store results
	left[0] = 1                     // leftmost product
	right[len(nums)-1] = 1          // rightmost product
	for i := 1; i < len(nums); i++ {
		j := len(nums) - i - 1 // for right products (iterate backwards)
		left[i] = nums[i-1] * left[i-1]
		right[j] = nums[j+1] * right[j+1]
	}
	for i := 0; i < len(nums); i++ {
		ans[i] = left[i] * right[i] // results in answer array will be left * right
	}
	return ans
}

/*

attempt1:

func productExceptSelf(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		prod := 1
		prod_nums := noIndex(nums, i) // exclude current index in product calculations
		for _, v := range prod_nums { // calculate product
			prod = v * prod
		}
		//fmt.Printf("prodnums = %d, prod = %d, nums = %d\n", prod_nums, prod, nums)
		ans = append(ans, prod)
	}
	return ans
}

func noIndex(slice []int, index int) []int {
	if index < 0 || index >= (len(slice)) {
		return slice
	}
	newSlice := append(slice[:index], slice[index+1:]...) // modifies original slice?
	return newSlice
}

attempt 2:


func productExceptSelf(nums []int) []int {
	ans := []int{}
	prod := 1
	for i := 0; i < len(nums); i++ {
		left := nums[:i]
		right := nums[i+1:]
		for _, v := range left {
			prod = prod * v
		}
		for _, v := range right {
			prod = prod * v
		}
		ans = append(ans, prod)
		prod = 1
	}
	return ans
}




*/

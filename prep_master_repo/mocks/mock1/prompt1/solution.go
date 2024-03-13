package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	nums2 := []int{1, 2, 2}
	m := 6
	n := 3
	fmt.Println(merge(nums1, m, nums2, n))
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	count := 0 // keep track of which element to merge from nums 2
	for index, v := range nums1 {
		if v == 0 && index >= m { // make sure we only replace zeroes in back of nums 1
			nums1[index] = nums2[count] // merge num2 element into nums 1
			count++
		}
	}
	sort.Ints(nums1) // increasing order
	return nums1
}

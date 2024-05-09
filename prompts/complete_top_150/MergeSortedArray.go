package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	count := 0
	for index, v := range nums1 {
		if v == 0 && index >= m {
			nums1[index] = nums2[count]
			count++
		}
	}
	sort.Ints(nums1)
	return nums1
}

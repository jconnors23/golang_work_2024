package main

import (
	"math"
	"sort"
)

// func main() {
// 	nums := []int{1, 1, 1, 2, 2, 2, 3, 5, 6, 7}
// 	fmt.Println(removeDuplicates2(nums))
// }

func removeDuplicates2(nums []int) int {
	filler := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == curr {
				count++
				if count >= 3 {
					nums[j] = filler
					count--
				}
			}
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return numUnique(nums, filler)
}

func numUnique(nums []int, filler int) int {
	count := 0
	for _, v := range nums {
		if v != filler {
			count++
		}
	}
	return count
}

/*
iterate through nums track current element count
	when reaching >=3 index of curr
		 	fill -1 for index
			if no non unique available, fill -1

	after duplicates handled, resort array in non decreasing order via sort . slices package

move each of these to back of arr by swapping with first available non unique element
*/

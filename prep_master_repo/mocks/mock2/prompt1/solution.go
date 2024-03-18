package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates2(nums))
}

func removeDuplicates2(nums []int) int {
	new := []int{}
	for i := 0; i < len(nums); i++ {
		ele := nums[i]
		if !contains(new, ele) {
			new = append(new, ele)
		}
	}
	nums = new
	fmt.Printf("%d\n", nums)
	return len(nums)
}

func contains(new []int, target int) bool {
	for _, v := range new {
		if v == target {
			return true
		}
	}
	return false
}

//func removeDuplicates(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		ele := nums[i]
// 		for index, v := range nums {
// 			if index == i { // ignore value we are currently at in arr
// 				continue
// 			}
// 			// fmt.Printf("V: %d\n", v)
// 			// fmt.Printf("Element: %d\n", ele)
// 			if v == ele {
// 				if index >= len(nums) { // if at last element
// 					nums = nums[:len(nums)-1]
// 					fmt.Printf("After Last element: %d\n", nums)
// 					continue
// 				}
// 				fmt.Printf("V: %d\n", v)
// 				nums = append(nums[:index], nums[index+1:]...) // remove duplicate element
// 				fmt.Printf("After Normal remove: %d\n", nums)
// 			}
// 		}
// 	}
// 	fmt.Printf("%d\n", nums)
// 	return len(nums)
// }

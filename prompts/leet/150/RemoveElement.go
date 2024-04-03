package main

import "fmt"

func main() {
	nums := []int{3, 3}
	fmt.Println(removeElement(nums, 2))
}

func removeElement(nums []int, val int) int {
	track := findFirst(val, nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != val && i >= track { // if num is further back in nums compared to val
			// swap elements in place
			fmt.Printf("before swap %d, %v \n", nums, track)
			nums[track] = nums[i]
			nums[i] = val
			// reset index of val
			track = findFirst(val, nums)
			fmt.Printf("after swap %d, %v \n", nums, track)
		}
	}
	return placeBlanks(val, nums)
}

func findFirst(val int, nums []int) int {
	var index = -1
	for i, v := range nums {
		if v == val {
			index = i
			break
		}
	}
	if index >= 0 {
		return index
	} else {
		return len(nums)
	}
}

func placeBlanks(val int, nums []int) int {
	var count int
	for _, v := range nums {
		if v == val {
			v = -1
		} else {
			count++
		}
	}
	return count
}

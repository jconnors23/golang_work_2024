/*

Given an array a that contains only numbers in the range from 1 to a.length,
find the first duplicate number for which the second occurrence has the minimal index.
 In other words, if there are more than 1 duplicated numbers,
 return the number for which the second occurrence has a smaller index than the
 second occurrence of the other number does. If there are no such elements, return -1.

For a = [2, 1, 3, 5, 3, 2], the output should be solution(a) = 3.

// init duplicate two dup objects, counter var
// when dup found, add index of 2nd to first dup
// if & when second dup found, add index of 2nd to 2nd dup
// compare and return, if no dupes return -1

// edge - >2 dupes, edge - > 3 of a kind or more

 comment test
*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 1, 3, 5, 3, 2}
	fmt.Println(solution(arr))
}

func solution(arr []int) int {
	var count int
	var dupi1, dupi2 []int
	// find duplicates (bf)
	for i := 0; i < len(arr); i++ {
		current := arr[i]
		for j := 0; j < len(arr); j++ {
			if arr[j] == current {
				if count == 0 {
					dupi1 = append(dupi1, current, j) // add duplicate number and its index
					count++
					continue
				} else {
					dupi2 = (append(dupi2, current, j))
					count++
					continue
				}
			}
		}
	}
	if dupi1[1] < dupi2[1] {
		return dupi1[1]
	} else {
		return dupi2[1]
	}
}

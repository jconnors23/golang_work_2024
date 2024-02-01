/*

Given an array a that contains only numbers in the range from 1 to a.length,
find the first duplicate number for which the second occurrence has the minimal index.
 In other words, if there are more than 1 duplicated numbers,
 return the number for which the second occurrence has a smaller index than the
 second occurrence of the other number does. If there are no such elements, return -1.

For a = [2, 1, 3, 5, 3, 2], the output should be solution(a) = 3.

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

}

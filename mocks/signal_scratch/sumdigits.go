/*

You are given a two-digit integer n.
Return the sum of its digits.

Example

For n = 29, the output should be
solution(n) = 11.

*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 29
	fmt.Println(solution(num))
}
func solution(num int) int {
	sum := 0
	digits := strconv.Itoa(num)
	for i := 0; i < len(digits); i++ {
		count, err := strconv.Atoi(string(digits[i]))
		if err != nil {
			_ = err
		}
		sum += count
	}
	return sum
}

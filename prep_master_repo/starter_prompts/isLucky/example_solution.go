package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 1230
	fmt.Println(solution(num))
}

func solution(n int) bool {
	str_num := strconv.Itoa(n)  // convert integer to string that we can iterate over
	half := len(str_num) / 2    // constraint that the input has even number of digits
	str_num_2 := str_num[half:] // slice second half of string
	sum1, sum2 := 0, 0
	for i := 0; i < len(str_num)/2; i++ {
		sum1 += int(str_num[i])   // convert digit to int, increment sum first half
		sum2 += int(str_num_2[i]) // sum second half
	}
	return sum1 == sum2 // if sums are equal return true else return false
}

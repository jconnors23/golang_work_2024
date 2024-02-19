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
	str_num := strconv.Itoa(n)
	half := len(str_num) / 2
	str_num_2 := str_num[half:]
	sum1, sum2 := 0, 0
	for i := 0; i < len(str_num)/2; i++ {
		sum1 += int(str_num[i])
		sum2 += int(str_num_2[i])
	}
	return sum1 == sum2
}

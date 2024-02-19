package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution("aabaa"))
}

func solution(inputString string) bool {
	// check for equality on all chars except middle
	length := len(inputString)
	if len(inputString) == 1 {
		return true
	}
	count := length
	for i := 0; i < len(inputString); i++ {
		if len(inputString)%2 == 1 { // if odd num chars
			if i == int((float64(length / 2))) { // if at middle char ignore in checking
				return true // can assume true once at middle char
			}
		}

		if inputString[i] != inputString[count-1] {
			return false
		}
		count--
	}
	return true
}

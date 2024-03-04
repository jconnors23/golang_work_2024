package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution("aabaa"))
}

func solution(inputString string) bool {
	// check for equality on all characters except middle, if all except middle are equivalent input is a palindrome
	length := len(inputString)
	if len(inputString) == 1 {
		return true
	}
	point2 := length                                       // use count variable to have pointer at end of word
	for point1 := 0; point1 < len(inputString); point1++ { // use 'i' to have pointer at beginning of word
		if len(inputString)%2 == 1 { // if odd num chars
			if point1 == int((float64(length / 2))) { // if at middle char ignore in checking
				return true // can assume true once at middle char
			}
		}

		if inputString[point1] != inputString[point2-1] { // if characters are not the same, not a palindrome
			return false
		}
		point2-- // move our pointer inward for further comparisons
	}
	return true
}

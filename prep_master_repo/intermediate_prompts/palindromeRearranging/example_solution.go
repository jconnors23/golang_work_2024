package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution("aabb"))
}

func solution(a string) bool {
	numChars := len(a)
	hasOddLength := numChars%2 == 1 // boolean representing if word has odd number of characters
	m := make(map[string]int)       // mapping of each char in string
	for i := 0; i < len(a); i++ {   // iterate through word
		currentChar := string(a[i])        // current character in word
		_, exist := m[string(currentChar)] // if it exists in map, it has already been counted
		if exist {
			continue
		}
		count := strings.Count(a, string(a[i])) // count number of char within word
		m[string(currentChar)] = count          // set value of key character to its count in the word
	}
	// check values for any odd, if isOdd == false, return false
	numCharOdd := 0
	for k := range m { // iterate through the maps values, counting for num odd chars
		if m[k]%2 == 1 { // if odd count for character in word
			numCharOdd += 1
			if !hasOddLength { // if word length is even, but there is a odd number of chars, it can not be rearranged to form a palindrome
				return false
			}
			if numCharOdd > 1 { // if num of odd chars within word is greater than 1, cant be rearranged
				return false
			}
		}
	}
	return true
}

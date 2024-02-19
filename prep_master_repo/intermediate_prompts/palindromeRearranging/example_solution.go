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
	isOdd := numChars%2 == 1
	m := make(map[string]int)
	// mapping of each char in string
	for i := 0; i < len(a); i++ {
		currentChar := string(a[i])
		_, exist := m[string(currentChar)]
		if exist {
			continue
		}
		count := strings.Count(a, string(a[i]))
		m[string(currentChar)] = count
	}
	// check values for any odd, if isOdd == false, return false
	numOdd := 0
	for k := range m {
		if m[k]%2 == 1 {
			numOdd += 1
			if !isOdd {
				return false
			}
			if numOdd > 1 {
				return false
			}
		}
	}
	return true
}

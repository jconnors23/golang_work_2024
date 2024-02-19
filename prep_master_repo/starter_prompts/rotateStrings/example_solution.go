package main

import "fmt"

func main() {
	str := "abcde"
	goal := "cdeab"
	fmt.Println(rotateString(str, goal))
}
func rotateString(s, goal string) bool {
	lives := len(s) * 2 // assume after 2x shift trials that rotation possibility is false
	for lives > 0 {
		trial := shift(s)
		if trial == goal {
			return true
		}
		s = trial
		lives--
	}
	return false
}

func shift(s string) string {
	after := s[1:]
	return after + string(s[0])
}

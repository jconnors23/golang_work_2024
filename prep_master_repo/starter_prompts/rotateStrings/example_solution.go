package main

import "fmt"

func main() {
	str := "abcde"
	goal := "cdeab"
	fmt.Println(rotateString(str, goal))
}
func rotateString(s, goal string) bool {
	trials := len(s) * 2 // assume after 2 times len of input shift trials that rotation possibility is false
	for trials > 0 {     // while alive
		trial := shift(s)
		if trial == goal {
			return true
		}
		s = trial // have to set input to modified string, keep shifting until out of trials
		trials--
	}
	return false
}

func shift(s string) string { // shift helper function
	after := s[1:]              // take everything except first char, slices in Go are inclusive of first index and exclusive of end index
	return after + string(s[0]) // concatenate end of our slice with the first char of input to complete the shift
}

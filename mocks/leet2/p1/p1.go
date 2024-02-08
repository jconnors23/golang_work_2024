/*

Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.

A shift on s consists of moving the leftmost character of s to the rightmost position.

For example, if s = "abcde", then it will be "bcdea" after one shift.


Example 1:

Input: s = "abcde", goal = "cdeab"
Output: true
Example 2:

Input: s = "abcde", goal = "abced"
Output: false

*/

package main

import "fmt"

func main() {
	fmt.Println(shiftGoal("abcde", "abced"))
}

func shiftGoal(s, goal string) bool {
	lives := len(s) * 2
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

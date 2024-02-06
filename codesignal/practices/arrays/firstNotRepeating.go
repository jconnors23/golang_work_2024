/*

Given a string s consisting of small English letters,
find and return the first instance of a non-repeating character in it.
If there is no such character, return '_'.

Example

For s = "abacabad", the output should be
solution(s) = 'c'.

There are 2 non-repeating characters in the string: 'c' and 'd'.
Return c since it appears in the string first.

For s = "abacabaabacaba", the output should be
solution(s) = '_'.

There are no characters in this string that do not repeat.

*/

package arrays

func FirstNotRepeating(str string) string {
	for i := 0; i < len(str); i++ {
		current := str[i]
		hasRepeat := false
		for j := 0; j < len(str); j++ {
			if str[j] == current && j != i {
				hasRepeat = true
			}
		}
		if hasRepeat {
			return string(current)
		}
	}
	return "_"
}

/*

package main

import t"
	"practices/arrays"
)(
	"fm

func main() {
	arr := []int{2, 1, 3, 5, 3, 2}
	fmt.Println(arrays.FindDuplicates(arr))
	fmt.Println(arrays.FirstNotRepeating("Strr"))
}


*/

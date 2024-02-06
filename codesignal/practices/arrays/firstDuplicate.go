/*

Given an array a that contains only numbers in the range from 1 to a.length,
find the first duplicate number for which the second occurrence has the minimal index.
 In other words, if there are more than 1 duplicated numbers,
 return the number for which the second occurrence has a smaller index than the
 second occurrence of the other number does. If there are no such elements, return -1.

For a = [2, 1, 3, 5, 3, 2], the output should be solution(a) = 3.

// init duplicate two dup objects, counter var
// when dup found, add index of 2nd to first dup
// if & when second dup found, add index of 2nd to 2nd dup
// compare and return, if no dupes return -1

// edge - >2 dupes, edge - > 3 of a kind or more

// unpassed case:
 [8, 4, 6, 2, 6, 4, 7, 9, 5, 8]

 e 6 r 4


*/

package arrays

func FindDuplicates(arr []int) int {
	var count int
	var dup1, dup2 []int
	for i := 0; i < len(arr); i++ {
		current := arr[i]
		for j := 0; j < len(arr); j++ {
			if arr[j] == current && j != i {
				// if first dup
				if count == 0 {
					dup1 = append(dup1, current, i, j) // add duplicate number and its 1st, 2nd index
					count++
				} else {
					dup2 = (append(dup2, current, i, j))
					count++
				}
			}
		}
	}
	if count == 0 {
		return -1 // if no dupes
	} else if count == 1 {
		return dup1[1] // return first index of only duplicate
	} else { // return (the number) smallest index of second occurence among dupes
		if dup1[2] < dup2[2] {
			return dup1[0]
		} else {
			return dup2[0]
		}
	}
}

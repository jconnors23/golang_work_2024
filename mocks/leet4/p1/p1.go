package main

import "fmt"

// potential solution:

// iterate through pos 0 & 1 of rec1, compare to pos 0,1 then pos 2, 3
// check: if both 0 & 1 are strictly less than both 0,1 or both 2,3, then check for strictly greater than
// strictly less and strictly greater need to hold true to return true, else return false
// rec1 = [0,0,2,2], rec2 = [1,1,3,3]
// need pair where one is less than and one is greater than
// abs val of distance calculations for positive area ??
func main() {
	// rec1 := []int{0, 0, 2, 2}
	// rec2 := []int{1, 1, 3, 3}
	// rec1 := []int{7, 8, 13, 15}
	// rec2 := []int{10, 8, 12, 20}
	rec1 := []int{0, 0, 1, 1}
	rec2 := []int{1, 0, 2, 1}

	fmt.Println(isRectangleOverlap(rec1, rec2))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	coord1 := []int{rec1[0], rec1[1]}
	coord2 := []int{rec1[2], rec1[3]}
	coord3 := []int{rec2[0], rec2[1]}
	coord4 := []int{rec2[2], rec2[3]}

	if coord3[0] < coord1[0] && coord3[1] < coord1[1] {
		if coord4[0] > coord1[0] && coord4[1] > coord1[1] {
			return true
		}
	}
	if coord4[0] < coord1[0] && coord4[1] < coord1[1] {
		if coord3[0] > coord1[0] && coord3[1] > coord1[1] {
			return true
		}
	}
	if coord3[0] < coord2[0] && coord3[1] < coord2[1] {
		if coord4[0] > coord2[0] && coord4[1] > coord2[1] {
			return true
		}
	}
	if coord4[0] < coord2[0] && coord4[1] < coord2[1] {
		if coord3[0] > coord2[0] && coord3[1] > coord2[1] {
			return true
		}
	}
	return false
}

// if rec1[0] < rec2[0] && rec1[1] < rec2[1] { // slt
// 	if rec1[2] > rec2[2] && rec1[3] > rec2[3] { // sgt
// 		return true
// 	}
// }
// if rec1[0] < rec2[2] && rec1[1] < rec2[3] { // slt
// 	if rec1[2] > rec2[0] && rec1[3] > rec2[1] { // sgt
// 		return true
// 	}
// }
// if rec1[0] < rec2[0] && rec1[1] < rec2[1] { // slt
// 	if rec1[2] > rec2[2] && rec1[3] > rec2[3] { // sgt
// 		return true
// 	}
// }

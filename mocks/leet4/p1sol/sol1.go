package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	x1, y1, x2, y2 := rec1[0], rec1[1], rec1[2], rec1[3]
	xx1, yy1, xx2, yy2 := rec2[0], rec2[1], rec2[2], rec2[3]

	if xx2 <= x1 || xx1 >= x2 {
		return false
	}
	if yy2 <= y1 || yy1 >= y2 {
		return false
	}
	return true
}

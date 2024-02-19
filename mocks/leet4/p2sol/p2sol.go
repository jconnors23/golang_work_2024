package main

func partitionLabels(s string) []int {
	result := []int{}

	// first, let's count the index that each character was last seen
	lastSeen := make([]int, 26)
	for index, char := range s {
		lastSeen[char-'a'] = index
	}

	// to keep track of the starts and ends of our partition
	start, end := 0, 0

	for index, char := range s {
		// update the end to be the max of current end and last seen of this char
		// because if it was seen later, we wont be able to partition here
		if lastSeen[char-'a'] > end {
			end = lastSeen[char-'a']
		}

		// check if current index is the end (we found a partition!)
		if index == end {
			result = append(result, end-start+1)
			start = index + 1
		}
	}

	return result
}
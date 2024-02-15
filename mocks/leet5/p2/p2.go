package main

import (
	"fmt"
	"strings"
)

// group 1 & group 2
// compare in bw \n\t if diff then ..
// if group 2 != "file.ext"

func main() {
	str := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	fmt.Println(lengthLongestPath(str))
}

func lengthLongestPath(input string) int {
	if !(strings.Contains(input, "file.ext")) {
		return 0
	}
	given := "dir" // "/"

	count := 4
	for i := 0; i < len(input); i++ {
		for input[i] != '/' {

		}
		if input[i] == "/" {

		}
	}

	return count
}

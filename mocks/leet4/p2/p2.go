package main

import (
	"fmt"
	"strings"
)

func main() {
	// " ababcbaca defegde hijhklij "
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}

func partitionLabels(s string) []int {
	sizes := []int{}
	partitions := []string{}
	for i := 0; i < len(s); i++ {
		last := lastIndex(s, s[i])
		if last != -1 {
			position := last + 1
			part := s[i:position]
			for _, char := range part {
				if lastIndex(s, byte(char)) >= position {
					position := lastIndex(s, byte(char)) + 1
					part = s[i:position]
					fmt.Printf("newpart: %s, curr char: %s\n", part, string(char))

				}
			}
			partitions = append(partitions, part)
			i = last
		}

	}
	fmt.Println(partitions)
	for _, v := range partitions {
		sizes = append(sizes, len(v))
	}
	return sizes
}

func lastIndex(s string, c byte) int {
	return strings.LastIndex(s, string(c))
}

// for j := len(s) - 1; j > 0; j-- {
// 	if s[j] == s[i] {
// 		part := s[i : j+1]
// 		partitions = append(partitions, s[i:j+1])
// 		i = j + 1
// 	}
// }

package main

import "fmt"

func main() {
	input := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(input))
}

/*
init return arr
iterate through input
for each index of input, make map of chars ?
for each key of each map, check if exact val exists in other mappings ?
if so, push char to return arr
*/

type PositionMaps struct {
	Map map[string]int
}

func commonChars(words []string) []string {
	commons := []string{}
	// positions := PositionMaps
	str := words[0] // original map
	m := make(map[string]int)
	for _, char := range str {
		m[string(char)] += 1
	}
	for _, word := range words {
		if word == str {
			continue
		}
		comparison := make(map[string]int) // mapping of word in question
		for _, char := range word {
			comparison[string(char)] += 1
		}
		// compare
		for k, v := range m {
			if m[k] == comparison[k] {
				if v > 1 {
					count := v
					for count > 0 {
						if countOf(commons, k) < v {
							commons = append(commons, k)
						}
						count--
					}
				} else {
					if !contains(commons, string(k)) {
						commons = append(commons, k)
					}
				}
			} else {
				remove(commons, k)
			}

		}
	}
	return commons
}

func remove(commons []string, k string) []string {
	for i, v := range commons {
		if v == k {
			return append(commons[:i], commons[i+1:]...)
		}
	}
	return commons
}

func contains(arr []string, k string) bool {
	for _, v := range arr {
		if v == k {
			return true
		}
	}
	return false
}

func countOf(arr []string, k string) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == k {
			count++
		}
	}
	return count
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	k := 4
	str := "5F3Z-2e-9-w"
	fmt.Println(licenseKeyFormatting(str, k))
}

func licenseKeyFormatting(s string, k int) string {
	dash := "-"
	// strings method replace will take away all dashes from our input, -1 means replace all occurences
	wo_dashes := strings.Replace(s, string(dash), "", -1) // 2-5G-3J - > 25G3J
	// handling for an edge case
	if len(wo_dashes) == 1 {
		str := []string{wo_dashes}
		return final(str)
	}
	var formatted []string          // string slice representing formatted license key
	len_first := len(wo_dashes) % k // this int represents the length of  first group
	var tracker int                 // this int represents length of each subsequent group
	if len_first > 0 {              // if there is a remainder, first group will be length of that remainder
		for i := 0; i < len_first; i++ {
			formatted = append(formatted, string(wo_dashes[i]))
			tracker++
		}
		formatted = append(formatted, dash)
	}
	for tracker < len(wo_dashes) {
		for i := 0; i < k; i++ { // while less than amount k, append to group
			formatted = append(formatted, string(wo_dashes[tracker]))
			tracker++
		}
		if tracker == len(wo_dashes) {
			return final(formatted)
		}
		// append formatted license key with a dash in between groups of length k
		formatted = append(formatted, dash)
	}
	return final(formatted)
}

// final formatting helper method
func final(s []string) string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i]) // requirement to uppercase chars in formatting
	}
	return strings.Join(s, "") // take slice join into string for return
}

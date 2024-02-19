package main

import (
	"fmt"
	"strings"
)

/*
 1. iterate through s, remove all dashes, then get len(s) for count of chars
 2. divide count by k, if remainder, first group = remainder(s) for new string
 3. for each next group of k chars, push to new string, then add a dash
 4. return result


*/

func main() {
	k := 4
	str := "5F3Z-2e-9-w"
	fmt.Println(licenseKeyFormatting(str, k))
}

func licenseKeyFormatting(s string, k int) string {
	dash := "-"
	wo_dashes := strings.Replace(s, string(dash), "", -1) // 2-5G-3J - > 25G3J
	if len(wo_dashes) == 1 {
		str := []string{wo_dashes}
		return final(str)
	}
	//fmt.Printf("string w/o dashes: %s\n", wo_dashes)
	var tracker int
	var formatted []string
	//group_count := math.Floor(float64(len(wo_dashes) / k))
	len_first := len(wo_dashes) % k
	//fmt.Println("remainder: ", len_first)
	if len_first > 0 {
		for i := 0; i < len_first; i++ {
			formatted = append(formatted, string(wo_dashes[i]))
			tracker++
		}
		formatted = append(formatted, dash)
		//fmt.Printf("first group %s\n", formatted)
	}
	for tracker < len(wo_dashes) {
		for i := 0; i < k; i++ {
			//fmt.Println("tracker count:", tracker)
			formatted = append(formatted, string(wo_dashes[tracker]))
			//fmt.Printf("formatted %s\n", formatted)
			tracker++
		}
		//fmt.Println("tracker count:", tracker)
		if tracker == len(wo_dashes) {
			return final(formatted)
		}
		formatted = append(formatted, dash)
	}
	return final(formatted)
}

func final(s []string) string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i])
	}
	return strings.Join(s, "")
}

package main

import (
	"slices"
)

func lengthOfLongestSubstring(s string) int {
	globalMax := 0
	currentMax := 0
	var subs []string

	for _, value := range s {
		strValue := string(value)
		valueIndex := slices.Index(subs, strValue)

		if valueIndex == -1 { // value not in substring
			subs = append(subs, strValue)
			currentMax = currentMax + 1
			if currentMax > globalMax {
				globalMax = currentMax
			}
		} else {
			// delete up to the existing occurence of value
			subs = subs[valueIndex+1:]
			// add the new item to the updated substring
			subs = append(subs, strValue)
			// update the current max
			currentMax = len(subs)
		}

		/* fmt.Println("Current Character: ", strValue)
		fmt.Println("Current Substring: ", subs)
		println("Current Max: ", currentMax)
		println("Global Max: ", globalMax)
		*/
	}
	return globalMax
}

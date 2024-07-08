package main

import (
	"fmt"
	"os"
)

func split(str, sep string) []string {
	// Create an empty slice to store the split strings
	var result []string

	// Set the starting index of the current substring to 0
	start := 0

	for i := 0; i < len(str); i++ {
		if i+len(sep) <= len(str) && str[i:i+len(sep)] == sep {			
			// Update the start index to the end of the separator
			result = append(result, str[start:i])
			start = i + len(sep)
			// We decrement by 1 since i is 0 index based
			i += len(sep) - 1			
		}
	}
	// After the loop, append the remaining substring from the start index to end of the string to
	// the result slice	
	result = append(result, str[start:])
	return result
}

// This commented out split function shows how to use make without append in slices.
// func split(str, sep string) []string {

// We start with the assumption that there is at least one segment with or without sep.
// 	estimatedSplits := 1
// 	for i := 0; i < len(str); i++ {		
// 		if i+len(sep) <= len(str) && str[i:i+len(sep)] == sep {
// 			estimatedSplits++
// 			// Increment the position by the length of the separator minus 1
// 			// to avoid checking the same separator again
// 			i += len(sep) - 1
// 		}
// 	}


// 	// Create a slice with the estimated capacity
// 	result := make([]string, estimatedSplits)
// 	start, index := 0, 0

// 	// Iterate through the string again to perform the actual splitting
// 	for i := 0; i < len(str); i++ {
// 		if i+len(sep) <= len(str) && str[i:i+len(sep)] == sep {
// 			// Assign the substring from the start position to the current position to the result slice
// 			result[index] = str[start:i]
// 			index++
// 			// Update the start position to the position after the separator
// 			start = i + len(sep)
// 			// Increment the position by the length of the separator minus 1
// 			// to avoid checking the same separator again
// 			i += len(sep) - 1
// 		}
// 	}

// 	// Assign the remaining substring (from the last separator to the end of the string) to the result slice
// 	result[index] = str[start:]
// 	// Return the result slice, trimming any unused capacity
// 	return result
// }


func main() {
	if len(os.Args) != 1{
		return
	}
	str := "HelloHAhowHAareHAyou"
	

	fmt.Printf("%#v\n", split(str, "HA"))
}

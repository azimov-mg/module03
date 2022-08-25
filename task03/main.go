package main

import (
	"fmt"
	"sort"
)

func main() {
	SliceString := []string{"lemon", "banana", "apple", "cherry"}
	count := 0

	sort.Slice(SliceString, func(i, j int) bool {
		return SliceString[i] < SliceString[j]
	})

	for key, value := range SliceString {
		if value[key] < value[key+1] {
			count++
			// fmt.Println(count)
		}
	}

	if count == 3 {
		fmt.Println(SliceString, true)
	} else {
		fmt.Println(SliceString, false)
	}
}

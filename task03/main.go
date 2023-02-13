package main

import "fmt"

func isSorted(words []string) bool {
	for i := 1; i < len(words); i++ {
		if words[i-1] > words[i] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"apple", "banana", "cherry"}
	fmt.Println(isSorted(words))
}

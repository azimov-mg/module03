package main

import "fmt"

func stringToMap(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	return counts
}

func main() {
	s := "съешь ещё этих мягких французских булок, да выпей чаю"
	counts := stringToMap(s)
	for char, count := range counts {
		fmt.Printf("'%c': %d\n", char, count)
	}
}

package main

import "fmt"

func main() {
	nums := []int{2, 6, 8, 2, 4, 9, 1, 1, 1, 7}
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}

	if count >= 2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

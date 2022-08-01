package main

import "fmt"

func SameArr(nums []int) bool {
	count := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
				if count >= 2 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 0, 6, 7, 4, 7, 0}
	fmt.Println(nums, SameArr(nums))
}

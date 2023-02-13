package main

import "fmt"

func main() {
	Arr := [...]int{1, 2, 3, 2, 5, 1}

	for i := 0; i < len(Arr); i++ {
		count := 0
		for j := 0; j < len(Arr); j++ {
			if Arr[i] == Arr[j] {
				count++
			}
		}
		if count > 1 {
			fmt.Println(Arr[i], "is true")
		} else {
			fmt.Println(Arr[i], "is false")
		}
	}
}

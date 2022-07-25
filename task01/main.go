package main

import "fmt"

type HealthCheck struct {
	ServiceID int
	status    string
}

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

func main() {
	GCheck := GenerateCheck()
	fmt.Println("ID will be print here")

	for _, val := range GCheck {
		if val.status == PassStatus {
			fmt.Println(val.ServiceID)
		}
	}
}

func GenerateCheck() []HealthCheck {
	GenerateSlice := make([]HealthCheck, 0, 5)
	for i := 0; i <= 5; i++ {
		if i%2 == 0 && i != 0 {
			GenerateSlice = append(GenerateSlice, HealthCheck{i, PassStatus})
		} else {
			GenerateSlice = append(GenerateSlice, HealthCheck{i, FailStatus})
		}
	}
	return GenerateSlice
}

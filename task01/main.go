package main

import (
	"fmt"
)

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type HealthCheck struct {
	ServiceID int
	status    string
}

func GenerateCheck() []*HealthCheck {

	var checks []*HealthCheck
	for i := 0; i <= 5; i++ {
		s := &HealthCheck{ServiceID: i}
		if s.ServiceID%2 == 0 && s.ServiceID != 0 {
			s.status = PassStatus
		} else {
			s.status = FailStatus
		}
		checks = append(checks, s)
	}
	return checks
}

func main() {
	fmt.Println("The identification and its status will be shown:")

	gChecks := GenerateCheck()

	for _, check := range gChecks {
		if check.status == PassStatus {
			fmt.Printf("Service ID : %d, Status: %s\n", check.ServiceID, check.status)
		}
	}
}

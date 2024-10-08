package hw02

import (
	"fmt"

	"github.com/mar4ehk0/go/hw06_testing/internal/hw02/employee"
)

func PrintStaff(employees []employee.Employee) {
	for i := 0; i < len(employees); i++ {
		fmt.Println(employees[i])
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func do() {
	// Create a slice to store Employee objects
	employees := make([]Employee, 0)

	// Generate a large number of Employee objects
	for i := 0; i < 1000000; i++ {
		employee := Employee{
			Name:   "John Doe",
			Age:    30,
			Salary: 50000.0,
		}
		time.Sleep(100 * time.Millisecond)
		employees = append(employees, employee)
	}

	// Print the names of all employees with salaries above 45000
	printHighSalaryEmployees(employees)
}

func printHighSalaryEmployees(employees []Employee) {
	highSalaryEmployees := make([]string, 0)

	// Iterate over all employees and filter those with salaries above 45000
	for _, employee := range employees {
		if employee.Salary > 45000.0 {
			highSalaryEmployees = append(highSalaryEmployees, employee.Name)
		}
	}

	// Print the names of high-salary employees
	fmt.Println("High Salary Employees: " + strings.Join(highSalaryEmployees, ", "))
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	do()
}

package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee := Employee{LastName: "Doe", FirstName: "Johnson"}
	employee.ID = 1001
	employee.Address = "AAA"
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "AAA"
	fmt.Println(employee)
}

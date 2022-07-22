package main

import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {
	employee := Employee{
		Person: Person{
			FirstName: "aaa",
			LastName:  "bbb",
		},
		ManagerID: 1001,
	}
	employee.Address = "ccc"
	fmt.Println(employee)

	fmt.Println(employee.FirstName)
}

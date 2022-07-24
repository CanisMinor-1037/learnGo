package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrNotFound = errors.New("Employee not found!")

func main() {
	employee, err := getInfomation(1001)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT FOUND: %v\n", err)
	} else {
		fmt.Print(employee)
	}
}

func getInfomation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

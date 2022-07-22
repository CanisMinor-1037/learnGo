package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["bob"] = 31
	studentsAge["john"] = 32
	age, exist := studentsAge["alice"]
	if exist {
		fmt.Println("Alice's age is", age)
	} else {
		fmt.Println("Not found")
	}
}

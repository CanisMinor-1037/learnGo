package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println("Before", studentsAge)
	delete(studentsAge, "bob")
	fmt.Println("After", studentsAge)
}

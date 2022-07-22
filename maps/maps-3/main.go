package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 31
	studentsAge["bob"] = 32
	fmt.Println("Bob's age is", studentsAge["bob"])
}

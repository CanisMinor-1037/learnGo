package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	for name, age := range studentsAge {
		fmt.Printf("Name:%s\tAge:%d\n", name, age)
	}
}

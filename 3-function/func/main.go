package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := 1
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("sum: ", sum)
}

func sum(str1 string, str2 string) int {
	number1, _ := strconv.Atoi(str1)
	number2, _ := strconv.Atoi(str2)
	return number1 + number2
}

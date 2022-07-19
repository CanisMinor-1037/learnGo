package main

import "fmt"

func main() {
	var array [3]int
	array[1] = 10
	fmt.Println("array[0] =", array[0])
	fmt.Println("array[1] =", array[1])
	fmt.Println("array[len(array) - 1] =", array[len(array)-1])
}

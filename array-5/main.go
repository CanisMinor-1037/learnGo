package main

import "fmt"

func main() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println(twoD)
}

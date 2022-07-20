package main

import "fmt"

func main() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		fmt.Println("You entered:", val)
		if val < 0 {
			panic("panic")
		} else if val == 0 {
			fmt.Println("0 is neither negative nor positive")
		}
	}
}

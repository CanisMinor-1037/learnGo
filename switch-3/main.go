package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Firday":
		fmt.Println("It's time to learn Some Go")
	default:
		fmt.Println("It's weekend, time to reset!")
	}

	fmt.Println(time.Now().Weekday().String())
}

package main

import (
	"fmt"
)

func main() {
	var (
		a, b int
	)
	fmt.Scan(&a, &b)

	checkNum := a * b
	if checkNum%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

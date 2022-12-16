package main

import (
	"fmt"
)

func main() {
	var num string
	var count int
	fmt.Scan(&num)
	for _, c := range num {
		if string(c) == "1" {
			count++
		}
	}
	fmt.Println(count)
}

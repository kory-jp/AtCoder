package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < 100; i++ {
		fmt.Println(N)
		N = N - 1
		if N == 0 {
			break
		}
	}
	fmt.Println(0)
}

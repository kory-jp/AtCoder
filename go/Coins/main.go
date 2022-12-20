package main

import "fmt"

func main() {
	var (
		a                int
		b                int
		c                int
		x                int
		numberOfPatterns int
	)
	const A, B, C int = 500, 100, 50

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	if !checkNum(a, b, c, x) {
		return
	}
	compareWithOneTypeOfCoin(a, A, x, &numberOfPatterns)
	compareWithOneTypeOfCoin(b, B, x, &numberOfPatterns)
	compareWithOneTypeOfCoin(c, C, x, &numberOfPatterns)
	compareWithTwoTypeOfCoin(a, A, b, B, x, &numberOfPatterns)
	compareWithTwoTypeOfCoin(a, A, c, C, x, &numberOfPatterns)
	compareWithTwoTypeOfCoin(c, C, b, B, x, &numberOfPatterns)
	compareWithThreeTypeOfCoin(a, A, b, B, c, C, x, &numberOfPatterns)

	fmt.Println(numberOfPatterns)
}

func checkNum(a, b, c, x int) bool {
	return 0 <= a && a <= 50 && 0 <= b && b <= 50 && 0 <= c && c <= 50 && 1 <= a+b+c && 1 <= x && x <= 20000 && x%50 == 0
}

func compareWithOneTypeOfCoin(numCoin int, coinPrice int, x int, numberOfPatterns *int) {
	for i := 0; i < numCoin+1; i++ {
		totalAmount := coinPrice * i
		if totalAmount == x {
			*numberOfPatterns = *numberOfPatterns + 1
		}
	}
}

func compareWithTwoTypeOfCoin(numCoinA int, coinPriceA int, numCoinB int, coinPriceB int, x int, numberOfPatterns *int) {
	for a := 1; a < numCoinA+1; a++ {
		for b := 1; b < numCoinB+1; b++ {
			totalAmount := coinPriceA*a + coinPriceB*b
			if totalAmount == x {
				*numberOfPatterns = *numberOfPatterns + 1
			}
		}
	}
}

func compareWithThreeTypeOfCoin(numCoinA int, coinPriceA int, numCoinB int, coinPriceB, numCoinC int, coinPriceC, x int, numberOfPatterns *int) {
	for a := 1; a < numCoinA+1; a++ {
		for b := 1; b < numCoinB+1; b++ {
			for c := 1; c < numCoinC+1; c++ {
				totalAmount := coinPriceA*a + coinPriceB*b + coinPriceC*c
				if totalAmount == x {
					*numberOfPatterns = *numberOfPatterns + 1
				}
			}
		}
	}
}

package main

import "fmt"

func main() {
	var (
		N             int
		A             int
		B             int
		analyedValues []int
		checkedValues []int
	)

	fmt.Scanf("%d %d %d", &N, &A, &B)

	if !validation(N, A, B) {
		return
	}

	for i := 1; i <= N; i++ {
		list := digit(i, analyedValues)
		sum := sumNum(list)
		if checkNum(sum, A, B) {
			checkedValues = append(checkedValues, i)
		}
	}
	fmt.Println(sumNum(checkedValues))
}

func validation(N, A, B int) bool {
	return 1 <= N && N <= 100000 && 1 <= A && A <= 36 && 1 <= B && B <= 36 && A <= B
}

func digit(i int, list []int) []int {
	if i > 0 {
		return digit(i/10, append(list, i%10))
	}
	return list
}

func sumNum(list []int) int {
	var sum int
	for _, v := range list {
		sum = sum + v
	}
	return sum
}

func checkNum(sum, A, B int) bool {
	return A <= sum && sum <= B
}

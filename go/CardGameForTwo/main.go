package main

import (
	"fmt"
)

func main() {
	var (
		N int
		A int
		B int
	)

	fmt.Scanf("%d", &N)
	numArr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &numArr[i])
	}

	if !validation(N, numArr) {
		return
	}

	for i := 0; i <= N; i++ {
		if i%2 == 0 {
			idx, biggestNum := getBiggestNumber(numArr)
			A = A + biggestNum
			numArr = remakeArray(idx, numArr)
		} else {
			idx, biggestNum := getBiggestNumber(numArr)
			B = B + biggestNum
			numArr = remakeArray(idx, numArr)
		}
	}
	fmt.Println(A - B)
}

func validation(N int, numArr []int) bool {
	if N < 1 && 100 < N {
		return false
	}
	for i := range numArr {
		if numArr[i] < 1 || 100 < numArr[i] {
			return false
		}
	}
	return true
}

func getBiggestNumber(numArr []int) (int, int) {
	checkIdx := 0
	checkNum := 0
	for i, v := range numArr {
		if checkNum < v {
			checkIdx = i
			checkNum = v
		}
	}
	return checkIdx, checkNum
}

func remakeArray(num int, oldArr []int) []int {
	var newArr []int
	for i := range oldArr {
		if num != i {
			newArr = append(newArr, oldArr[i])
		}
	}
	return newArr
}

package main

import "fmt"

func main() {
	var N int

	fmt.Scanf("%d", &N)
	numArr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &numArr[i])
	}

	unique := SliceUnique(numArr)

	fmt.Println(len(unique))
}

func SliceUnique(target []int) (unique []int) {
	m := map[int]bool{}
	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

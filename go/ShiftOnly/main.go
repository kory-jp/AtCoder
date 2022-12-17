package main

import (
	"fmt"
)

func main() {
	// 受け取るデータ数
	var dataCount int
	// 受け取ったデータをまとめる配列
	// var numArr []int
	// 読み込み関数
	fmt.Scanf("%d", &dataCount)

	numArr := make([]int, dataCount)
	for i := 0; i < dataCount; i++ {
		fmt.Scanf("%d", &numArr[i])
	}

	// 除数をおこなった回数
	var repeatDiviCount int

	// 読み込んだデータの要素数が受け取るデータ数と一致しているか?
	// 要素がすべて偶数か確認 再帰関数
	// repeatDiviCountは整数型で定義しているが、&をつけることで整数データをもったアドレス情報に変換される
	// repeatCheckEvenの第三引数は*をつけてポインタ型の整数値なので引数が一致して渡せる
	repeatCheckEven(dataCount, numArr, &repeatDiviCount)
	fmt.Println(repeatDiviCount)
}

// 再帰関数
func repeatCheckEven(dataCount int, numArr []int, repeatDiviCount *int) {
	if checkEven(dataCount, numArr) {
		remakeArr(dataCount, numArr)
		*repeatDiviCount = *repeatDiviCount + 1
		repeatCheckEven(dataCount, numArr, repeatDiviCount)
	}
}

// 偶数かを確認
func checkEven(dataCount int, numArr []int) bool {
	for i := 0; i < dataCount; i++ {
		if numArr[i]%2 != 0 {
			return false
		}
	}
	return true
}

func remakeArr(dataCount int, numArr []int) []int {
	for i := 0; i < dataCount; i++ {
		numArr[i] = numArr[i] / 2
	}
	return numArr
}

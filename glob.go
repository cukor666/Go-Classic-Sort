package main

import "fmt"

type MyNumber interface {
	~int | ~int8 | ~int32 | ~int64 | ~float32 | ~float64
}

func swap[T MyNumber](arr []T, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func printArr[T MyNumber](arr []T) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func arrMin[T MyNumber](arr []T) T {
	n := len(arr)
	res := arr[0]
	for i := 1; i < n; i++ {
		res = min(arr[i], res)
	}
	return res
}

func arrMax[T MyNumber](arr []T) T {
	n := len(arr)
	res := arr[0]
	for i := 1; i < n; i++ {
		res = max(arr[i], res)
	}
	return res
}

func getDigits(num int) (digit int) {
	if num == 0 {
		return 1
	}
	for num%10 > 0 {
		digit++
		num /= 10
	}
	return
}

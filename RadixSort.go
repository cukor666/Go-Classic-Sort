package main

import "math"

func RadixSort(arr []int) {
	temp := make([][]int, 10)
	digits := getDigits(arrMax(arr))
	for k := 0; k < digits; k++ {
		for _, v := range arr {
			index := v / int(math.Pow10(k)) % 10
			temp[index] = append(temp[index], v)
		}
		j := 0
		for _, integers := range temp {
			for _, integer := range integers {
				arr[j] = integer
				j++
			}
		}
		clear(temp)
	}
}

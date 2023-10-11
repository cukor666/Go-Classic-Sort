package main

func CountSort[T ~int | ~int64 | ~int32 | ~int8](arr []T) {
	low := arrMin(arr)
	high := arrMax(arr)
	n := high - low + 1
	temp := make([]T, n)
	for _, v := range arr {
		i := v - low
		temp[i]++
	}
	var i T = 0
	var j T = 0
	for ; i < n; i++ {
		for temp[i] > 0 {
			arr[j] = i + low
			temp[i]--
			j++
		}
	}
}

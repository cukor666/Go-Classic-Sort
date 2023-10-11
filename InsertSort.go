package main

func InsertSort[T MyNumber](arr []T) {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0 && temp < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}

package main

func SelectSort[T MyNumber](arr []T) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			swap(arr, i, minIndex)
		}
	}
}

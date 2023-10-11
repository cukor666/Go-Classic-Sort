package main

func QuickSort[T MyNumber](arr []T) {
	subQuickSort(arr, 0, len(arr)-1)
}

func subQuickSort[T MyNumber](arr []T, left, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	prviot := arr[i]
	for i < j {
		for i < j && arr[j] >= prviot {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= prviot {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = prviot
	subQuickSort(arr, left, i-1)
	subQuickSort(arr, i+1, right)
}

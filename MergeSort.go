package main

func MergeSort[T MyNumber](arr []T) {
	n := len(arr)
	temp := make([]T, n)
	mergeSort(arr, 0, n-1, temp)
}

func merge[T MyNumber](arr []T, left, mid, right int, temp []T) {
	i := left
	j := mid + 1
	k := left
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}
	for x := left; x <= right; x++ {
		arr[x] = temp[x]
	}
}

func mergeSort[T MyNumber](arr []T, left, right int, temp []T) {
	if left >= right {
		return
	}
	mid := left + ((right - left) >> 1)
	mergeSort(arr, left, mid, temp)
	mergeSort(arr, mid+1, right, temp)
	merge(arr, left, mid, right, temp)
}

package main

func HeapSort[T MyNumber](arr []T) {
	n := len(arr)
	for n > 0 {
		heapify(arr, n)
		swap(arr, 0, n-1)
		n--
	}
}

func heapify[T MyNumber](arr []T, n int) {
	current := n/2 - 1
	for current >= 0 {
		i := current
		left := 2*i + 1
		for left+1 <= n {
			var best int
			if left+1 < n && arr[left+1] > arr[left] {
				best = left + 1
			} else {
				best = left
			}
			if arr[best] > arr[i] {
				swap(arr, i, best)
				i = best
				left = 2*i + 1
			} else {
				break
			}
		}
		current--
	}
}

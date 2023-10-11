package main

func ShellSort[T MyNumber](arr []T) {
	n := len(arr)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h > 0 {
		for i := h; i < n; i++ {
			temp := arr[i]
			j := i - h
			for ; j >= 0 && temp < arr[j]; j -= h {
				arr[j+h] = arr[j]
			}
			arr[j+h] = temp
		}
		h /= 3
	}
}

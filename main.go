package main

func main() {
	//arr := [...]int{5, 2, 1, 3, 4, 7, 8, 6, 9, 0, 11, 12, 10}
	arr := [...]int{0, 0, 3, 3, -5, -5, -8, -8, 1, 2, 3, 8, 4, 4, 9, 5, -5, 6, 6, 7, 2, 1}
	nums := arr[:]
	printArr(nums)
	//BubbleSort(nums)
	//SelectSort(nums)
	//InsertSort(nums)
	//ShellSort(nums)
	//QuickSort(nums)
	//MergeSort(nums)
	//HeapSort(nums)
	CountSort(nums)
	printArr(nums)
}

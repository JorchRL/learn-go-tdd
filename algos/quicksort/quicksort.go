package algos

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		p := _partition(arr, low, high)
		QuickSort(arr, low, p-1)
		QuickSort(arr, p+1, high)
	}
	return arr
}

func _partition(arr []int, p, r int) int {
	pivot := arr[r]
	i := p - 1
	for j := p; j < r; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

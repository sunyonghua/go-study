package quick_sort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var left, right []int
	flag := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < flag {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(QuickSort(left), flag), QuickSort(right)...)
}

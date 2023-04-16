package sort

import "fmt"

func QuickSort(arr []int) []int {
	arrLen := len(arr)

	if arrLen <= 1 {
		return arr
	}
	var left, right []int
	flag := arr[0]

	for i := 1; i < arrLen; i++ {
		if arr[i] < flag {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(QuickSort(left), flag), QuickSort(right)...)
}

func BubbleSort(arr []int) {
	fmt.Println("bubble sort")
}

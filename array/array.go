package main

import (
	"fmt"
	"go-study/array/quick_sort"
)

func main() {

	// 写个快排
	arr := []int{1, 4, 199, 6, 5, 10}
	var result = quick_sort.QuickSort(arr)
	fmt.Println(result)
}

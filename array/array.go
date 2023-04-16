package main

import (
	"fmt"
	"go-study/array/sort"
)

func main() {

	// 写个快排
	arr := []int{1, 4, 199, 6, 5, 10}
	var result = sort.QuickSort(arr)
	fmt.Println(result)
	sort.BubbleSort(arr)
}

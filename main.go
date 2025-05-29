package main

import (
	"fmt"
	"slices"
)

func bubbleSort(arr []int32) int32 {
	n := len(arr)
	slices.Sort(arr)
	fmt.Println("Sorted array:", arr)

	return arr[n/2]
}

func main() {
	fmt.Println("Hello, playground")
	numbers := []int32{122, 77, 12, 15, 22, 56, 99}
	fmt.Println("Unsorted array:", numbers)

	myNum := bubbleSort(numbers)
	fmt.Println("Middle:", myNum)
}

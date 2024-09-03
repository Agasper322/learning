package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var l int
	fmt.Println("please input array length")
	fmt.Scan(&l)
	if l <= 0 {
		fmt.Println("input error")
		return
	}
	arr := createRandArray(l)
	fmt.Println("your array: ", arr)
	start := time.Now()
	sortedArray := mergeSort(arr)
	duration := time.Since(start)
	fmt.Printf("sorted array:\n%v\n", sortedArray)
	fmt.Printf("duration: %v microseconds", time.Duration.Microseconds(duration))
}

func mergeSort(arr []int) (sortedArr []int) {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}

func createRandArray(l int) (arr []int) {
	for i := 0; i < l; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

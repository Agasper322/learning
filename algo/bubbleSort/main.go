package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Time Complexity: O(n^2)
Auxiliary Space: O(1)
*/

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
	sortedArray, duration := bubbleSort(arr)
	fmt.Printf("sorted array:\n%v\n", sortedArray)
	fmt.Printf("duration: %v microseconds", time.Duration.Microseconds(duration))
}

func bubbleSort(arr []int) (sortedArr []int, duration time.Duration) {
	start := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	duration = time.Since(start)
	return arr, duration
}

func createRandArray(l int) (arr []int) {
	for i := 0; i < l; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

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
	sortedArray, duration := insertionSort(arr)
	fmt.Printf("sorted array:\n%v\n", sortedArray)
	fmt.Printf("duration: %v microseconds", time.Duration.Microseconds(duration))
}

func insertionSort(arr []int) (sortedArr []int, duration time.Duration) {
	start := time.Now()
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
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

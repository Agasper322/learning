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
	fmt.Println("your array:\n", arr)
	sortedArr, duration := selectionSort(arr)
	fmt.Printf("sorted array: \n%v\nduration: %v microseconds \n", sortedArr, time.Duration.Microseconds(duration))
}

func selectionSort(arr []int) (sortedArr []int, duration time.Duration) {
	start := time.Now()
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
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

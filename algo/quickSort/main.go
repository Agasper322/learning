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
	sortedArray := quickSort(arr)
	duration := time.Since(start)
	fmt.Printf("sorted array:\n%v\n", sortedArray)
	fmt.Printf("duration: %v microseconds", time.Duration.Microseconds(duration))
}

func quickSort(arr []int) (sortedArr []int) {
	if len(arr) < 2 {
		return arr
	}
	key := arr[0]
	var less, greater []int
	for _, i := range arr[1:] {
		if i <= key {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}
	result := append(quickSort(less), key)
	result = append(result, quickSort(greater)...)
	return result
}

func createRandArray(l int) (arr []int) {
	for i := 0; i < l; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

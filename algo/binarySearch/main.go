package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

func main() {
	var l, s int
	fmt.Println("please input array length")
	fmt.Scan(&l)
	if l <= 0 {
		fmt.Println("input error")
		return
	}
	fmt.Println("please input search value")
	fmt.Scan(&s)
	arr := createRandArray(l)
	fmt.Println("your array: ", arr)
	start := time.Now()
	sortedArr, index := binarySearch(arr, s)
	duration := time.Since(start)
	fmt.Printf("sortedArr:\n%v\n", sortedArr)
	fmt.Printf("index:\n%v\n", index)
	fmt.Printf("duration: %v microseconds", time.Duration.Microseconds(duration))
}

func binarySearch(arr []int, search int) (sortedArr []int, index int) {
	arr = sortArr(arr)
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == search {
			return arr, mid
		} else if arr[mid] < search {
			left = mid + 1
		} else {
			right = right - 1
		}
	}

	return arr, -1
}

func sortArr(arr []int) (sortedArr []int) {
	slices.Sort(arr)
	return arr
}

func createRandArray(l int) (arr []int) {
	for i := 0; i < l; i++ {
		arr = append(arr, rand.Intn(10))
	}
	return arr
}

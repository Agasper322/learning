package main

import (
	"fmt"
	"math/rand"
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
	index := linearSearch(arr, s)
	duration := time.Since(start)
	fmt.Printf("index:\n%v\n", index)
	fmt.Printf("duration: %v microseconds", time.Duration.Microseconds(duration))
}

func linearSearch(arr []int, search int) int {
	for i, v := range arr {
		if search == v {
			return i
		}
	}
	return -1
}

func createRandArray(l int) (arr []int) {
	for i := 0; i < l; i++ {
		arr = append(arr, rand.Intn(10))
	}
	return arr
}

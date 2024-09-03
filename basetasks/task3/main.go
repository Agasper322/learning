package main

import (
	"fmt"
	"sync"
)

/*
3. Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2*2+3*3+4*4….) с использованием конкурентных вычислений.
*/

func main() {
	arr := 10
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go squareSum(arr, wg)
	fmt.Println("waiting goroutines")
	wg.Wait()
	fmt.Println("done")
}

func squareSum(arr int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for i := 0; i < arr; i++ {
		if i%2 == 0 {
			sum += i * i
		}
		fmt.Println(sum)
	}
}

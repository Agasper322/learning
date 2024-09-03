package main

import (
	"fmt"
	"sync"
)

/*
Задача 2. Написать программу, которая конкурентно рассчитает значение
квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты
в stdout
*/

func main() {
	arr := []int{2, 4, 6, 8, 10}
	res := make([]int, 5)
	wg := &sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res = append(res, double(arr[i]))
		}()
		wg.Wait()
	}
	fmt.Println(res)
}

func double(i int) int {
	return i * i
}

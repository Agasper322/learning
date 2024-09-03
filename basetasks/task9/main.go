package main

import (
	"fmt"
)

/*
Напишите функцию SquareSum(nums []int) int, которая принимает массив чисел, возводит каждое число в квадрат, а затем возвращает их сумму.

Например, для []int{1, 2, 2} функция должна вернуть 9, так как 1² + 2² + 2² = 9.

Для пустого массива функция должна вернуть 0.

Примеры вызова функции:

SquareSum([]int{1, 2}) // 5
SquareSum([]int{0, 3, 4, 5}) // 50
SquareSum([]int{}) // 0
SquareSum([]int{-1, -2, -3}) // 14
*/

func main() {
	nums := []int{1, 2}
	fmt.Println(SquareSum(nums))
}

func SquareSum(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v * v
	}
	return sum
}

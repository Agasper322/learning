package main

import (
	"fmt"
)

/* Напишите функцию ReverseNumbers(numbers []int) []int, которая принимает на вход срез неопределенной длины, состоящий из целых чисел, разворачивает его и возвращает.

В случае, если функция получает пустой срез, вернуть она должна его же.
*/

func main() {
	fmt.Println(ReverseNumbers([]int{}))
}

func ReverseNumbers(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

package main

import (
	"fmt"
)

/*
Напишите функцию Average(array[] int) int, принимающую массив чисел и возвращающую среднее арифметическое для этого массива, округленное до ближайшего целого.

Среднее арифметическое (в математике и статистике) — разновидность среднего значения. Определяется как число, равное сумме всех чисел множества, делённой на их количество.

Примеры вызова функции:

Average([]int{1, 2, 3}) // 2
Average([]int{5, 3, 5, 4, 4, 5}) / 4
Average([]int{2, 5, 7}) / 5

*/

func main() {
	array := []int{2, 5, 7}
	fmt.Println(Average(array))
}

func Average(array []int) int {
	len := len(array)
	var sum int
	var average float64
	for _, v := range array {
		sum += v
	}
	average = float64(sum) / float64(len)
	integerPart := int(average)
	fractionalPart := average - float64(integerPart)

	if fractionalPart >= 0.5 {
		return integerPart + 1
	}
	return integerPart
}

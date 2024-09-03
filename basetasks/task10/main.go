package main

import (
	"fmt"
)

/*
Факториал — это произведение всех натуральных чисел от 1 до некоторого числа n включительно.

Напишите функцию Factorial(n int) int, которая будет принимать целое число n и возвращать значение факториала для него.

Считайте, что факториал от нуля равен единице.

Примеры вызова функции:

Factorial(0) // 1
Factorial(3) // 6
Factorial(5) // 120
*/

func main() {
	var n int
	fmt.Println("введите число:")
	fmt.Scan(&n)
	fmt.Println(Factorial(n))
}

func Factorial(n int) int {
	var sum int = 1
	for i := 1; i <= n; i++ {
		sum = sum * i
		fmt.Printf("sum: %d\n i:%d\n", sum, i)
	}
	return sum
}

// !5 = 1 * 2 * 3 * 4 * 5 = 120

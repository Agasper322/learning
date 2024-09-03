package main

import (
	"fmt"
	"strings"
)

/*
Палиндром — слово, предложение или последовательность символов, которая абсолютно одинаково читается как в привычном направлении, так и в обратном. К примеру, “Шалаш” — это палиндром, а “стол” и “гусеница” — нет.

Напишите функцию IsPalindrome(str string) bool, которая принимает в качестве аргумента строку str и определяет, является ли она палиндромом.

В случае, если переданная строка является палиндромом, функция должна вернуть true, иначе - false.

Регистр символов не должен оказывать влияние на результат работы функции.
*/

func main() {
	fmt.Println(IsPalindrome("Шалаш"))
}

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	word := []rune(str)
	start := 0
	end := len(word) - 1
	for start < end {
		if word[start] != word[end] {
			return false
		}
		start++
		end--
	}
	return true
}

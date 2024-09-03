package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Print("Enter integer: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Invalid input")
		return
	}
	fmt.Println(fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	arr := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return arr
}

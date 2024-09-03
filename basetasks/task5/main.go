package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Задача 5. Разработать программу, которая будет последовательно отправлять значения в канал,
 а с другой стороны канала — читать. По истечению N секунд программа должна завершаться
*/

func main() {
	var n int
	fmt.Println("Введите количество секунд, сколько будет работать программа:")
	fmt.Scan(&n)
	channel := make(chan int)
	defer close(channel)
	go write(channel)
	go read(channel)
	time.Sleep(time.Duration(n) * time.Second)
}

func write(channel chan int) {
	for {
		data := rand.Intn(10)
		channel <- data
	}
}

func read(channel chan int) {
	for data := range channel {
		fmt.Printf("data is: %d\n", data)
	}
}

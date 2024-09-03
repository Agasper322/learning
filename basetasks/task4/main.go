package main

import (
	"fmt"
	"time"
)

/*
	Задача 4. Реализовать постоянную запись данных в канал (главный поток). Реализовать
	набор из N воркеров, которые читают произвольные данные из канала и
	выводят в stdout. Необходима возможность выбора количества воркеров при
	старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
	способ завершения работы всех воркеров.
*/

func main() {
	var numWorkers int
	fmt.Println("введите количество горутин:")
	fmt.Scan(&numWorkers)
	channel := make(chan int)
	for i := 0; i <= numWorkers; i++ {
		go worker(i, channel)
	}

	for {
		channel <- time.Now().Second()
		time.Sleep(time.Millisecond * 100)
	}

}

func worker(name int, channel <-chan int) {
	for {
		num := <-channel
		fmt.Printf("Горутина: %d, значение: %d\n", name, num)
	}
}

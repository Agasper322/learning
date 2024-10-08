package main

import (
	"fmt"
)

/*
По городу движется автобус, который на каждой остановке забирает и высаживает несколько человек.

Напишите функцию PassengersCount(busStops [][2]int) int, которая должна вернуть количество людей, оставшихся в автобусе после последней остановки. Несмотря на то, что это последняя остановка, автобус может быть не пустым, и некоторые люди все еще могут находиться в автобусе (возможно, они спят).

busStops это массив (срез) пар целых чисел. Каждая пара представляет собой информацию о количестве людей, которые садятся в автобус (первое число) и количество людей, которые выходят из автобуса (второе число).

Примеры вызова функции:

Number([][2]int{{10, 0}}) // 10
Number([][2]int{{10, 0}, {3, 5}, {5, 8}}) // 5
Number([][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}) // 21
Number([][2]int{{0, 0}}) // 0
*/

func main() {
	fmt.Println(PassengersCount([][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}}))
}

func PassengersCount(busStops [][2]int) int {
	var enter int
	for i := range len(busStops) {
		enter += busStops[i][0] - busStops[i][1]

	}
	return enter
}

package main

import (
	"fmt"
)

/*

Напишите функцию GetIntersection(a, b []int) которая принимает два целочисленных массива произвольной длинны и возвращает новый массив, представляющий их пересечение.

Пересечение массивов включает только те элементы, которые присутствуют в обоих массивах.

Например, пересечением массивов [1, 2, 2, 3] и [2, 2, 4] будет [2, 2], потому что число 2 появляется дважды и в первом, и во втором.

Примеры вызова функции:

GetIntersection([]int{1, 2, 2, 1}, []int{2, 2}))  // []int{2, 2}
GetIntersection([]int{8, 5, 3}, []int{6, 8, 5}) // []int{8, 5}
GetIntersection([]int{7, 7, 3}, []int{7, 5, 7, 2}) // []int{7, 7}
*/

func main() {
	fmt.Println(GetIntersection([]int{8, 5, 3}, []int{6, 8, 5}))
}

func GetIntersection(a, b []int) []int {
	someMap := make(map[int]int)
	var intersection []int
	for _, v := range a {
		someMap[v]++
	}
	for _, v := range b {
		if count, ok := someMap[v]; ok && count > 0 {
			intersection = append(intersection, v)
			someMap[v]--
		}
	}
	return intersection
}

package main

import "fmt"

/*
Задача 1. Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
*/

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
	Doing string
}

func main() {
	action := Action{
		Human: Human{Name: "Nikita", Age: 27},
		Doing: "programming",
	}
	fmt.Println(action)
}

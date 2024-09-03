package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(High("Я забыл дневник"))
}

func High(s string) string {
	var words []string
	var word string
	words = strings.Split(s, " ")
	for i := range words {
		if len(words[i]) > len(word) {
			word = words[i]
		}

	}
	return word
}

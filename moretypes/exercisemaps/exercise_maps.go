package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	word_count := make(map[string]int)
	words := strings.Split(s, " ")
	for _, word := range words {
		word_count[word] += 1
	}

	return word_count
}

func main() {
	fmt.Println(WordCount("Hello World World Bye"))
}

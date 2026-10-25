package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	var result int

	f := func() int {
		result, a, b = a, b, a+b
		return result
	}

	return f
}

func main() {
	f := fibonacci()

	for range 10 {
		fmt.Println(f())
	}
}

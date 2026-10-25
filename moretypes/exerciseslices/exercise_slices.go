package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
	}

	for x := range dx {
		y := math.Pow(float64(x), 2)
		if int(y) < dy {
			pic[dy-1-int(y)][x] = 1
		}

	}

	for _, row := range pic {
		var row_strs []string
		for _, value := range row {
			row_strs = append(row_strs, strconv.Itoa(int(value)))
		}
		fmt.Println(strings.Join(row_strs, " "))
	}

	return pic
}

func main() {
	Pic(30, 30)
}

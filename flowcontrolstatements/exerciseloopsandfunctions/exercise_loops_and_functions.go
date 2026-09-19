package main

import "fmt"

func isSqrt(z, x, delta float64) bool {
	return z*z >= x-delta && z*z <= x+delta
}

func Sqrt(x float64, delta float64) float64 {
	z := 1.0
	for !isSqrt(z, x, delta) {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2, 0.00000001))
}

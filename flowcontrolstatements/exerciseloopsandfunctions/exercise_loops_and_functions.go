package main

import "fmt"

func isDeltaEqual(y, x, delta float64) bool {
	return y >= x-delta && y <= x+delta
}

func Sqrt(x float64, delta float64) float64 {
	z := 1.0
	z_squared := z * z
	for !isDeltaEqual(z_squared, x, delta) {
		z -= (z_squared - x) / (2 * z)
		z_squared = z * z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2, 0.00000001))
}

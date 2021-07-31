package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("Negative numbers don't have square roots!")
	}
	z := 1. // estimation of sqrt(x)
	y := 0. // previous value of z
	for math.Abs(z-y) > 0.00000000001 {
		y = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2)) // standard library
}

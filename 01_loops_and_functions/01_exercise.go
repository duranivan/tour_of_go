package main

import (
	"fmt"
	"math"
)

// Sqrt computes the square root of x using Newton's method
// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	i := 1
// 	for i <= 11 {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Println(z)
// 		i++
// 	}
// 	return z
// }

// Sqrt computes the square root of x using Newton's method
func Sqrt(x float64) float64 {
	z := x / 2
	n := 0
	for {
		if math.Abs((z*z-x)/(2*z)) < 1e-7 {
			break
		}
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		n++
	}

	fmt.Println("Number of iterations:", n)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

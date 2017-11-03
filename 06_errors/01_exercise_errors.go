package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt computes the square root of x using Newton's method
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
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
	return z, nil
}

func printSqrt(x float64) {
	result, e := Sqrt(x)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(result)
}

func main() {
	// fmt.Println(ErrNegativeSqrt(-2).Error())
	// fmt.Println(Sqrt(2))
	// fmt.Println(Sqrt(-2))
	printSqrt(2)
	printSqrt(-2)
}

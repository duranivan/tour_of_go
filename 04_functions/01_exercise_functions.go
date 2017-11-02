// Implement a fibonacci function that returns a function (a closure) that
// returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...)

package main

import "fmt"

// Returns the xth fibonacci number using recursion
func fibo(x int) int {
	if x == 1 {
		return 0
	}
	if x == 2 {
		return 1
	}
	return fibo(x-1) + fibo(x-2)
}

func fibonacci() func(int) int {
	return func(x int) int {
		n0, n1 := 0, 1
		if x == 1 {
			return n0
		}
		if x == 2 {
			return n1
		}

		for i := 3; i <= x; i++ {
			n0, n1 = n1, n0+n1
		}
		return n1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i + 1))
	}

	fmt.Println(fibo(6))
}

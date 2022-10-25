package main

import "fmt"

func Fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Factorial(n uint) (result uint) {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}

func main() {
	fmt.Printf("%d\n", Fibonacci(7))
	fmt.Printf("%d\n", Factorial(7))
}

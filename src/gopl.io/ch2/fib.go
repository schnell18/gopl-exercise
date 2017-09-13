package main

import "fmt"

func main() {
	fmt.Printf("fib(%d) = %d\n", 1, fib(1))
	fmt.Printf("fib(%d) = %d\n", 2, fib(2))
	fmt.Printf("fib(%d) = %d\n", 3, fib(3))
	fmt.Printf("fib(%d) = %d\n", 4, fib(4))
	fmt.Printf("fib(%d) = %d\n", 5, fib(5))
	fmt.Printf("fib(%d) = %d\n", 6, fib(6))
	fmt.Printf("fib(%d) = %d\n", 7, fib(7))
	fmt.Printf("fib(%d) = %d\n", 8, fib(8))
	fmt.Printf("fib(%d) = %d\n", 9, fib(9))
	fmt.Printf("fib(%d) = %d\n", 10, fib(10))
	fmt.Printf("fib(%d) = %d\n", 11, fib(11))
}

func fib(n int) int {
	a, b := 1, 1
	for i := 2; i < n; i++ {
		a, b = b, a + b
	}
    return b
}


package main

import "fmt"

func main() {
	fmt.Printf("gcd(%d, %d) = %d\n", 22, 4, gcd(22, 4))
	fmt.Printf("gcd(%d, %d) = %d\n", 13, 23, gcd(13, 23))
	fmt.Printf("(%d mod %d) = %d\n", 13, 23, 13%23)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
    return x
}


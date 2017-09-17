package main

import (
	"fmt"
)

func main() {
	s := "善始善终"
	fmt.Printf("% X\n", s)
	r := []rune(s)
	fmt.Printf("%#X\n", r)
	fmt.Printf("%d\n", r)
	fmt.Println(string(12345678))
}
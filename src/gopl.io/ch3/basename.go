package main

import (
	"fmt"
)

func main() {
	s := "善/始/善/终.d.go"
	fmt.Printf("%s\n", basename(s))
	s = "/Users/justin/work/mandelbrot.go"
	fmt.Printf("%s\n", basename(s))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i--{
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i--{
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
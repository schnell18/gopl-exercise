package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	tc "gopl.io/ch2/tempconv"
)

func main() {
	dps := os.Args[1:]
	if len(dps) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			c := convert(input.Text())
			fmt.Printf("%v => %v\n", tc.Celsius(c), tc.CToF(tc.Celsius(c)))
		}
	} else {
		for _, r := range os.Args[1:] {
			c := convert(r)
			fmt.Printf("%v => %v\n", tc.Celsius(c), tc.CToF(tc.Celsius(c)))
		}
	}
}

func convert(raw string) float64 {
	c, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid temperature %s\n", raw)
		return 0.0
	}
	return c
}


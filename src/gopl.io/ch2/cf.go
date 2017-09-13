package main

import (
	"fmt"
	"os"
	"strconv"
	tc "gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tc.Fahrenheit(t)
		c := tc.Celsius(t)
		k := tc.Kelvin(t)
		fmt.Printf(
			"%s = %s, %s = %s, %s = %s\n",
			f, tc.FToC(f), c, tc.CToF(c), k, tc.KToC(k))
	}
}



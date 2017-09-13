package main

import (
	"fmt"
	tc "gopl.io/ch2/tempconv"
)

func main() {
	c := tc.FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
	k := tc.KToC(tc.Kelvin(300.0))
	fmt.Println(k.String())
	fmt.Printf("%v\n", k)
	fmt.Printf("%s\n", k)
	fmt.Println(k)
	fmt.Printf("%g\n", k)
	fmt.Println(float64(k))
}



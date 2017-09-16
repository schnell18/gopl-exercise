package main

import (
	"os"
)

func main() {
	width := 600
	height := 320
	writeSurface(os.Stdout, width, height)
}
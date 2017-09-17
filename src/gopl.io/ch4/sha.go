package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"flag"
	"os"
	"bufio"
)

func main() {
	algPtr := flag.String("m", "sha256", "Specify the hash algorithm to use")
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		switch *algPtr {
			case "sha256":
				fmt.Printf("%X\n", sha256.Sum256(input.Bytes()))
			case "sha384":
				fmt.Printf("%X\n", sha512.Sum384(input.Bytes()))
			case "sha512":
				fmt.Printf("%X\n", sha512.Sum512(input.Bytes()))
		}
	}
}
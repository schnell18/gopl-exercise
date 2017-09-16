package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
	    w.Header().Set("Content-Type", "image/png")
		writeImage(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("Listening on localhost:8000...")
}
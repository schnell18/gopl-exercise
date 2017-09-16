package main

import (
	"log"
	"net/http"
	"strconv"
	"fmt"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		width := 600
		height := 320
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			if k == "width" {
				width, _ = strconv.Atoi(v[0])
			} else if k == "height" {
				height, _ = strconv.Atoi(v[0])
			}
		}
	    w.Header().Set("Content-Type", "image/svg+xml")
		writeSurface(w, width, height)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("Listening on localhost:8000...")
}
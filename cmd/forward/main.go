package main

import (
	"log"
	"net/http"

	"github.com/btwiuse/forward"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(forward.Forward))
	if err != nil {
		log.Fatalln(err)
	}
}

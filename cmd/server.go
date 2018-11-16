package main

import (
	"log"
	"net/http"

	"github.com/eidosam/tortoise/pkg/tortoise"
)

func main() {
	http.HandleFunc("/", tortoise.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"log"
	"net/http"

	_ "router"
)

func main() {
	log.Println("Server is running at :9090.")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

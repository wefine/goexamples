package main

import (
	"fmt"
	"net/http"
    "time"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("hello")

	fmt.Fprint(w, "Hello, world!")
	time.Sleep(time.Second * 10)
}

func main() {
    log.Println("loading...")
    time.Sleep(time.Second * 10)
    log.Println("loading...finished!")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

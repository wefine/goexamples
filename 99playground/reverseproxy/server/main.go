package main

import (
	"net/http"
	"os"
	"strconv"
	"log"
)

func main() {
	port := "9091"

	if len(os.Args) > 1 {
		if _, err := strconv.Atoi(os.Args[1]); err == nil {
			port = os.Args[1]
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		println("--->", port, req.URL.String())
	})

	log.Println("Listen on : ", port)
	http.ListenAndServe(":"+port, nil)
}

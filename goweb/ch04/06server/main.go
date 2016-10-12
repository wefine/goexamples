package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	runWithMux(server)
	//runWithHandlerFun()

	log.Println("Listening...")
	server.ListenAndServe()
}

func runWithHandlerFun() {
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Go Web Development")
	})
}

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func runWithMux(server *http.Server) {
	mux := http.NewServeMux()
	mh1 := &messageHandler{"Welcome to Go Web Development"}
	mux.Handle("/welcome", mh1)

	server.Handler = mux
}

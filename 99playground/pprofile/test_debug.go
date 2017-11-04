package main

import (
	"log"
	"net/http"
	_ "net/http"
)
import "net/http/pprof"

func main() {
	http.HandleFunc("/debug/pprof/", pprof.Index)
	http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	log.Println(http.ListenAndServe(":7070", nil))
}

// http://localhost:6060/debug/pprof

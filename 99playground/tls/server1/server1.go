package main

import (
    "net/http"
    "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server.\n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}

func main() {
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServeTLS(":7443", "/ssl/server.crt", "/ssl/server.key", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

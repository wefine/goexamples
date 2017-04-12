package simple

import (
	"net/http/httputil"
	"net/url"
	"net/http"
	"log"
)

func main() {
	port := "9090"

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:9091",
	})

	log.Println("Listen on : ", port)
	http.ListenAndServe(":"+port, proxy)
}

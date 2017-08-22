package main

import "./reverseproxy"

func main() {
    reverseproxy.NewReverseProxy(reverseproxy.ProxyConfig{
        Cert:   "/ssl/server.crt",
        Key:    "/ssl/server.key",
        UseTLS: true,
        Listen: ":8888",
        To:     "http://localhost:8000",
    })
}

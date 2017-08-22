package main

import "./reverseproxy"

func main() {
    go reverseproxy.NewWsProxy(reverseproxy.ProxyConfig{
        Cert:   "/ssl/server.crt",
        Key:    "/ssl/server.key",
        UseTLS: true,
        Listen: ":6666",
        To:     "http://localhost:8080",
    })

    reverseproxy.NewHttpProxy(reverseproxy.ProxyConfig{
        Cert:   "/ssl/server.crt",
        Key:    "/ssl/server.key",
        UseTLS: true,
        Listen: ":8888",
        To:     "http://localhost:8000",
    })
}

package main

import (
    "net/http/httputil"
    "net/url"
    "log"
    "time"
    "net/http"
)

func main() {
    NewHttpProxy(ProxyConfig{
        Listen:":8888",
        To:"http://localhost:80",
    })
}

type ProxyConfig struct {
    Listen        string
    To            string
    Cert          string
    Key           string
    UseTLS        bool
    IsWebSocket   bool
    useLogging    bool
    flushInterval time.Duration
}

func parseTo(config ProxyConfig) (*url.URL, error) {
    to, err := url.Parse(config.To)
    if err != nil {
        log.Fatalln("Fatal parsing -to:", err)
        return nil, err
    }

    log.Printf("listen on: %s", config.Listen)
    log.Printf("proxy to: %s", config.To)

    return to, nil
}

func NewHttpProxy(config ProxyConfig) {
    // default value
    config.flushInterval = 0

    to, err := parseTo(config)
    if err != nil {
        return
    }

    hp := httputil.NewSingleHostReverseProxy(to)
    hp.FlushInterval = config.flushInterval

    var handler http.Handler
    handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        hp.ServeHTTP(w, r)
    })

    server := &http.Server{Addr: config.Listen, Handler: handler}
    log.Fatalln(server.ListenAndServe())
}

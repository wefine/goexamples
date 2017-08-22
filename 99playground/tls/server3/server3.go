package main

import (
    "crypto/tls"
    "log"
    "net/http"
    "net/url"
    "./wsutil"
    "./websocketproxy"
)

func getMuxHandler() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
        w.Write([]byte("This is an example server.\n"))
    })

    return mux
}

func getWsHandler1() http.Handler  {
    backendURL := &url.URL{Scheme: "wss://", Host: "10.114.51.34:6443"}
    p := wsutil.NewSingleHostReverseProxy(backendURL)

    return p
}

func getWsHandler2() http.Handler  {
    u := &url.URL{Scheme: "ws://", Host: "10.114.51.34:8080"}

    return websocketproxy.NewProxy(u)
}

func main() {

    cfg := &tls.Config{
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
            tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
        },
    }
    srv := &http.Server{
        Addr:         ":7443",
        Handler:      getWsHandler2(),
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    log.Fatal(srv.ListenAndServeTLS("/ssl/server.crt", "/ssl/server.key"))
}

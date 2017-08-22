package main

import (
    "crypto/tls"
    "log"
    "net/http"
    "net/url"
    "./wsutil"
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
    backendURL := &url.URL{Scheme: "ws://", Host: "localhost:8080"}
    p := wsutil.NewSingleHostReverseProxy(backendURL)

    return p
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

    listen := ":7553"
    log.Println(listen)
    srv := &http.Server{
        Addr:         listen,
        Handler:      getWsHandler1(),
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    log.Fatal(srv.ListenAndServeTLS("/ssl/server.crt", "/ssl/server.key"))
}

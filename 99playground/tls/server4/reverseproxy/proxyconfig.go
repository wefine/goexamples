package reverseproxy

import (
    "log"
    "net/url"
    "crypto/tls"
    "net/http"
    "./wsproxy"
)

type ProxyConfig struct {
    Listen, To string
    Cert, Key  string
    UseTLS     bool
}

func NewReverseProxy(config ProxyConfig) {

    to, err := url.Parse(config.To)
    if err != nil {
        log.Fatalln("Fatal parsing -where:", err)
    }

    proxy := wsproxy.NewSingleHostReverseProxy(to)

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

    log.Printf("listen on: %s", config.Listen)
    log.Printf("proxy to: %s", config.To)

    if config.UseTLS {
        srv := &http.Server{
            Addr:         config.Listen,
            Handler:      proxy,
            TLSConfig:    cfg,
            TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
        }
        log.Fatal(srv.ListenAndServeTLS("/ssl/server.crt", "/ssl/server.key"))
    } else {
        srv := &http.Server{
            Addr:    config.Listen,
            Handler: proxy,
        }
        log.Fatal(srv.ListenAndServe())
    }

}

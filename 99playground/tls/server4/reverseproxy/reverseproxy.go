package reverseproxy

import (
    "log"
    "net/url"
    "crypto/tls"
    "net/http"
    "net"
    "io/ioutil"
    "bytes"
    "./ws"
    "time"
    "net/http/httputil"
)

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

type ConnectionErrorHandler struct{ http.RoundTripper }

func (c *ConnectionErrorHandler) RoundTrip(req *http.Request) (*http.Response, error) {
    resp, err := c.RoundTripper.RoundTrip(req)
    if _, ok := err.(*net.OpError); ok {
        r := &http.Response{
            StatusCode: http.StatusServiceUnavailable,
            Body:       ioutil.NopCloser(bytes.NewBufferString("503 Backend Unavailable")),
        }
        return r, nil
    }
    return resp, err
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

func NewWsProxy(config ProxyConfig) {
    to, err := parseTo(config)
    if err != nil {
        return
    }

    proxy := ws.NewSingleHostReverseProxy(to)

    if config.UseTLS {
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

func NewHttpProxy(config ProxyConfig) {
    // default value
    config.useLogging = true
    config.flushInterval = 0

    to, err := parseTo(config)
    if err != nil {
        return
    }

    hp := httputil.NewSingleHostReverseProxy(to)
    hp.Transport = &ConnectionErrorHandler{http.DefaultTransport}
    hp.FlushInterval = config.flushInterval

    var handler http.Handler
    handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        hp.ServeHTTP(w, r)
    })

    if config.useLogging {
        handler = &LoggingMiddleware{handler}
    }

    server := &http.Server{Addr: config.Listen, Handler: handler}

    switch {
    case config.UseTLS:
        err = server.ListenAndServeTLS(config.Cert, config.Key)
    default:
        err = server.ListenAndServe()
    }

    log.Fatalln(err)
}

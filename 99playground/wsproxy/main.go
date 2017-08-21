package main

import (
    "./wsutil"
    "net/http"
    "net/url"
)

func main() {
    //backend := ":8080"
    proxy := ":8080"

    // an webscket echo server
    //backendHandler := websocket.Handler(func(ws *websocket.Conn) {
    //	io.Copy(ws, ws)
    //})
    //// run both servers and give them a second to start up
    //http.ListenAndServe(backend, backendHandler)

    // make a proxy pointing at that backend url
    backendURL := &url.URL{Scheme: "ws://", Host: "localhost:8080"}
    p := wsutil.NewSingleHostReverseProxy(backendURL)
    http.ListenAndServe(proxy, p)
    //go http.ListenAndServe(backend, backendHandler)
    //go http.ListenAndServe(proxy, p)
    //time.Sleep(1 * time.Second)

    //// connect to the proxy
    //origin := "http://localhost/"
    //ws, err := websocket.Dial("ws://"+proxy, "", origin)
    //if err != nil {
    //    log.Fatal(err)
    //}
    //
    //// send a message along the websocket
    //msg := []byte("isn't yhat awesome?")
    //if _, err := ws.Write(msg); err != nil {
    //    log.Fatal(err)
    //}
    //
    //// read the response from the proxy
    //resp := make([]byte, 4096)
    //if n, err := ws.Read(resp); err != nil {
    //    log.Fatal(err)
    //} else {
    //    fmt.Printf("%s\n", resp[:n])
    //}
}

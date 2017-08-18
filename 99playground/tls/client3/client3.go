package main

import (
    "log"
    "crypto/tls"
    "net/http"
    "io/ioutil"
)

func main() {

    // Setup HTTPS client
    tlsConfig := &tls.Config{
        InsecureSkipVerify: true,
    }

    tlsConfig.BuildNameToCertificate()
    transport := &http.Transport{TLSClientConfig: tlsConfig}
    client := &http.Client{Transport: transport}

    req, err := http.NewRequest("GET", "https://localhost:6443/version", nil)
    req.SetBasicAuth("admin", "see@123")
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    token := resp.Request.Header.Get("authorization")

    log.Println("token=" + token[6:])

    defer resp.Body.Close()

    // Dump response
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(string(data))
}

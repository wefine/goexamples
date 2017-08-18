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

    // Do GET something
    resp, err := client.Get("https://127.0.0.1:7443/")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // Dump response
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(string(data))
}

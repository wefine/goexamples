package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "net/url"
)

func remoteResponse(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    return string(body);
}

func main() {
    os.Clearenv()
    remote := "http://localhost:9090/patch/jos/v1/configmaps"
    u, _ := url.Parse(remote)

    q := u.Query()
    q.Set("username", "user")
    q.Set("password", "passwd")
    u.RawQuery = q.Encode()

    result := remoteResponse(u.String())
    fmt.Printf("%s", result)
}

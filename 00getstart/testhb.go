package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "strings"
    "time"
)

func main() {
    hbUrl := "http://192.168.1.11:8124/heartbeat.jsp"

    doHeartBeat(hbUrl)
}

func doHeartBeat(hbUrl string) {
    time.Sleep(time.Second * 10)

    resp, err := http.Get(hbUrl)
    if err != nil {
        log.Println("Can not connect, sleep 10s and then try to check again!")
        doHeartBeat(hbUrl)
    } else {
        body, _ := ioutil.ReadAll(resp.Body)
        content := strings.TrimSpace(string(body))
        resp.Body.Close()

        if content != "APPOK" {
            log.Println("Response is not 'APPOK', sleep 10s and then try to check again!")
            doHeartBeat(hbUrl)
        } else {
            log.Println("App is works now!")
        }
    }
}

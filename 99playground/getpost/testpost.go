package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	os.Setenv("http_proxy", "")

	url := "http://devdocker:2375/containers/5a5b816b1000af689dc09e2e40d8d5411db949f1a09e3063d8be8b892fd80c42/exec"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{
	  "AttachStdin": false,
	  "AttachStdout": true,
	  "AttachStderr": true,
	  "DetachKeys": "ctrl-p,ctrl-q",
	  "Tty": false,
	  "Cmd": [
		"date"
	  ],
	  "Env": [
		"FOO=bar",
		"BAZ=quux"
	  ]
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

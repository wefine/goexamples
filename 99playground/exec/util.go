package main

import (
    "sync"
    "strings"
    "os/exec"
    "os"
    "log"
)

func ExeCmd(cmd string) {
    wg := new(sync.WaitGroup)

    wg.Add(1)
    go _exeCmd(cmd, wg)

    wg.Wait()
}

func _exeCmd(cmd string, wg *sync.WaitGroup) {
    log.Printf("Running command: %s\n", cmd)

    parts := strings.Fields(cmd)
    head := parts[0]
    parts = parts[1:]

    out, err := exec.Command(head, parts...).Output()
    if err != nil {
        log.Printf("Error: %s\n", err)
    }
    log.Printf("Command output: %s\n\n", out)
    wg.Done()
}

func CheckPathExisted(path string) bool {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return false
    }

    return true
}

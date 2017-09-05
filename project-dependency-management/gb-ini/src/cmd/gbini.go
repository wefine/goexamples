package main

import (
    "os/user"
    "log"
    "os"
    "fmt"
)

func main() {

    home := os.Getenv("HOME")
    log.Println(home)

    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    log.Println(usr.HomeDir)

    user, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Hello " + user.Name)
    fmt.Println("====")
    fmt.Println("Id: " + user.Uid)
    fmt.Println("Username: " + user.Username)
    fmt.Println("Home Dir: " + user.HomeDir)
}
